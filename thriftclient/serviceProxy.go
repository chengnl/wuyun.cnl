package thriftclient

type ServiceProxy struct {
	client   interface{}
	ct       *ctransport
	provider connectionProvider
}

func NewServiceProxy(client interface{}, ct *ctransport, provider connectionProvider) *ServiceProxy {
	return &ServiceProxy{client: client, ct: ct, provider: provider}
}
func (s *ServiceProxy) returnConnection() error {
	return s.provider.returnConnection(s.ct)
}
func (s *ServiceProxy) distoryConnection() error {
	return s.provider.distoryConnection(s.ct)
}
func (s *ServiceProxy) clearConnection() error {
	return s.provider.clearConnection(s.ct.n)
}
