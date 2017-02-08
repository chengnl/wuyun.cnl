package thriftclient

type nodeLoaderImpl struct {
	rf *registerFatcory
}

func NewNodeLoaderImpl() *nodeLoaderImpl {
	return &nodeLoaderImpl{rf: NewRegisterFacory()}
}
func (load *nodeLoaderImpl) load(ID, version string) []*node {
	return load.rf.GetRegister().QueryService(ID, version).getNodes()
}
