package stt

// 查询hash表，获取远端节点，返回远端节点的客户端
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// 通过HTTP客户端访问远端节点，拉取数据
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
