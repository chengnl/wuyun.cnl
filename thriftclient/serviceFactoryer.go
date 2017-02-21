package thriftclient

import (
	"git.apache.org/thrift.git/lib/go/thrift"
)

type serviceFactoryer interface {
	CreateService(ID, version string, timeOut int64, genClient func(thrift.TTransport, thrift.TProtocolFactory) interface{}) (*ServiceProxy, error)
}
