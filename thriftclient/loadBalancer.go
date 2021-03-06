package thriftclient

//负载均衡服务节点接口
type loadBalancer interface {
	//返回服务一个节点
	getNode(serviceID string, nodes []*node) (*node, error)
}
