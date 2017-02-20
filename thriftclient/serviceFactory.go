package thriftclient

import (
	"git.apache.org/thrift.git/lib/go/thrift"
)

type ServiceFactory struct {
	router serviceRouter
}

func NewServiceFactory() *ServiceFactory {
	return &ServiceFactory{router: NewServiceRouterCommonImpl()}
}
func (factory *ServiceFactory) CreateService(factoryer serviceFactoryer, ID, version string, timeOut int64) (*ServiceProxy, error) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	ct, err := factory.router.routeService(ID, version, timeOut)
	if err != nil {
		return nil, err
	}
	client := factoryer.GenClient(ID, version, ct.transport, protocolFactory)
	proxy := NewServiceProxy(client, ct, factory.router.getConnectionProvider())
	return proxy, nil
}
