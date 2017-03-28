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
	fmt.Printf("service size := %d\n", len(nodes))
	if len(nodes) == 0 {
		return nil, NewNodeException(NO_NODE_SERVICE, fmt.Sprintf("none  node for ID:=%s,version:=%s\n", ID, version))
	}
	var n *node
	var err error
	var transport thrift.TTransport
	for retry := 0; retry < MAX_TRY_NUM; retry++ {
		n, err = s.lb.getNode(ID, nodes)
		if err == nil {
			//获取连接
			transport, err = s.provider.getConnection(n, timeOut)
			if err == nil {
				break
			}
		}
	}
	if err != nil {
		s.checkNode(ID, version, n, 1)
		return nil, err
	}
	ct := NewCTransport(n, transport)
	return ct, nil
}

func (s *serviceRouterCommonImpl) returnConnection(ct *ctransport) error {
	//TODO 相关节点错误计数清0
	return s.provider.returnConnection(ct)
}
func (s *serviceRouterCommonImpl) errorHandler(ID, version string, err thrift.TTransportException, ct *ctransport) {
	s.provider.distoryConnection(ct)
	errSep := 1
	switch err.TypeId() {
	case thrift.TIMED_OUT:
		errSep = 2
	case thrift.END_OF_FILE:
		errSep = 2
	case thrift.UNKNOWN_TRANSPORT_EXCEPTION:
		errSep = 1
	}
	s.checkNode(ID, version, ct.n, errSep)
}

//异常累加权重
func (s *serviceRouterCommonImpl) checkNode(ID, version string, n *node, errSep int) error {
	//TODO增加到错误node，map，累加错误技术
	//记录节点，定时器检测节点可用性，超过一定检测次数，设置节点不健康，clear node连接
	//s.provider.clearConnection(ct)
	//n.SetHealthy(false)
	return nil
}
