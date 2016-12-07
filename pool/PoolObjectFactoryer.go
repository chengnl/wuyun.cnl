package pool

//创建对象工厂类接口
type PoolObjectFactoryer interface {
	//创建对象
	CreateObj() (interface{}, error)
	//销毁对象
	DestroyObj(interface{}) error
	//验证对象
	ValidateObj(interface{}) error
}
