package thriftclient

import "fmt"

type node struct {
	host     string
	port     int
	disabled bool
	healthy  bool
	priority int8
}

func NewNode() *node {
	return &node{disabled: true, healthy: false}
}
func NewNode1(host string, port int) *node {
	return &node{host: host, port: port, disabled: true, healthy: false}
}
func (n *node) GetHost() string {
	return n.host
}
func (n *node) GetPort() int {
	return n.port
}
func (n *node) GetDisable() bool {
	return n.disabled
}
func (n *node) SetDisable(disabled bool) {
	n.disabled = disabled
}
func (n *node) GetHealthy() bool {
	return n.healthy
}
func (n *node) SetHealthy(healthy bool) {
	n.healthy = healthy
}
func (n *node) SetPriority(priority int8) {
	n.priority = priority
}
func (n *node) GetPriority() int8 {
	return n.priority
}
func (n *node) getNodeKey() string {
	return fmt.Sprintf("%s:%d", n.host, n.port)
}
