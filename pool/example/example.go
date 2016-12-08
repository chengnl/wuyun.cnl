package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/chengnl/wuyun.cnl/pool"
)

type PoolObjectFactory struct{}

func NewFactory() *PoolObjectFactory {
	return &PoolObjectFactory{}
}
func (f *PoolObjectFactory) CreateObj() (interface{}, error) {
	return net.Dial("tcp", "127.0.0.1:8080")
}

func (f *PoolObjectFactory) DestroyObj(c interface{}) error {
	return c.(net.Conn).Close()
}
func (f *PoolObjectFactory) ValidateObj(c interface{}) error {
	fmt.Println("OBJ valid")
	return nil
}

func init() {
	go func() {
		l, err := net.Listen("tcp", "127.0.0.1:8080")
		if err != nil {
			log.Fatal(err)
		}
		defer l.Close()

		for {
			conn, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}

			go func() {
				buffer := make([]byte, 256)
				conn.Read(buffer)
				fmt.Println(string(buffer))
			}()
		}
	}()
}

func main() {
	config := &pool.PoolConfig{MAX_ACTIVE: 100, MAX_IDLE: 50, MIN_IDLE: 5, MAX_WAIT: 60 * time.Millisecond, IsVaildObj: true}
	p, err := pool.NewGenObjectPool(NewFactory(), config)
	if err != nil {
		fmt.Println(err)
	}
	p.InitPool()
	fmt.Printf("activeNum:%d;idleNum:%d\n", p.GetNumActive(), p.GetNumIdle())
	poolObject, err := p.GetObject()
	if err != nil {
		fmt.Println(err)
	} else {
		poolObject.(net.Conn).Write([]byte("hello"))
		p.ReturnObject(poolObject)
		//fmt.Printf("activeNum:%d;idleNum:%d\n", p.GetNumActive(), p.GetNumIdle())
		//p.DestroyObject(poolObject)
		//p.Clear()
	}
}
