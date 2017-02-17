package thriftclient

type thriftclient interface {
	routeService(ID, version string, timeOut int64) (*ctransport, error)
}
