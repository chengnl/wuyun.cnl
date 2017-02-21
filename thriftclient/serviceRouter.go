package thriftclient

import (
	"git.apache.org/thrift.git/lib/go/thrift"
)

type serviceRouter interface {
	routeService(ID, version string, timeOut int64) (*ctransport, error)
	returnConnection(ct *ctransport) error
	errorHandler(ID, version string, err thrift.TTransportException, ct *ctransport)
}
