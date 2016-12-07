package pool

import (
	"errors"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var POOL_IS_FULL = "POOL_FULL"

type genObjectPool struct {
	config    *PoolConfig
	factory   PoolObjectFactoryer
	objNum    int64
	idlObject chan interface{}
	lock      sync.Mutex
}

func DefaultGenObjectPool(f PoolObjectFactoryer) (ObjectPooler, error) {
	DEFAULTNUM := 5
	param := &PoolConfig{MAX_IDLE: DEFAULTNUM, MIN_IDLE: 0, MAX_ACTIVE: DEFAULTNUM, MAX_WAIT: time.Duration(-1), IsVaildObj: false}
	return NewGenObjectPool(f, param)
}
func NewGenObjectPool(f PoolObjectFactoryer, param *PoolConfig) (ObjectPooler, error) {
	if param.MAX_IDLE > param.MAX_ACTIVE || param.MAX_ACTIVE <= 0 || param.MAX_ACTIVE > math.MaxInt32 {
		return nil, errors.New("invalid config,check MAX_ACTIVE or MAX_IDLE set.")
	}
	if param.MIN_IDLE > param.MAX_IDLE {
		return nil, errors.New("invalid config, must be set MIN_IDLE < MAX_IDLE.")
	}
	if f == nil {
		return nil, errors.New("factory is nil")
	}
	idleObj := make(chan interface{}, param.MAX_IDLE)
	return &genObjectPool{config: param, factory: f, objNum: 0, idlObject: idleObj}, nil
}
func (p *genObjectPool) InitPool() error {
	if p.config.MIN_IDLE > 0 {
		for i := 0; i < p.config.MIN_IDLE; i++ {
			obj, err := p.factory.CreateObj()
			if err != nil {
				return err
			}
			if obj == nil {
				return errors.New("factory create obj is nil")
			}
			atomic.AddInt64(&p.objNum, 1)
			p.idlObject <- obj
		}
	}
	return nil
}
func (p *genObjectPool) GetObject() (interface{}, error) {
	waitTime := p.config.MAX_WAIT
	startTime := time.Now()
	if waitTime <= 0 {
		waitTime = time.Duration(math.MaxInt64)
	}
	for {
		if p.idlObject == nil {
			return nil, errors.New("idlObject is nil,pool is Clear!")
		}
		select {
		case poolObj := <-p.idlObject:
			if p.config.IsVaildObj && !p.isValidObject(poolObj) {
				continue
			}
			return poolObj, nil
		default:
			poolObj, err := p.createObject()
			if poolObj == nil {
				if err != nil && err.Error() == POOL_IS_FULL {
					select {
					case <-time.After(waitTime):
						return nil, errors.New("Timeout waiting for idle object")
					case poolObj = <-p.idlObject:
						if p.config.IsVaildObj && !p.isValidObject(poolObj) {
							dtime := time.Since(startTime)
							waitTime = waitTime - dtime
							continue
						}
						return poolObj, nil
					}
				} else {
					return nil, err
				}
			} else {
				return poolObj, nil
			}
		}
	}

}
func (p *genObjectPool) ReturnObject(poolObj interface{}) error {
	if poolObj == nil {
		return errors.New("poolObj is nil")
	}
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.idlObject == nil {
		return p.DestroyObject(poolObj)
	}
	select {
	case p.idlObject <- poolObj:
		return nil
	default:
		return p.DestroyObject(poolObj)
	}
}
func (p *genObjectPool) DestroyObject(poolObj interface{}) error {
	if poolObj == nil {
		return errors.New("poolObj is nil")
	}
	defer atomic.AddInt64(&p.objNum, -1)
	err := p.factory.DestroyObj(poolObj)
	if err != nil {
		return err
	}
	return nil
}
func (p *genObjectPool) Clear() {
	p.lock.Lock()
	conns := p.idlObject
	p.idlObject = nil
	p.lock.Unlock()
	if conns == nil {
		return
	}
	close(conns)
	for poolObj := range conns {
		p.factory.DestroyObj(poolObj)
		atomic.AddInt64(&p.objNum, -1)
	}
}
func (p *genObjectPool) GetNumIdle() int {
	return len(p.idlObject)

}
func (p *genObjectPool) GetNumActive() int {
	return int(p.objNum) - len(p.idlObject)
}

func (p *genObjectPool) createObject() (interface{}, error) {
	num := atomic.AddInt64(&p.objNum, 1)
	if num > int64(p.config.MAX_ACTIVE) || num > int64(math.MaxInt32) {
		atomic.AddInt64(&p.objNum, -1)
		return nil, errors.New(POOL_IS_FULL)
	}
	poolObj, err := p.factory.CreateObj()
	if err != nil {
		atomic.AddInt64(&p.objNum, -1)
		return nil, err
	}
	if poolObj == nil {
		return nil, errors.New("factory create obj is nil")
	}
	return poolObj, nil
}
func (p *genObjectPool) isValidObject(poolObj interface{}) bool {
	if poolObj == nil {
		return false
	}
	err := p.factory.ValidateObj(poolObj)
	if err != nil {
		p.DestroyObject(poolObj)
		return false
	}
	return true
}
