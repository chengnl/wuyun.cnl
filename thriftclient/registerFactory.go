package thriftclient

type registerFatcory struct {
	r register
}

var RegisterFactory *registerFatcory

func init() {
	RegisterFactory = NewRegisterFacory()
}

func NewRegisterFacory() *registerFatcory {
	return &registerFatcory{r: NewRegisterImpl()}
}
func (factory *registerFatcory) getRegister() register {
	return factory.r
}
