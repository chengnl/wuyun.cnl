package thriftclient

type serviceRouter interface {
	routeService(ID, version string, timeOut int64) (*ctransport, error)
}
