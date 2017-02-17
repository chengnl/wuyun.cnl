package thriftclient

import "github.com/chengnl/wuyun.cnl/pool"
import "sync"
import "git.apache.org/thrift.git/lib/go/thrift"
import "time"
import "fmt"

type transportConnectionProvider struct {
	poolMap map[string]pool.ObjectPooler
	mux     *sync.Mutex
	//最大空闲对象数量
	maxIdle int
	//最小空闲对象数量
	minIdle int
	//最大存活对象数量
	maxActive int
	//获取对象最大等待时间，<=0 一直等待
	maxWait time.Duration
	//是否在获取对象的时候验证对象有效性
	isVaildObj bool
}

func NewTransportConnectionProvider() *transportConnectionProvider {
	return &transportConnectionProvider{poolMap: make(map[string]pool.ObjectPooler), mux: new(sync.Mutex),
		maxIdle: 50, minIdle: 5, maxActive: 100, maxWait: 10 * time.Millisecond, isVaildObj: false}
}
func (provider *transportConnectionProvider) setParams(maxIdle, minIdle, maxActive int, maxWait time.Duration) {
	provider.maxActive = maxActive
	provider.maxIdle = maxIdle
	provider.minIdle = minIdle
	provider.maxWait = maxWait
}
func (provider *transportConnectionProvider) getConnection(node *node, timeOut int64) (thrift.TTransport, error) {
	var err error
	key := node.getNodeKey()
	pool, ok := provider.poolMap[key]
	if !ok {
		provider.mux.Lock()
		defer provider.mux.Unlock()
		pool, ok = provider.poolMap[key]
		if !ok {
			//创建连接池
			pool, err = provider.createPool(node, timeOut)
			if err != nil {
				return nil, err
			}
			provider.poolMap[key] = pool
		}
	}
	transport, e := pool.GetObject()
	if e != nil {
		return nil, err
	}
	return transport.(thrift.TTransport), nil
}
func (provider *transportConnectionProvider) returnConnection(ct *ctransport) error {
	key := ct.n.getNodeKey()
	transport := ct.transport
	pool, ok := provider.poolMap[key]
	if ok {
		if transport != nil {
			pool.ReturnObject(transport)
		}
	} else {
		fmt.Printf("returnConnection key:=%s ,pool not exist/n", key)
	}
	return nil
}
func (provider *transportConnectionProvider) distoryConnection(ct *ctransport) error {
	key := ct.n.getNodeKey()
	transport := ct.transport
	pool, ok := provider.poolMap[key]
	if ok {
		if transport != nil {
			pool.DestroyObject(transport)
		}
	} else {
		fmt.Printf("distoryConnection key:=%s ,pool not exist/n", key)
	}
	return nil
}
func (provider *transportConnectionProvider) clearConnection(node *node) error {
	pool, ok := provider.poolMap[node.getNodeKey()]
	if ok {
		pool.Clear()
	} else {
		fmt.Printf("clearConnection key:=%s ,pool not exist/n", node.getNodeKey())
	}
	return nil
}

func (provider *transportConnectionProvider) createPool(node *node, timeOut int64) (pool.ObjectPooler, error) {
	tf := NewTransportFactory(node.GetHost(), node.GetPort(), timeOut)
	config := pool.NewPoolConfig(provider.maxIdle, provider.minIdle, provider.maxActive, provider.maxWait,
		provider.isVaildObj)
	pool, err := pool.NewGenObjectPool(tf, config)
	if err != nil {
		return nil, err
	}
	pool.InitPool()
	return pool, nil
}
