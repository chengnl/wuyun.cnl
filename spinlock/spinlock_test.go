package spin

import (
	"sync"
	"testing"
)

type SpinMap struct {
	SpinLock
	m map[int]bool
}

func (sm *SpinMap) Add(i int) {
	sm.Lock()
	sm.m[i] = true
	sm.UnLock()
}

func (sm *SpinMap) Get(i int) (b bool) {
	sm.Lock()
	b = sm.m[i]
	sm.UnLock()
	return
}

type MutexMap struct {
	sync.Mutex
	m map[int]bool
}

func (mm *MutexMap) Add(i int) {
	mm.Lock()
	mm.m[i] = true
	mm.Unlock()
}

func (mm *MutexMap) Get(i int) (b bool) {
	mm.Lock()
	b = mm.m[i]
	mm.Unlock()
	return
}

type RWMutexMap struct {
	sync.RWMutex
	m map[int]bool
}

func (mm *RWMutexMap) Add(i int) {
	mm.Lock()
	mm.m[i] = true
	mm.Unlock()
}

func (mm *RWMutexMap) Get(i int) (b bool) {
	mm.RLock()
	b = mm.m[i]
	mm.RUnlock()
	return
}

const N = 10

var (
	sm   = SpinMap{m: map[int]bool{}}
	mm   = MutexMap{m: map[int]bool{}}
	rwmm = RWMutexMap{m: map[int]bool{}}
)

func BenchmarkSpinL(b *testing.B) {
	var wg sync.WaitGroup
	b.N = 100
	for i := 0; i < b.N; i++ {
		wg.Add(N * 2)
		for i := 0; i < N; i++ {
			go func(i int) {
				sm.Add(i)
				wg.Done()
			}(i)
			go sm.Get(i)

			go func(i int) {
				sm.Get(i - 1)
				sm.Get(i + 1)
				wg.Done()
			}(i)
		}
	}
	wg.Wait()
}

func BenchmarkMutex(b *testing.B) {
	var wg sync.WaitGroup
	b.N = 100
	for i := 0; i < b.N; i++ {
		wg.Add(N * 2)
		for i := 0; i < N; i++ {
			go func(i int) {
				mm.Add(i)
				wg.Done()
			}(i)
			go mm.Get(i)
			go func(i int) {
				mm.Get(i - 1)
				mm.Get(i + 1)
				wg.Done()
			}(i)
		}
	}
	wg.Wait()
}

func BenchmarkRWMutex(b *testing.B) {
	var wg sync.WaitGroup
	b.N = 100
	for i := 0; i < b.N; i++ {
		wg.Add(N * 2)
		for i := 0; i < N; i++ {
			go func(i int) {
				rwmm.Add(i)
				wg.Done()
			}(i)
			go rwmm.Get(i)
			go func(i int) {
				rwmm.Get(i - 1)
				rwmm.Get(i + 1)
				wg.Done()
			}(i)
		}
	}
	wg.Wait()
}
