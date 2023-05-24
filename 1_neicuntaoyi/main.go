package main

import (
	"fmt"
	"unsafe"
)

func main() {
	p := new(person)
	fmt.Println(*p)
	//Name是person的第一个字段不用偏移，即可通过指针修改
	pName := (*string)(unsafe.Pointer(p))
	*pName = "微客鸟窝"
	fmt.Println(*p)
	//Age并不是person的第一个字段，所以需要进行偏移，这样才能正确定位到Age字段这块内存，才可以正确的修改
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p))+unsafe.Offsetof(p.Age)))
	*pAge = 20
	fmt.Println(*p)
}
type person struct {
	Name string
	Age int
}