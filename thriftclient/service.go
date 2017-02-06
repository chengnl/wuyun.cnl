package thriftclient

import "container/list"

type service struct {
	id      string
	version string
	nodes   *list.List
}

func NewService(id, version string) *service {
	return &service{id: id, version: version, nodes: list.New()}
}
func (s *service) addNode(node node) {
	s.nodes.PushBack(node)
}
func (s *service) getNodes() *list.List {
	return s.nodes
}
