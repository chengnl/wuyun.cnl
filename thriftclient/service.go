package thriftclient

type service struct {
	ID           string
	Version      string
	nodes        []*node
	userPriority int8 //主机选择策略 1：使用优先级，优先使用主机，如果有问题才使用备机  2：使用备机,3:都使用
}

func NewService(ID, version string) *service {
	return &service{ID: ID, Version: version, nodes: make([]*node, 0), userPriority: 1}
}
func (s *service) addNode(node *node) {
	s.nodes = append(s.nodes, node)
}
func (s *service) getNodes() []*node {
	switch s.userPriority {
	case 1:
		nodes := make([]*node, 0)
		for _, ne := range s.nodes {
			if ne.GetPriority() == 1 {
				nodes = append(nodes, ne)
			}
		}
		if len(nodes) == 0 {
			return s.nodes
		}
		return nodes
	case 2:
		nodes := make([]*node, 0)
		for _, ne := range s.nodes {
			if ne.GetPriority() == 2 {
				nodes = append(nodes, ne)
			}
		}
		return nodes
	case 3:
		return s.nodes
	default:
		return nil
	}
}
