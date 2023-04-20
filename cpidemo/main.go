package main

import (
	"fmt"
	"runtime"
)

func main()  {
	cpunum := runtime.NumCPU()
	fmt.Println(cpunum)
}
