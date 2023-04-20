package main

import (
	"fmt"
	"time"
)
func main() {
	test1()
}

func test1() {
	ch := make(chan string)
	go func(){
		fmt.Println("微客鸟窝")
		ch <- "执行完毕"
	}()

	fmt.Println("我是无尘啊")
	time.Sleep(time.Second*2)
	value := <-ch
	fmt.Println("获取的chan的值为：",value)
}
// 上面的操作就是一个无缓冲 channel，通道的容量是0，它不能存储数据，只是起到了传输的作用，
//所以无缓冲 channel 的发送和接收操作是同时进行的

// 在声明的时候，我们可以传入第二个参数，即「channel容量大小」，这样就是创建了一个有缓冲 channel。
//有缓冲 channel 内部有一个队列
//发送操作是向队列尾部追加元素，如果队列满了，则阻塞等待，直到接收操作从队列中取走元素。
//接收操作是从队列头部取走元素，如果队列为空，则阻塞等待，直到发送操作向队列追加了元素。
//可以通过内置函数 cap 来获取 channel 的容量，通过内置函数 len 获取 channel 中元素个数。

//使用内置函数 close :close(ch)
//channel 关闭了就不能再向其发送数据了，否则会引起 1_panic 异常。
//可以从关闭了的 channel 中接收数据，如果没数据，则接收到的是元素类型的零值。、


// select 可以实现多路复用，即同时监听多个 channel。
//发现哪个 channel 有数据产生，就执行相应的 case 分支
//如果同时有多个 case 分支可以执行，则会随机选择一个
//如果一个 case 分支都不可执行，则 select 会一直等待

func test2() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
			fmt.Println("--", i)
		}
	}
}