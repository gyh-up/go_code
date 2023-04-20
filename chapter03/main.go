package main

import (
	"fmt"
	"unsafe"
)

var a1 = 100
var a2 = "fff"
var (
	a3 = "qqq"
	a4 = 1111
)
func main()  {
	//var i int
	//i = 1
	//fmt.Println("i = ", i)
	//var num = 10.10
	//fmt.Println("num = ", num)
	//name := "str"
	//fmt.Println("name = ", name)

	//var n1, n2, n3 int
	//fmt.Println("n1=", n1,"n2=", n2, "n3 =",n3)

	//var n1, n2, n3 = 111, "str", 100.102
	//fmt.Println("n1=", n1,"n2=", n2, "n3 =",n3)

	//n1, n2, n3 := 111, "str", 100.102
	//fmt.Println("n1=", n1,"n2=", n2, "n3 =",n3)

	fmt.Println(a1,a2)
	fmt.Println(a3,a4)

	//返回变量类型
	//Printf 做格式化输出
	var a5 = 100
	fmt.Printf("%T \n",a5)
	//变量占用的字节大小
	var a6 int8
	var a7 int
	fmt.Printf("%T, %d \n", a7,unsafe.Sizeof(a7))
	fmt.Printf("%T, %d \n", a6,unsafe.Sizeof(a6))

	var a8 = 1.1
	fmt.Printf("%T \n",a8)
}
