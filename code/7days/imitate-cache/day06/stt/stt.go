package stt

import (
	"fmt"
	"log"
	"stt/singleflight"
	"sync"
)

// 回调接口，用户用于返回对应的数值
type Getter interface {
	Get(key string) ([]byte, error)
}

// 定义一个方法对象 ,作用：强制转换为Getter对象，然后使用Getter对象调用Get方法执行
type GetterFunc func(key string) ([]byte, error)

// 该方法对象实现了Getter接口的 Get方法
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// cache 执行的主体，用于查询和插入缓存，如果缓存中不存在，则查询远程节点，如果远程节点没有，则调用回调函数获取插入并返回
type Group struct {
	name      string // group name
	getter    Getter
	mainCache cache
	peers     PeerPicker // 获取远端节点

	loader *singleflight.Group // 多次相同key的请求只查询一次
}

var (
	rwMu sync.RWMutex
	// 缓存命名空间
	groups = make(map[string]*Group)
)

func NewGroup(name string, capacity int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	rwMu.Lock()
	defer rwMu.Unlock()
	g := &Group{
		name:   name,
		getter: getter,
		mainCache: cache{
			capacity: capacity,
		},
		loader: &singleflight.Group{},
	}
	groups[name] = g
	return g
}

func GetGroup(name string) *Group {
	rwMu.RLock()
	g := groups[name]
	rwMu.RUnlock()
	return g
}

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	// 从本地缓存中获取
	value, ok := g.mainCache.get(key)
	if ok {
		log.Println("cache hit")
		return value, nil
	}
	// 本地缓存不存在，则去其他缓存获取或者重新拉取数据
	return g.load(key)
}

func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{cloneBytes(bytes)}
	// 存储在本地缓存中
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}

func (g *Group) RegisterPeers(peers PeerPicker) {
	if g.peers != nil {
		panic("RegisterPeerPicker called more than once")
	}
	g.peers = peers
}

func (g *Group) load(key string) (value ByteView, err error) {
	// 防止缓存击穿
	val, err := g.loader.Do(key, func() (interface{}, error) {
		// 如果远端有，则尝试从远端获取
		if g.peers != nil {
			peer, ok := g.peers.PickPeer(key)
			if ok {
				value, err := g.getFromPeer(peer, key)
				if err == nil {
					return value, nil
				}
				log.Println("failed get from peer", err)
			}
		}
		// 从本地获取
		return g.getLocally(key)
	})

	if err == nil {
		return val.(ByteView), nil
	}
	return
}

func (g *Group) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	value, err := peer.Get(g.name, key)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: value}, nil
}
