package thriftclient

import "fmt"

const (
//DEMO_SERVICE = [2]string{"demoService", "1.0"}
)

type serviceFactory struct {
	router serviceRouter
}

func NewServiceFactory() *serviceFactory {
	return &serviceFactory{router: NewServiceRouterCommonImpl()}
}
func (factory *serviceFactory) getDemoService() (*serviceProxy, error) {
	obj, err := factory.createService(demo.TestService, 5000)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (factory *serviceFactory) createService(obj interface{}, timeOut int64) (*serviceProxy, error) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	switch obj.(type) {
	case demo.TestService:
		var client *demo.TestServiceClient
		ct, err := factory.router.routeService(DEMO_SERVICE[0], DEMO_SERVICE[1], timeOut)
		if err != nil {
			return nil, err
		}
		client = demo.NewTestServiceClientFactory(ct.transport, protocolFactory)
		proxy := NewServiceProxy(ct, factory.router)
		return proxy, nil
	default:
		return nil, fmt.Errorf("serviceName=%s not found\n", serviceName)
	}
}
