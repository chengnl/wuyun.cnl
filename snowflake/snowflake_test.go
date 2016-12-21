package snowflake

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestSnowflake(t *testing.T) {
	s, err := Snowflake(1, 1)
	startTime := time.Now()
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < 10000; i++ {
			fmt.Println(s.GetSeqID())
		}

	}
	fmt.Println(time.Since(startTime).Nanoseconds() / 1e6)
}

func TestSnowflake_concurrent(t *testing.T) {
	runtime.GOMAXPROCS(4)
	s, err := Snowflake(1, 1)
	var wg sync.WaitGroup
	startTime := time.Now()
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < 10000; i++ {
			wg.Add(1)
			go func() {
				for j := 0; j < 1000; j++ {
					s.GetSeqID()
				}
				wg.Done()
			}()
		}

	}
	wg.Wait()
	fmt.Println(time.Since(startTime).Nanoseconds() / 1e6)

	//     fmt.Println(int64(6217313433218585064>>22))
//     fmt.Println(int64(6217313433218585064&(1<<12-1)))
//     fmt.Println(int64(6217313433222778880>>22))
//     fmt.Println(int64(6217313433222778880&(1<<12-1)))
}
