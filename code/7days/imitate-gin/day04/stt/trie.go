package stt

import "strings"

type node struct {
	pattern  string  // 待匹配的路由 从root 到当前的路由，在叶子节点时才会有值
	part     string  // 路由的一部分
	children []*node //  子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 插入时匹配使用，返回单个
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		// 如果匹配到了路由的一部分，或者是模糊匹配，则返回下一个节点
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 查询时匹配使用，返回多个
func (n *node) matchChildren(part string) []*node {
	// 如果多个子节点匹配成功，则返回一个数组
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 递归插入，第一次调用 current 是 0
func (n *node) insert(pattern string, parts []string, current int) {
	if len(parts) == current {
		n.pattern = pattern
		return
	}
	// 获取url的一部分
	part := parts[current]
	// 在当前节点的子节点中查询，看是否可以匹配到
	child := n.matchChild(part)
	// 节点不存在，则创建节点
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*', // 字符串第一个值是: 或者 * 表示模糊匹配
		}
		n.children = append(n.children, child)
	}
	// 继续下一层
	child.insert(pattern, parts, current+1)
}

// 查询获取匹配点，递归查询
func (n *node) search(parts []string, current int) *node {
	// 到达parts的大长度，或者 n当前是模糊匹配
	if len(parts) == current || strings.HasPrefix(n.part, "*") {
		// n当前节点 非叶子节点，没有匹配到
		if n.pattern == "" {
			return nil
		}
		// 是叶子节点或者达到了模糊匹配的条件，可以将当前进行返回
		return n
	}

	part := parts[current]
	// 继续往下匹配节点，获得到匹配的节点数组
	children := n.matchChildren(part)
	for _, child := range children {
		// 深度搜索
		result := child.search(parts, current+1)
		if result != nil {
			return result
		}
	}
	return nil
}
