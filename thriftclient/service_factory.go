package thriftclient

type serviceFactory struct {
}

func NewServiceFactory() *serviceFactory {
	return &serviceFactory{}
}
