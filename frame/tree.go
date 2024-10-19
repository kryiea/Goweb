package frame

import (
	"errors"
	"strings"
)

// 代表树结构
type Tree struct {
	root *node // 根节点
}

// 代表节点
type node struct {
	isLast   bool                // 代表这个节点是否可以成为最终的路由规则。该节点是否能成为一个独立的uri, 是否自身就是一个终极节点
	segment  string              // uri中的字符串，代表这个节点表示的路由中某个段的字符串
	handlers []ControllerHandler // 中间件 + 控制器
	childs   []*node             // 代表这个节点下的子节点
}

// 判断一个segment是否是通配符
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// 过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}
	// 如果segment是通配符，则所有下一层子节点都满足需求
	if isWildSegment(segment) {
		return n.childs
	}

	nodes := make([]*node, 0, len(n.childs))
	// 过滤所有的下一层子节点
	for _, cnode := range n.childs {
		// 如果当前节点是通配符，则满足需求，添加到结果中
		if isWildSegment(cnode.segment) {
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment {
			nodes = append(nodes, cnode)
		}
	}
	return nodes
}

func (n *node) mathNode(uri string) *node {
	// 使用分隔符将uri切割， 切割一次， 得到第一部分和第二部分
	segments := strings.SplitN(uri, "/", 2)
	// 第一部分用于匹配下一个节点
	segment := segments[0]
	if !isWildSegment(segment) {
		// 过滤所有满足segment的子节点
		segment = strings.ToUpper(segment)
	}
	// 匹配符合的下一层子节点
	cnodes := n.filterChildNodes(segment)
	// 如果当前子节点没有一个符合，那么说明这个uri一定是之前不存在, 直接返回nil
	if len(cnodes) == 0 || cnodes == nil {
		return nil
	}

	// 如果只有一个 segment, 说明uri已经是最后一个节点了
	if len(segments) == 1 {
		// 说明uri已经是最后一个节点了
		for _, tn := range cnodes {
			if tn.isLast {
				return tn
			}
		}
		// 都不是最后一个节点
		return nil
	}

	// 如果有2个segment,说明uri还没有结束,递归每个子节点继续进行查找
	for _, tn := range cnodes {
		tnMatch := tn.mathNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}

	return nil
}

func Newnode() *node {
	return &node{
		isLast:   false,
		segment:  "",
		handlers: []ControllerHandler{},
		childs:   []*node{},
	}
}

// newTree
func NewTree() *Tree {
	root := Newnode()
	return &Tree{root}
}

// 增加路由节点
func (tree *Tree) AddRouter(uri string, handlers []ControllerHandler) error {
	n := tree.root

	// 1. 判断是否冲突
	if n.mathNode(uri) != nil {
		return errors.New("router exist:" + uri)
	}
	// 2. 增加路由规则
	// 将uri使用 ‘/’ 切割成多个字符串
	segments := strings.Split(uri, "/")
	// 逐个遍历每个字符串
	for index, segment := range segments {
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1

		var objNode *node // 标记是否有合适的子节点

		// 遍历当前节点的所有子节点，看是否有合适的子节点
		childNodes := n.filterChildNodes(segment)
		if len(childNodes) > 0 {
			for _, conde := range childNodes {
				if conde.segment == segment {
					objNode = conde
					break
				}

			}
		}

		// 如果没有合适的子节点，则创建一个新的节点
		if objNode == nil {
			cnode := Newnode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handlers = handlers
			}
			n.childs = append(n.childs, cnode)
			objNode = cnode
		}
		n = objNode
	}
	return nil
}

// 匹配 uri 的 handler
func (tree *Tree) FindHandler(uri string) []ControllerHandler {
	// 直接复用matchNode函数，uri是不带通配符的地址
	mathNode := tree.root.mathNode(uri)

	if mathNode == nil {
		return nil
	}
	return mathNode.handlers
}
