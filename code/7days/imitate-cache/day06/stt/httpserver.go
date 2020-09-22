package stt

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"stt/consistenthash"
	"sync"
)

const (
	defaultBasePath = "/_cache/"
	defaultReplicas = 50 // 默认每个节点有50个虚拟节点
)

type HTTPPool struct {
	baseUri  string // 记录主机名和ip地址与端口
	basePath string // 作为通信节点地址的前缀，http://example.com/_cache/

	mu          sync.Mutex //
	peers       *consistenthash.Map
	httpClients map[string]*httpClient // keyed by e.g. "http://localhost:8088
}

func NewHTTPPool(baseUri string) *HTTPPool {
	return &HTTPPool{
		baseUri:  baseUri,
		basePath: defaultBasePath,
	}
}

func (p *HTTPPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", p.baseUri, fmt.Sprintf(format, v...))
}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 访问指定的前缀的路径
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpected path:" + r.URL.Path)
	}
	p.Log("%s %s", r.Method, r.URL.Path)
	// 解析路径，并到指定的缓存中获取数据
	// /<basepath>/<groupname>/<key>
	// 分割获取指定的groupname和key
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0]
	key := parts[1]
	group := GetGroup(groupName)

	if group == nil {
		http.Error(w, "no such group", http.StatusNotFound)
		return
	}
	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// application/octet-stream默认的未知类型
	//w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Type", "text/plain")
	w.Write(view.ByteSlice())
}

// peers 是主机ip与port
func (p *HTTPPool) Set(peers ...string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.peers = consistenthash.New(defaultReplicas, nil)
	p.peers.Add(peers...)
	p.httpClients = make(map[string]*httpClient, len(peers))
	for _, peer := range peers {
		p.httpClients[peer] = &httpClient{
			baseURL: peer + p.basePath,
		}
	}
}

func (p *HTTPPool) PickPeer(key string) (PeerGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	peer := p.peers.Get(key)
	if peer != "" && peer != p.baseUri {
		return p.httpClients[peer], true
	}
	return nil, false
}

var _ PeerPicker = (*HTTPPool)(nil)
