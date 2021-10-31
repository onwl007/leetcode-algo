package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func main() {
	// fast()
	singleNotify()
}

// 异步请求，类似于 Node.js 中的 async/await
func promise() {
	// a, b := longTimeRequest(), longTimeRequest1()
	// result := sumSquares(<-a, <-b)
	// fmt.Println(result)
	ra := make(chan int32, 2)
	go longTime(ra)
	go longTime(ra)
	fmt.Println(sumSquares(<-ra, <-ra))
}

func longTimeRequest() <-chan int32 {
	r := make(chan int32)
	go func() {
		time.Sleep(3 * time.Second) // 模拟一个工作负载
		r <- 3
	}()

	return r
}

func longTimeRequest1() <-chan int32 {
	r := make(chan int32)
	go func() {
		time.Sleep(4 * time.Second) // 模拟一个工作负载
		r <- 4
	}()

	return r
}

// 超时
// 在一些请求/回应用例中，一个请求可能因为种种原因导致需要超出预期的时长才能得到回应，有时甚至永远得不到回应
func requestWithTimeout(timeout time.Duration) (int, error) {
	c := make(chan int)
	go func(c chan int) {
		time.Sleep(2 * time.Second)
		c <- 10
	}(c) // 可能需要超出预期的时长回应

	select {
	case data := <-c:
		return data, nil
	case <-time.After(timeout):
		return 0, errors.New("超时了")
	}
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func longTime(r chan<- int32) {
	time.Sleep(3 * time.Second) // 模拟工作负载
	r <- 3
}

func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)+1
	// 睡眠 1 秒、2 秒、3 秒
	time.Sleep(time.Duration(rb) * time.Second)
	c <- ra
}

// 采用最快回应
// 利用缓冲通道
// 注意：如果有 N 个数据源，为了防止被舍弃的回应对应的协程永久阻塞，则专属数据用的通道必须为一个容量至少为 N-1 的缓冲通道
func fast() {
	rand.Seed(time.Now().Unix())

	starTime := time.Now()
	c := make(chan int32, 5) // 必须用一个缓冲通道
	for i := 0; i < cap(c); i++ {
		go source(c)
	}
	rnd := <-c // 只有第一个回应被使用了
	fmt.Println(time.Since(starTime))
	fmt.Println(rnd)
}

// 使用通道实现通知
// 如果一个通道中无值可接收，则此通道上的下一个接收操作将阻塞到另一个协程发送一个值到此通道为止

// 单对单通知
// 向一个通道发送一个值来实现单对单通知
func singleNotify() {
	values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}) // 也可以是缓冲的

	// 排序协程
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{} // 通知排序已完成
	}()
	// 并发地做一些其他事情

	<-done // 等待通知
	fmt.Println(values[0], values[len(values)-1])
}

// 从一个通道接收一个值来实现单对单通知
// 如果一个通道的数据缓冲队列已满（非缓冲的通道的数据缓冲队列总是满的）但它的发送协程队列为空，则向此通道发送一个值将阻塞，
// 直到另外一个协程从此通道接收一个值为止。
// 所以我们可以通过从一个通道接收数据来实现单对单通知。一般我们使用非缓冲通道来实现这样的通知
func receiveSingleNotify() {
	// 此信号通道也可以缓冲为 1，如果这样，则在下面这个协程创建之前，我们必须向其中写入一个值
	done := make(chan struct{})

	go func() {
		fmt.Println("Hello")
		time.Sleep(time.Second * 2)
		// 使用一个接收操作来通知主协程
		<-done
	}()

	done <- struct{}{} // 阻塞在此，等待通知
	fmt.Println("world")
}

type T = struct{}

// 多对单和单对多通知
// 该例中的方式在实践中用的并不多。
// 在实践中，我们多使用 WaitGroup 来实现多对单通知，使用关闭一个通道的方式来实现单对多通知
func multiNotify() {
	ready, done := make(chan T), make(chan T)
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	// 模拟一个初始化过程
	time.Sleep(time.Second * 3 / 2)
	// 单对多通知
	ready <- T{}
	ready <- T{}
	ready <- T{}

	// 等待被多对单通知
	<-done
	<-done
	<-done
}

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready // 阻塞在此，等待通知
	fmt.Println("Worker#", id, "开始工作")
	// 模拟一个工作负载
	time.Sleep(time.Second * time.Duration(id+1))
	fmt.Println("Worker#", id, "工作完成")
	done <- T{} // 通知主协程
}

// 通过关闭一个通道来实现群发通知

// 将通道用做互斥锁
// 容量为 1 的缓冲通道也可以用做多次性二元信号量（即互斥锁）尽管这样的互斥锁效率不如 sync 标准库中提供的互斥锁高效
// 1. 通过发送操作来加锁，通过接收操作来解锁
// 2. 通过接收操作来加锁，通过发送操作来解锁
func mutexChan() {
	mutex := make(chan struct{}, 1) // 容量必须为 1

	counter := 0
	increase := func() {
		mutex <- T{} // 加锁
		counter++
		<-mutex // 解锁
	}

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done
	fmt.Println(counter) // 2000
}

// 对话（俗称乒乓）
// 两个协程可以通过一个通道进行对话，整个过程宛如打乒乓球一样
type Ball uint64

func Play(playerName string, table chan Ball) {
	var lastValue Ball = 1
	for {
		ball := <-table // 接球
		fmt.Println(playerName, ball)
		ball += lastValue
		if ball < lastValue { // 溢出结束
			os.Exit(0)
		}
		lastValue = ball
		table <- ball // 回球
		time.Sleep(time.Second)
	}
}

func pingpang() {
	table := make(chan Ball)
	go func() {
		table <- 1 // 裁判发球
	}()
	go Play("A: ", table)
	Play("B: ", table)
}

// 速率限制
// 在此例中，任何一分钟时段内处理的请求数不会超过200
type Request interface{}

func handle(r Request) {
	fmt.Println(r.(int))
}

const (
	RateLimitPeriod = time.Minute
	RateLimit       = 200 // 任何一分钟内最多处理 200 个请求
)

func handleRequests(requests <-chan Request) {
	quotas := make(chan time.Time, RateLimit)

	go func() {
		tick := time.NewTicker(RateLimitPeriod / RateLimit)
		defer tick.Stop()
		for t := range tick.C {
			select {
			case quotas <- t:
			default:
			}
		}
	}()

	for r := range requests {
		<-quotas
		go handle(r)
	}
}

// 如果在开始的一分钟内没有任何请求，则接下来的的某个瞬时时间点可能会同时处理最多 200 个请求
// 这可能会造成卡顿情况。我们可以将速率限制和峰值限制一并使用来避免出现这样的情况
func RateExample() {
	requests := make(chan Request)
	go handleRequests(requests)
	// time.Sleep(time.Minute)
	for i := 0; ; i++ {
		requests <- i
	}
}
