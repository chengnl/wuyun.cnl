package thriftclient

import (
	"container/list"
	"time"
)

type registerImpl struct {
	serviceMap map[string]service
}

var services map[string]*list.List

func init() {
	//模拟数据
	services = make(map[string]*list.List)
	serviceInfo1 := "userService_1.0"
	nodeList := list.New()
	node := NewNode1("192.168.1.100", 8080)
	node.SetDisable(false)
	node.SetHealthy(true)
	nodeList.PushBack(node)

	node1 := NewNode1("192.168.1.101", 8080)
	node1.SetDisable(false)
	node1.SetHealthy(true)
	nodeList.PushBack(node1)
	services[serviceInfo1] = nodeList
}

func NewRegisterImpl() *registerImpl {
	registerImpl := &registerImpl{serviceMap: make(map[string]service)}
	registerImpl.checkNodeState()
	return registerImpl
}
func (r *registerImpl) RegisterNode(s *service, node *node) {

}
func (r *registerImpl) QueryService(id, version string) *service {
	return nil
}
func (r *registerImpl) checkNodeState() {
	go func() {
		for {
			//每隔1s定期检测节点状态
			time.AfterFunc(1*time.Second, func() {
				for _, service := range r.serviceMap {
					id := service.ID
					version := service.Version
					nodes := getService(id, version)

				}
			})
		}
	}()
}

//模拟服务节点配置，查询返回服务节点信息
func getService(ID, version string) *list.List {
	return services[ID+"_"+version]
}
