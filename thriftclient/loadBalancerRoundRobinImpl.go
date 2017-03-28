package thriftclient

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)

type loadBalancerRoundRobinImpl struct {
	counterMap map[string]*int32
	mux        *sync.Mutex
}

func NewLoadBalancerRoundRobinImpl() *loadBalancerRoundRobinImpl {
	return &loadBalancerRoundRobinImpl{counterMap: make(map[string]*int32), mux: new(sync.Mutex)}
}

func (l *loadBalancerRoundRobinImpl) getNode(serviceID string, nodes []*node) (*node, error) {
	counter, ok := l.counterMap[serviceID]
	if !ok {
		l.mux.Lock()
		defer l.mux.Unlock()
		if counter, ok = l.counterMap[serviceID]; !ok {
			c := int32(0)
			counter = &c
			l.counterMap[serviceID] = counter
		}
		atomic.AddInt32(l.counterMap[serviceID], 1)
	}
	if *counter > (math.MaxInt32 - 1000) {
		atomic.StoreInt32(counter, int32(0))
	}
	for step := 0; step < len(nodes); step++ {
		node := nodes[*counter%int32(len(nodes))]
		if !node.GetDisable() && node.GetHealthy() {
			return node, nil
		}
		fmt.Printf("node:%v 不可用/n", node)
	}
	return nil, NewNodeException(NO_AVAILABLE_NODE, fmt.Sprintf("none available node for service:%s", serviceID))
}
