package thriftclient

type serviceProxy struct {
	client   interface{}
	ct       *ctransport
	provider connectionProvider
}

func NewServiceProxy(client interface{}, ct *ctransport, provider connectionProvider) *serviceProxy {
	return &serviceProxy{client: client, ct: ct, provider: provider}
}
func (s *serviceProxy) returnConnection() error {
	return s.provider.returnConnection(s.ct)
}
func (s *serviceProxy) distoryConnection() error {
	return s.provider.distoryConnection(s.ct)
}
func (s *serviceProxy) clearConnection() error {
	return s.provider.clearConnection(s.ct.n)
}
