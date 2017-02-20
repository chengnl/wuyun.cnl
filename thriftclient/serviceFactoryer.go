package thriftclient

import (
	"git.apache.org/thrift.git/lib/go/thrift"
)

type serviceFactoryer interface {
	createService(ID, version string, timeOut int64) (*serviceProxy, error)
	genClient(ID, version string, t thrift.TTransport, f thrift.TProtocolFactory) interface{}
}
