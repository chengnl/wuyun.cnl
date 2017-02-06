package thriftclient

type registerImpl struct {
}

func init() {

}
func NewRegisterImpl() *registerImpl {
	return &registerImpl{}
}
func (r *registerImpl) registerService(s service, node node) {

}
func (r *registerImpl) queryService(id, version string) *service {
	return nil
}
