package thriftclient

type register interface {
	RegisterNode(s *service, node *node)
	QueryService(id, version string) *service
}
