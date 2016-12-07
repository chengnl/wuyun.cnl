package pool

import (
	"fmt"
	"log"
	"net"
	"sync"
	"testing"
	"time"
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
func TestNew(t *testing.T) {
	pool, err := DefaultGenObjectPool(NewFactory())
	if err != nil {
		fmt.Println(err)
	}
	pool.InitPool()
	fmt.Printf("activeNum:%d;idleNum:%d\n", pool.GetNumActive(), pool.GetNumIdle())
}
func TestNew_config(t *testing.T) {
	config := &PoolConfig{MAX_ACTIVE: 100, MAX_IDLE: 50, MIN_IDLE: 5, MAX_WAIT: 250}
	pool, err := NewGenObjectPool(NewFactory(), config)
	if err != nil {
		fmt.Println(err)
	}
	pool.InitPool()
	fmt.Printf("activeNum:%d;idleNum:%d\n", pool.GetNumActive(), pool.GetNumIdle())
}
func TestCreate(t *testing.T) {
	pool, err := DefaultGenObjectPool(NewFactory())
	if err != nil {
		fmt.Println(err)
	}
	poolObject, err := pool.GetObject()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("activeNum:%d;idleNum:%d\n", pool.GetNumActive(), pool.GetNumIdle())
		poolObject.(net.Conn).Write([]byte("hello"))
		time.Sleep(1 * time.Second)
	}
}
func TestReturn(t *testing.T) {
	pool, err := DefaultGenObjectPool(NewFactory())
	if err != nil {
		fmt.Println(err)
	}
	poolObject, err := pool.GetObject()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("return activeNum:%d;idleNum:%d\n", pool.GetNumActive(), pool.GetNumIdle())
		poolObject.(net.Conn).Write([]byte("hello"))
		time.Sleep(1 * time.Second)
		pool.ReturnObject(poolObject)
		fmt.Printf("return activeNum:%d;idleNum:%d\n", pool.GetNumActive(), pool.GetNumIdle())
	}
}
func TestDestory(t *testing.T) {
	pool, err := DefaultGenObjectPool(NewFactory())
	if err != nil {
		fmt.Println(err)
	}
	poolObject, err := pool.GetObject()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("activeNum:%d;idleNum:%d\n", pool.GetNumActive(), pool.GetNumIdle())
		poolObject.(net.Conn).Write([]byte("hello"))
		time.Sleep(1 * time.Second)
		pool.DestroyObject(poolObject)
		fmt.Printf("activeNum:%d;idleNum:%d\n", pool.GetNumActive(), pool.GetNumIdle())
	}
}
func TestPoolFull(t *testing.T) {
	config := &PoolConfig{MAX_ACTIVE: 100, MAX_IDLE: 50, MIN_IDLE: 5, MAX_WAIT: 60 * time.Millisecond, IsVaildObj: true}
	pool, err := NewGenObjectPool(NewFactory(), config)
	if err != nil {
		fmt.Println(err)
	}
	pool.InitPool()
	fmt.Printf("pool full activeNum:%d;idleNum:%d\n", pool.GetNumActive(), pool.GetNumIdle())
	var wc sync.WaitGroup
	errnum := 0
	for i := 0; i < 200; i++ {
		wc.Add(1)
		go func(a int) {
			poolObject, err := pool.GetObject()
			if err != nil {
				fmt.Println(err)
				errnum++
			} else {
				fmt.Printf("go %d activeNum:%d;idleNum:%d\n", a, pool.GetNumActive(), pool.GetNumIdle())
				poolObject.(net.Conn).Write([]byte("hello"))
				time.Sleep(50 * time.Millisecond)
				pool.ReturnObject(poolObject)
				fmt.Printf("go %d activeNum:%d;idleNum:%d\n", a, pool.GetNumActive(), pool.GetNumIdle())
			}
			wc.Done()
		}(i)
	}
	// go func() {
	// 	time.Sleep(100 * time.Millisecond)
	// 	pool.Clear()
	// }()
	wc.Wait()
	fmt.Println(errnum)
}
