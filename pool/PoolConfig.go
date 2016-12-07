package pool

import "time"

//对象池配置信息
type PoolConfig struct {
	//最大空闲对象数量
	MAX_IDLE int
	//最小空闲对象数量
	MIN_IDLE int
	//最大存活对象数量
	MAX_ACTIVE int
	//获取对象最大等待时间，<=0 一直等待
	MAX_WAIT time.Duration
	//是否在获取对象的时候验证对象有效性
	IsVaildObj bool
}
