package session

import (
	"sync"
	"time"
	"video-server/api-server/dao"
	"video-server/api-server/domain/entity"
	"video-server/api-server/utils"
)

// 在go 1.9之后出现，对读速度很快，写有个全局锁，比较慢
var sessionMap *sync.Map

// 在main方法之前被调用
func init() {
	sessionMap = &sync.Map{}
}

// LoadSessionsFromDB 加载数据库中的session信息
func LoadSessionsFromDB() {
	m, err := dao.GetSessionsMap()
	if err != nil {
		return
	}
	// 将返回的map中的数据写入到cache中
	m.Range(func(k, v interface{}) bool {
		ss := v.(*entity.Session)
		sessionMap.Store(k, ss)
		return true
	})
}

// GenerateNewSessionID 创建一个sessionId 先存储到mysql中，然后存储到cache中
func GenerateNewSessionID(username string) (string, error) {
	id, _ := utils.NewUUID()
	// 设置过期时间
	// server side  session valid time 30 *60
	ttl := nowInMilli() + 30*60*1000

	// 保存到本地数据库
	err := dao.AddSession(id, ttl, username)
	if err != nil {
		return "", err
	}
	// 保存到cache中
	sessionMap.Store(id, &entity.Session{username, ttl})
	return id, nil
}

// IsSessionExpired 没有过期，返回(name,false) 过期返回("",true)
func IsSessionExpired(sid string) (string, bool) {

	v, isOk := sessionMap.Load(sid)
	// 不存在说明过期
	if !isOk {
		return "", true
	}

	session := v.(*entity.Session)
	ct := nowInMilli()
	// 超时
	if session.TTL > ct {
		// 过期删除数据库
		sessionMap.Delete(sid)
		dao.DeleteSession(sid)
		return "", true
	}
	// 没有超时
	return session.Username, false
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}
