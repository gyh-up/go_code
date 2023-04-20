package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	test()
}

var sum = 0
func test1() {
	//开启100个协程来让 sum + 1
	for i := 1; i <= 100000; i++ {
		go add1()
	}
	// 睡眠两秒防止程序提前退出
	time.Sleep(2 * time.Second)
	fmt.Println("sum:",sum)
}
func add1(){
	sum += 1
}

// 所谓资源竞争，就是在程序中，同一块内存同时被多个 channel 访问。
//对于这个共享的资源（内存）每个 channel 都有不同的操作，就有可能造成数据紊乱。
//使用 go build、go run、go test 命令时，添加 -race 标识可以检查代码中是否存在资源竞争
var mutex = sync.Mutex{}
func test2 () {
	//开启100个协程来让 sum + 1
	for i := 1; i <= 100000; i++ {
		go add2()
	}
	// 睡眠两秒防止程序提前退出
	time.Sleep(2 * time.Second)
	fmt.Println("sum:",sum)
}
func add2(){
	mutex.Lock()
	defer mutex.Unlock() //使用defer语句，确保锁一定会被释放
	sum += 1
}

// 上面我们使用互斥锁，来防止多个协程同时对 sum 做加法操作的时候产生数据错乱。
//RWMutex为读写锁，当读取竞争资源的时候，因为数据不会改变，所以不管多少个 channel 读都是并发安全的。
//因为可以多个协程同时读，不再相互等待，所以在性能上比互斥锁会有很大的提升。
var rwmutex = sync.RWMutex{}
func test3 () {
	//开启100个协程来让 sum + 1
	for i := 1; i <= 10000; i++ {
		go add3()
	}
	for i := 1; i<= 10; i++ {
		go fmt.Println("sum:",getSum3())
	}
	// 睡眠两秒防止程序提前退出
	time.Sleep(2 * time.Second)
	fmt.Println("sum:", sum)
}
func add3(){
	mutex.Lock()
	defer mutex.Unlock() //使用defer语句，确保锁一定会被释放
	sum += 1
}
func getSum3() int {
	//rwmutex.RLock() //使用读写锁
	//defer  rwmutex.RUnlock()
	return sum
}

// 上面的示例中，我们都是要了 time.Sleep(2 * time.Second)，来防止：主函数 mian 返回，提前退出程序。
//但是我们并不知道程序真正什么时候执行完，所以只能设置个长点的时间避免程序提前退出，这样会产生性能问题。
//这时候我们就用到了 sync.WaitGroup ，它可以监听程序的执行，一旦全部执行完毕，程序就能马上退出。
func test4() {
	var wg sync.WaitGroup
	//因为要监控110个协程，所以设置计数器为110
	wg.Add(10010)
	for i := 1; i <= 10000; i++ {
		go func() {
			//计数器值减1
			defer wg.Done()
			add4()
		}()
	}
	for i := 1; i <= 10; i++ {
		go func() {
			//计数器值减1
			defer wg.Done()
			fmt.Println("sum:", getSum4())
		}()
	}
	//一直等待，只要计数器值为0
	wg.Wait()
}
func add4() {
	mutex.Lock()
	defer mutex.Unlock() //使用defer语句，确保锁一定会被释放
	sum += 1
}
func getSum4() int {
	rwmutex.RLock() //使用读写锁
	defer rwmutex.RUnlock()
	return sum
}


// 有时候我们只希望代码执行一次，即使是在高并发的场景下，
//比如创建一个单例。这种情况可以使用 sync.Once 来保证代码只执行一次。
func test5() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	//用于等待协程执行完毕
	done := make(chan bool)
	//启动10个协程执行once.Do(onceBody)
	for i := 0; i < 10; i++ {
		go func() {
			//把要执行的函数(方法)作为参数传给once.Do方法即可
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		res, _ := <-done
		fmt.Println(res)
	}
}
//上面这个是Go语言自带的示例，虽然启动了10个协程来执行 onceBody  函数，但是 once.DO 方法保证 onceBody 函数只会执行一次。
//sync.Once 适合用于创建单例、只加载一次资源等只需要执行一次的场景。


// 	我们有一项任务，只有满足了条件情况下才能执行，否则就等着。如何获取这个条件呢？
//	可以使用 channel 的方式，但是 channel 适用于一对一，一对多就需要用到 sync.Cond。
//sync.Cond 是基于互斥锁的基础上，增加了一个通知队列，协程刚开始是等待的，通知的协程会从通知队列中唤醒一个或多个被通知的协程。
//sync.Cond 主要有以下几个方法：
//sync.NewCond(&mutex) //sync.Cond 通过sync.NewCond初始化，需要传入一个mutex，因为阻塞等待通知的操作以及通知解除阻塞的操作就是基于sync.Mutex来实现的。
//sync.Wait() //等待通知
//阻塞当前协程，直到被其他协程调用 Broadcast 或者 Signal 方法唤醒，使用的时候需要加锁，使用 sync.Cond 中的锁即可
//sync.Signal() //单发通知，随机唤醒一个协程
//sync.Broadcat() //广播通知，唤醒所有等待的协程。
func test()  {
	//3个人赛跑，1个裁判员发号施令
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(4) //3选手+1裁判
	for i := 1; i <= 3; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号选手已经就位")
			cond.L.Lock()
			cond.Wait() //等待发令枪响
			fmt.Println(num, "号选手开始跑……")
			cond.L.Unlock()
		}(i)
	}
	//等待所有goroutine都进入wait状态
	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("裁判：“各就各位~~预备~~”")
		fmt.Println("啪！！！")
		cond.Broadcast() //发令枪响
	}()
	//防止函数提前返回退出
	wg.Wait()
}