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
func (factory *ServiceFactory) CreateService(factoryer serviceFactoryer, ID, version string, timeOut int64) *ServiceProxy {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	ct, err := factory.router.routeService(ID, version, timeOut)
	if err != nil {
		return nil
	}
	client := factoryer.GenClient(ID, version, ct.transport, protocolFactory)
	if client == nil {
		return nil
	}
	proxy := NewServiceProxy(client, ct, factory.router.getConnectionProvider())
	return proxy
}
