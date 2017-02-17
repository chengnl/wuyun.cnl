package thriftclient

import "github.com/chengnl/wuyun.cnl/pool"
import "sync"
import "git.apache.org/thrift.git/lib/go/thrift"
import "time"

type transportConnectionProvider struct {
	poolMap map[string]pool.ObjectPooler
	mux     *sync.Mutex
}

func NewTransportConnectionProvider() *transportConnectionProvider {
	return &transportConnectionProvider{poolMap: make(map[string]pool.ObjectPooler), mux: new(sync.Mutex)}
}

func (provider *transportConnectionProvider) getConnection(node *node, timeOut int64) (thrift.TTransport, error) {
	key := node.getNodeKey()
	pool, ok := provider.poolMap[key]
	if !ok {
		provider.mux.Lock()
		defer provider.mux.Unlock()
		pool, ok = provider.poolMap[key]
		if !ok {
			//创建连接池
		}
	}
	return nil, nil
}
func (provider *transportConnectionProvider) returnConnection(ct *ctransport) error {
	return nil
}
func (provider *transportConnectionProvider) distoryConnection(ct *ctransport) error {
	return nil
}
func (provider *transportConnectionProvider) clearConnection(node *node) error {
	return nil
}

func (provider *transportConnectionProvider) createPool(node *node, timeOut int64) (pool.ObjectPooler, error) {
	tf := NewTransportFactory(node.GetHost(), node.GetPort(), timeOut)
	config := pool.n
	config := &pool.PoolConfig{pool.PoolConfig.MAX_ACTIVE: 100, pool.PoolConfig.MAX_IDLE: 50,
		pool.PoolConfig.MIN_IDLE: 5, pool.PoolConfig.MAX_WAIT: 60 * time.Millisecond, pool.PoolConfig.IsVaildObj: false}
	pool, err := pool.NewGenObjectPool(tf, config)
	if err != nil {
		return nil, err
	}
	pool.InitPool()
	return pool, nil
}
