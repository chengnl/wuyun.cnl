package snowflake

import (
	"fmt"
	"log"
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
		for i := 0; i < 10000000; i++ {
			s.GetSeqID()
		}

	}
	fmt.Println(time.Since(startTime).Nanoseconds() / 1e6)
}

func TestSnowflake_concurrent(t *testing.T) {
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
					m, err := fmt.Println("111")
					if err != nil {
						log.Fatalf("fmt err %d %v", m, err)
					}
					s.GetSeqID()
				}
				wg.Done()
			}()
		}

	}
	wg.Wait()
	fmt.Println(time.Since(startTime).Nanoseconds() / 1e6)
}
