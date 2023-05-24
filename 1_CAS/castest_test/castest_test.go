package castest_test

import (
	"fmt"
	"go_code/1_CAS/castest"
	"testing"
)


func TestLock(t *testing.T) {
	castest.Lock()
}

func BenchmarkLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		castest.Lock()
		fmt.Println(1)
	}
	//b.RunParallel(func(pb *testing.PB) {
	//	for pb.Next() {
	//		castest.Cas()
	//	}
	//})
}

func BenchmarkCas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		castest.Cas()
	}
	//b.RunParallel(func(pb *testing.PB) {
	//	for pb.Next() {
	//		castest.Cas()
	//	}
	//})
}