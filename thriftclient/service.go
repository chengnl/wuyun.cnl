package thriftclient

import "container/list"

type service struct {
	ID           string
	Version      string
	nodes        *list.List
	userPriority int8 //主机选择策略 1：使用优先级，优先使用主机，如果有问题才使用备机  2：使用备机,3:都使用
}

func NewService(ID, version string) *service {
	return &service{ID: ID, Version: version, nodes: list.New(), userPriority: 1}
}
func (s *service) addNode(node *node) {
	s.nodes.PushBack(node)
}
func (s *service) getNodes() *list.List {
	switch s.userPriority {
	case 1:
		nodes := list.New()
		for e := s.nodes.Front(); e != nil; e = e.Next() {
			if ne, ok := e.Value.(node); ok {
				if ne.GetPriority() == 1 {
					nodes.PushBack(ne)
				}
			}
		}
		if nodes.Len() == 0 {
			return s.nodes
		}
		return nodes
	case 2:
		nodes := list.New()
		for e := s.nodes.Front(); e != nil; e = e.Next() {
			if ne, ok := e.Value.(node); ok {
				if ne.GetPriority() == 2 {
					nodes.PushBack(ne)
				}
			}
		}
		return nodes
	case 3:
		return s.nodes
	default:
		return nil
	}
}
