package main

import "strings"

type HandlerBasedOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node
	handler  handlerFunc
}

func (n *HandlerBasedOnTree) ServeHTTP(c *Context) {

}

func (n *HandlerBasedOnTree) Route(pattern string, handlerFunc func(ctx *Context)) {
	pattern = strings.Trim(pattern, "/")
	//[user,friends]
	paths := strings.Split(pattern, "/")
	cur := n.root
	for _, path := range paths {
		matchChild, ok := n.findMatchChild(cur, path)
		if ok {
			cur = matchChild
		}
	}

}

//这个应该是 node 知道，而不是用我的人知道怎么找
func (n *HandlerBasedOnTree) findMatchChild(root *node, path string) (*node, bool) {
	for _, child := range root.children {
		if child.path == path {
			return child, true
		}
	}
	return nil, false
}
