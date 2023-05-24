package castest

import (
	"sync"
	"sync/atomic"
)

// 锁实现方式
func Lock() int64 {
	var count int64
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			mu.Lock()
			count = count + 1
			mu.Unlock()
		}(&wg)
	}
	wg.Wait()
	// count = 10000
	//fmt.Println("count = ", count)
	return count
}

// atomic CAS 原子操作
func Cas() int64 {
	var count int64
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 失败一直重试
			for {
				old := atomic.LoadInt64(&count)
				if atomic.CompareAndSwapInt64(&count, old, old+1) {
					break
				}
			}
		}(&wg)
	}
	wg.Wait()
	// count = 10000
	//fmt.Println("count = ", count)
	return count
}