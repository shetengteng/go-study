package stt

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)


// http客户端
type httpClient struct {
	baseURL string
}

// 访问远端节点group内的key对应的值
func (h *httpClient) Get(group string, key string) ([]byte, error) {
	// url.QueryEscape 对group和key中的特殊字符进行转义为base64进行传输
	u := fmt.Sprintf("%v%v/%v", h.baseURL, url.QueryEscape(group), url.QueryEscape(key))
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned: %v", res.Status)
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}
	return bytes, nil
}

// 表示编译器需要检查 httpClient需要实现 PeerGetter接口
var _ PeerGetter = (*httpClient)(nil)
