package thriftclient

type node struct {
	host     string
	port     int
	disabled bool
	healthy  bool
}

func NewNode() *node {
	return &node{disabled: true, healthy: false}
}
func NewNode1(host string, port int) *node {
	return &node{host: host, port: port, disabled: true, healthy: false}
}
func (n *node) getHost() string {
	return n.host
}
func (n *node) getPort() int {
	return n.port
}
func (n *node) getDisable() bool {
	return n.disabled
}
func (n *node) setDisable(disabled bool) {
	n.disabled = disabled
}
func (n *node) getHealthy() bool {
	return n.healthy
}
func (n *node) setHealthy(healthy bool) {
	n.healthy = healthy
}
