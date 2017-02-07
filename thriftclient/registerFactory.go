package thriftclient

type registerFatcory struct {
	r register
}

func NewRegisterFacory() *registerFatcory {
	return &registerFatcory{r: NewRegisterSimpleImpl()}
}
func (factory *registerFatcory) GetRegister() register {
	return factory.r
}
