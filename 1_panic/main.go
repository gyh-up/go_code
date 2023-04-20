package main

import "fmt"

// 一般我们不对panic异常做处理，但是如果有一些需要在程序崩溃前做处理的操作，可以使用内置的 recover 函数来恢复 1_panic 异常。
// 程序 1_panic 异常崩溃的时候，只有defer修饰的函数才会被执行，所以 recover 函数要结合 defer 关键字一起使用：
// 1_panic 是一种非常严重的错误，会使程序中断执行，所以 「如果不是影响程序运行的错误，使用 error 即可」
func main() {
	defer func() {
		if p := recover(); p != any(nil) {
			fmt.Println(p)
		}
	}()
	connectMySQL("","root","123456")
}
func connectMySQL(ip,username,password string) {
	if ip == "" {
		panic(any("ip不能为空"))
	}
	//省略其他代码
}