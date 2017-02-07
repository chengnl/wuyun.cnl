package thriftclient

import "container/list"

type nodeLoader interface {
	//加载服务相关所有节点
	load(ID, version string) *list.List
}
