package thriftclient

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"reflect"
)

type ServiceProxy struct {
	ID, version     string
	timeOut         int64
	genClient       func(thrift.TTransport, thrift.TProtocolFactory) interface{}
	router          serviceRouter
	protocolFactory *thrift.TBinaryProtocolFactory
	ct              *ctransport
}

func NewServiceProxy(ID, version string, timeOut int64, genClient func(thrift.TTransport, thrift.TProtocolFactory) interface{}, router serviceRouter) *ServiceProxy {
	return &ServiceProxy{ID: ID, version: version, timeOut: timeOut, genClient: genClient, router: router,
		protocolFactory: thrift.NewTBinaryProtocolFactoryDefault()}
}

func (s *ServiceProxy) Call(methodName string, params ...interface{}) (result interface{}, err error) {
	ct, err := s.router.routeService(s.ID, s.version, s.timeOut)
	if err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error:%v \n", p)
			s.router.returnConnection(ct)
		}
	}()
	s.ct = ct
	client := s.genClient(ct.transport, s.protocolFactory)
	v := reflect.ValueOf(client)
	f := v.MethodByName(methodName)
	if f.IsValid() {
		return nil, NewMethodException(NO_SUCH_METHOD, fmt.Sprintf("methodName=%s  is not valid\n", methodName))
	}
	pl := len(params)
	if pl != f.Type().NumIn() {
		s.router.returnConnection(ct)
		return nil, NewMethodException(NO_MACH_PARAMS, fmt.Sprintf("params is not match method=%s \n", methodName))
	}
	var in []reflect.Value
	if pl > 0 {
		in = make([]reflect.Value, pl)
		for k, param := range params {
			in[k] = reflect.ValueOf(param)
		}
	}
	out := f.Call(in)
	ol := len(out)
	switch ol {
	case 1:
		if out[0].Interface() != nil {
			err = out[0].Interface().(error)
		}
		break
	case 2:
		result = out[0].Interface()
		if out[1].Interface() != nil {
			err = out[1].Interface().(error)
		}
		break
	default:
		s.router.returnConnection(ct)
		return nil, NewMethodException(NO_MACH_RESULT, fmt.Sprintf("call result is not match method=%s \n", methodName))
	}
	if e, ok := err.(thrift.TTransportException); ok {
		s.router.errorHandler(s.ID, s.version, e, ct)
	} else {
		s.router.returnConnection(ct)
	}
	return
}
