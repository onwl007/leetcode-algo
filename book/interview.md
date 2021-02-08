# Golang

### 函数返回指针类型是安全的吗？

### Golang 中如何实现单例模式

### GC 机制是怎样的？

### 锁的实现机制怎样的

### 怎样对切片进行扩容

append 函数

### 无缓冲通道和有缓冲通道的应用场景

### 分布式事务如何做

### Context 有哪些方法，都是干什么的

### Docker & K8s

### 协程与协程之间如何通信

1. 共享全局变量
2. channel
3. context

### sync.Map
map 在并发情况下是有安全问题的，只读是线程安全的，同时写线程不安全，所以为了并发安全和高效，官方实现了 sync.Map

并发写 map 会存在什么问题呢？
```go
func main() {
    m := map[int]int{1:1}
    go do(m)
    go do(m)

    time.Sleep(1*time.Second)
    fmt.Println(m)
}

func do(m map[int]int) {
	i := 0
	for i < 10000 {
		m[1] = 1
		i++
	}
}

// fatal error: concurrent map writes
```

低配版解决方案，加一把大锁
```go
var s sync.RWMutex

func main() {
    m := map[int]int{1: 1}
	go do(m)
	go do(m)
	time.Sleep(1 * time.Second)
	fmt.Println(m)
}

func do(m map[int]int) {
	i := 0
	for i < 10000 {
		s.Lock()
		m[1] = 1
		s.Unlock()
		i++
	}
}
```
加锁不是最优解，一般都会有效率问题，简单的说就是加锁影响其他的元素操作了。

```go
func main() {
    m := sync.Map{}
	m.Store(1, 1)
	go do(m)
	go do(m)
	time.Sleep(1 * time.Second)
	fmt.Println(m.Load(1))
}

func do(m sync.Map) {
	i := 0
	for i < 10000 {
		m.Store(1, 1)
		i++
	}
}
```

# Node.js

# Redis

# Nginx

# MongoDB

# MySQL

# gRPC
