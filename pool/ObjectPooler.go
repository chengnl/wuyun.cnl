package pool

//对象池接口
type ObjectPooler interface {
	//初始化对象池
	InitPool() error
	//连接池中获取对象
	GetObject() (interface{}, error)
	//将对象返回连接池
	ReturnObject(interface{}) error
	//销毁对象池中对象
	DestroyObject(interface{}) error
	//清除对象池
	Clear()
	//获取空闲对象数
	GetNumIdle() int
	//获取正在活动的对象数
	GetNumActive() int
}
