package thriftclient

type registerSimpleImpl struct {
	serviceMap map[string]*service
}

func NewRegisterSimpleImpl() *registerSimpleImpl {
	registerImpl := &registerSimpleImpl{serviceMap: make(map[string]*service)}

	//模拟数据
	ID := "userService"
	version := "1.0"
	service := NewService(ID, version)
	registerImpl.serviceMap[ID+"_"+version] = service

	node := NewNode1("192.168.1.100", 8080)
	node.SetDisable(false)
	node.SetHealthy(true)
	node.SetPriority(1)
	registerImpl.RegisterNode(service, node)
	node1 := NewNode1("192.168.1.101", 8080)
	node1.SetDisable(false)
	node1.SetHealthy(true)
	node1.SetPriority(2) //备机
	registerImpl.RegisterNode(service, node1)
	return registerImpl
}
func (r *registerSimpleImpl) RegisterNode(s *service, node *node) {
	s.addNode(node)
}
func (r *registerSimpleImpl) QueryService(ID, version string) *service {
	return r.serviceMap[ID+"_"+version]
}
