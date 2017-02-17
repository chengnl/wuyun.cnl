package thriftclient

import (
	"errors"
	"fmt"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type transportFactory struct {
	serviceIP   string
	servicePort int
	timeOut     int64
}

func NewTransportFactory(serviceIP string, servicePort int, timeOut int64) *transportFactory {
	return &transportFactory{serviceIP: serviceIP, servicePort: servicePort, timeOut: timeOut}
}

//创建对象
func (tf *transportFactory) CreateObj() (interface{}, error) {
	transportFactory := thrift.NewTTransportFactory()
	var err error
	var transport thrift.TTransport
	addr := fmt.Sprintf("%s:%d", tf.serviceIP, tf.servicePort)
	transport, err = thrift.NewTSocketTimeout(addr, time.Duration(tf.timeOut)*time.Millisecond)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return nil, err
	}
	transport = transportFactory.GetTransport(transport)
	if err := transport.Open(); err != nil {
		fmt.Println("Error transport opening:", err)
		return nil, err
	}
	return transport, nil
}

//销毁对象
func (tf *transportFactory) DestroyObj(t interface{}) error {
	value, ok := t.(thrift.TTransport)
	if ok {
		if value.IsOpen() {
			value.Close()
		}
	}
	return nil

}

//验证对象
func (tf *transportFactory) ValidateObj(t interface{}) error {
	value, ok := t.(thrift.TTransport)
	if ok {
		if value.IsOpen() {
			return nil
		}
	}
	return errors.New("transport is not validate")
}
