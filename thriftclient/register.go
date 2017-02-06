package thriftclient

type register interface {
	registerService(s service, node node)
	queryService(id, version string) *service
}
