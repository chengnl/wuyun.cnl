package spinlock

import (
	"runtime"
	"sync/atomic"
)

type SpinLock struct {
	f int32
}

func (s *SpinLock) Lock() {
	for !atomic.CompareAndSwapInt32(&s.f, 0, 1) {
		runtime.Gosched()
	}
}

func (s *SpinLock) Unlock() {
	atomic.StoreInt32(&s.f, 0)
}
