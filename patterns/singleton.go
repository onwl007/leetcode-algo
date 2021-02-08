package patterns

import "sync"

type Singleton struct {
}

var instance1 *Singleton

// 非线程安全的单例模式，不能在多线程情况下使用
// 多个 goroutine 可以执行第一个检查，并且都将创建该 Singleton 类型的实例并互相覆盖。
// 无法保证它将在此处返回哪个实例，并且对该实例的其他进一步操作可能与开发人员的期望不一致。
func GetInstance1() *Singleton {
	if instance1 == nil {
		instance1 = &Singleton{}
	}
	return &Singleton{}
}

// 线程锁实现的单例模式
// 执行了过多的锁定，在实例已经创建的情况下，我们应该简单地返回缓存的单实例。
// 在高度并发的代码基础上，这可能会产生瓶颈，因为一次只有一个 goroutine 可以获得单例实例。
var (
	instance2 *Singleton
	mtx       sync.Mutex
)

func GetInstance2() *Singleton {
	mtx.Lock()
	defer mtx.Unlock()

	if instance2 == nil {
		instance2 = &Singleton{}
	}
	return instance2
}

// 检查锁的方式，双检测锁实现方式
// 你应该首先进行检查，以最小化任何主动锁定，因为IF语句的开销要比加锁小。
// 其次，我们希望等待并获取互斥锁，这样在同一时刻在那个块中只有一个执行。
// 但是，在第一次检查和获取互斥锁之间，可能有其他goroutine获取了锁，
// 因此，我们需要在锁的内部再次进行检查，以避免用另一个实例覆盖了实例。
var (
	instance3 *Singleton
	mut       sync.Mutex
)

func GetInstance3() *Singleton {
	if instance3 == nil {
		mut.Lock()
		defer mut.Unlock()

		if instance3 == nil {
			instance3 = &Singleton{}
		}
	}
	return instance3
}

// 更完美的方式，Go 语言风格的单例模式
// 使用 once.Do 来保证某个对象只会初始化一次，有一点需要注意的是，
// 这个方法只会被运行一次，哪怕 Do 函数里面发生了异常，对象初始化失败了，这个 Do 函数也不会被再次执行了
var (
	instance4 *Singleton
	once      sync.Once
)

func GetInstance4() *Singleton {
	if instance4 == nil {
		once.Do(func() {
			instance4 = &Singleton{}
		})
	}
	return instance4
}
