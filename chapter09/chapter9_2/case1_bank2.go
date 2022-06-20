package chapter9_2

// 在8.6节中，我们使用了一个buffered channel作为一个计数信号量，来保证最多只有20个goroutine会同时执行HTTP请求。
// 同理，我们可以用一个容量只有1的channel来保证最多只有一个goroutine在同一时刻访问一个共享变量。
// 一个只能为1和0的信号量叫做二元信号量(binary semaphore)。
// 这种互斥很实用，而且被sync包里的Mutex类型直接支持。它的Lock方法能够获取到token(这里叫锁)，并且Unlock方法会释放这个token（见bank3.go）

var (
	sema     = make(chan struct{}, 1) // a binary semaphore guarding balance
	balance2 int
)

func Deposit2(amount int) {
	sema <- struct{}{} // acquire token
	balance2 = balance2 + amount
	<-sema // release token
}

func Balance2() int {
	sema <- struct{}{} // acquire token
	b := balance2
	<-sema // release token
	return b
}
