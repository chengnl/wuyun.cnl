package thriftclient

type registerFatcory struct {
	r register
}

//可以修改为自己的节点注册实现
func NewRegisterFacory() *registerFatcory {
	return &registerFatcory{r: NewRegisterSimpleImpl()}
}
func (factory *registerFatcory) GetRegister() register {
	return factory.r
}
