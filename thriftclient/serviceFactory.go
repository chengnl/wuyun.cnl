package thriftclient

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"sync"
)

type ServiceFactory struct {
	router serviceRouter
}

var sf *ServiceFactory
var once sync.Once

func NewServiceFactory() *ServiceFactory {
	once.Do(func() {
		sf = &ServiceFactory{router: NewServiceRouterCommonImpl()}
	})
	return sf
}
func (factory *ServiceFactory) CreateService(ID, version string, timeOut int64, genClient func(thrift.TTransport, thrift.TProtocolFactory) interface{}) *ServiceProxy {
	proxy := NewServiceProxy(ID, version, timeOut, genClient, factory.router)
	return proxy
}
