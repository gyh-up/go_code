package main

import "fmt"

func main(){
	var b int = 10
	fmt.Println(&b)
	var a float32 = 3.22
	var ptr *float32 = &a
	fmt.Println(ptr)
	*ptr = 999
	fmt.Println(a)
}
