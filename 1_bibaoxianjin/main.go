package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	run2()
}

const  N = 10
func run1() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			m[i] = i
		}()
	}
	wg.Wait()
	println(len(m))
}

func run2 () {
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func() {
			fmt.Println(v)
		}()
		time.Sleep(time.Second)
	}
	select {}    // 阻塞模式

}