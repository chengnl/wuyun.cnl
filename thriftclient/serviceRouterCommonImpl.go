package thriftclient

import (
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

const MAX_TRY_NUM = 2

type serviceRouterCommonImpl struct {
	nl       nodeLoader
	lb       loadBalancer
	provider connectionProvider
}

func NewServiceRouterCommonImpl() *serviceRouterCommonImpl {
	return &serviceRouterCommonImpl{nl: NewNodeLoaderImpl(), lb: NewLoadBalancerRoundRobinImpl(), provider: NewTransportConnectionProvider()}
}
func (s *serviceRouterCommonImpl) routeService(ID, version string, timeOut int64) (*ctransport, error) {
	nodes := s.nl.load(ID, version)
	fmt.Printf("service size := %d/n", len(nodes))
	if len(nodes) == 0 {
		return nil, fmt.Errorf("none  node for ID:=%s,version:=%s/n", ID, version)
	}
	var n *node
	var err error
	var transport thrift.TTransport
	for retry := 0; retry < MAX_TRY_NUM; {
		n, err = s.lb.getNode(ID, nodes)
		if err == nil {
			//获取连接
			transport, err = s.provider.getConnection(n, timeOut)
			if err == nil {
				break
			}
		}
		if retry < MAX_TRY_NUM {
			retry++
		} else {
			return nil, err
		}
	}
	ct := NewCTransport(n, transport)
	return ct, nil
}

func (s *serviceRouterCommonImpl) getConnectionProvider() connectionProvider {
	return s.provider
}
