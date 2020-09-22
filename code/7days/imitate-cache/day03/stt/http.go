package stt

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const defaultBasePath = "/_cache/"

type HTTPPool struct {
	baseUri  string // 记录主机名和ip地址与端口
	basePath string // 作为通信节点地址的前缀，http://example.com/_cache/
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
