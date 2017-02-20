package thriftclient

import (
	"git.apache.org/thrift.git/lib/go/thrift"
)

type serviceFactory struct {
	router    serviceRouter
	factoryer serviceFactoryer
}

func NewServiceFactory(factoryer serviceFactoryer) *serviceFactory {
	return &serviceFactory{router: NewServiceRouterCommonImpl(), factoryer: factoryer}
}
func (factory *serviceFactory) createService(ID, version string, timeOut int64) (*serviceProxy, error) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	ct, err := factory.router.routeService(ID, version, timeOut)
	if err != nil {
		return nil, err
	}
	client := factory.factoryer.genClient(ID, version, ct.transport, protocolFactory)
	proxy := NewServiceProxy(client, ct, factory.router.getConnectionProvider())
	return proxy, nil
}
