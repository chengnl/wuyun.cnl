package thriftclient

import (
	"git.apache.org/thrift.git/lib/go/thrift"
)

type ServiceProxy struct {
	client   interface{}
	ct       *ctransport
	provider connectionProvider
}

func NewServiceProxy(client interface{}, ct *ctransport, provider connectionProvider) *ServiceProxy {
	return &ServiceProxy{client: client, ct: ct, provider: provider}
}
func (s *ServiceProxy) ReturnConnection() error {
	return s.provider.returnConnection(s.ct)
}
func (s *ServiceProxy) DistoryConnection() error {
	return s.provider.distoryConnection(s.ct)
}
func (s *ServiceProxy) ClearConnection() error {
	return s.provider.clearConnection(s.ct.n)
}
func (s *ServiceProxy) GetClient() interface{} {
	return s.client
}

//先简单处理下错误，后续增加节点检测机制
func (s *ServiceProxy) HandlerError(err error) {
	if _, ok := err.(thrift.TTransportException); ok {
		s.DistoryConnection()
	} else {
		s.ReturnConnection()
	}
}
