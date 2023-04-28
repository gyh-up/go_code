package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	//run()
}

//一个协程启动后，一般是代码执行完毕，自动退出，但是如果需要提前终止怎么办呢？
//一个办法是定义一个全局变量，协程中通过检查这个变量的变化来决定是否退出。
//这种办法须要加锁来保证并发安全，说到这里，有没有想的什么解决方案？
//「select + channel」 来实现：

func run1()  {
	var wg sync.WaitGroup
	stopWk := make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker1(stopWk)
	}()
	time.Sleep(3*time.Second) //工作3秒
	stopWk <- true //3秒后发出停止指令
	wg.Wait()
}
func worker1(stopWk chan bool){
	for {
		select {
		case <- stopWk:
			fmt.Println("下班咯~~~")
			return
		default:
			fmt.Println("认真摸鱼中，请勿打扰...")
		}
		time.Sleep(1*time.Second)
	}
}

//上面我们使用 select+channel 来实现了协程的终止，但是如果我们想要同时取消多个协程怎么办呢？如果需要定时取消又怎么办呢？
//此时，Context 就需要登场了，它可以跟踪每个协程，我们重写上面的示例：
func run2() {
	var wg sync.WaitGroup
	ctx, stop := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker2(ctx)
	}()
	go func() {
		defer wg.Done()
		worker2(ctx)
	}()
	time.Sleep(3*time.Second) //工作3秒
	stop() //3秒后发出停止指令
	wg.Wait()
}
func worker2(ctx context.Context){
	for {
		select {
		case <- ctx.Done():
			fmt.Println("下班咯~~~")
			return
		default:
			fmt.Println("认真摸鱼中，请勿打扰...")
		}
		time.Sleep(1*time.Second)
	}
}

// Context 是并发安全的，它是一个接口，可以手动、定时、超时发出取消信号、传值等功能，主要是用于控制多个协程之间的协作、取消操作。

//我们并不需要自己去实现 Context 接口，
//Go 语言提供了函数来生成不同的 Context，通过这些函数可以生成一颗 Context 树，
//这样 Context 就可以关联起来，父级 Context 发出取消信号，子级 Context 也会发出，这样就可以控制不同层级的协程退出。

// 「Background和TODO方法区别:」
//Background和TODO只是用于不同场景下:Background通常被用于主函数、初始化以及测试中，作为一个顶层的context，
//也就是说一般我们创建的context都是基于Background；
//而TODO是在不确定使用什么context的时候才会使用。
//如果一个 Context 有子 Context，在该 Context 取消时，其下的所有子 Context 都会被取消
//Context 不仅可以发出取消信号，还可以传值，可以把它存储的值提供其他协程使用。

func run() {
	var wg sync.WaitGroup
	ctx, stop := context.WithCancel(context.Background())
	valCtx := context.WithValue(ctx, "position","gopher")
	wg.Add(2)
	go func() {
		defer wg.Done()
		worker3(valCtx, "打工人1")
	}()
	go func() {
		defer wg.Done()
		worker3(valCtx, "打工人2")
	}()
	time.Sleep(3*time.Second) //工作3秒
	stop() //3秒后发出停止指令
	wg.Wait()
}
func worker3(valCtx context.Context, name string) {
	for {
		select {
		case <-valCtx.Done():
			fmt.Println("下班咯~~~")
			return
		default:
			position := valCtx.Value("position")
			fmt.Println(name, position, "认真摸鱼中，请勿打扰...")
		}
		time.Sleep(1 * time.Second)
	}
}

//Context 使用原则
//Context 不要放在结构体中，需要以参数方式传递
//Context 作为函数参数时，要放在第一位，作为第一个参数
//使用 context。Background 函数生成根节点的 Context
//Context 要传值必要的值，不要什么都传
//Context 是多协程安全的，可以在多个协程中使用