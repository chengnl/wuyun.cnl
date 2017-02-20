package thriftclient

import (
	"git.apache.org/thrift.git/lib/go/thrift"
)

type serviceFactoryer interface {
	CreateService(ID, version string, timeOut int64) *ServiceProxy
	GenClient(ID, version string, t thrift.TTransport, f thrift.TProtocolFactory) interface{}
}
