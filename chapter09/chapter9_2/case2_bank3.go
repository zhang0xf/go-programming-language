package chapter9_2

import "sync"

// 每次一个goroutine访问bank变量时(这里只有balance余额变量)，它都会调用mutex的Lock方法来获取一个互斥锁。
// 如果其它的goroutine已经获得了这个锁的话，这个操作会被阻塞直到其它goroutine调用了Unlock使该锁变回可用状态。mutex会保护共享变量。
// 在Lock和Unlock之间的代码段中的内容goroutine可以随便读取或者修改，这个代码段叫做临界区。
// 锁的持有者在其他goroutine获取该锁之前需要调用Unlock。
// goroutine在结束后释放锁是必要的，无论以哪条路径通过函数都需要释放，即使是在错误路径中，也要记得释放。

var (
	mu       sync.Mutex // guards balance
	balance3 int
)

func Deposit3(amount int) {
	mu.Lock()
	balance3 = balance3 + amount
	mu.Unlock()
}

func Balance3() int {
	mu.Lock()
	b := balance3
	mu.Unlock()
	return b
}

// 一系列的导出函数封装了一个或多个变量，那么访问这些变量唯一的方式就是通过这些函数来做(一个代理人保证变量被顺序访问)
// 由于在存款和查询余额函数中的临界区代码这么短--只有一行，没有分支调用--在代码最后去调用Unlock就显得更为直截了当。
// 此外，一个deferred Unlock即使在临界区发生panic时依然会执行，这对于用recover (§5.10)来恢复的程序来说是很重要的。
// defer调用只会比显式地调用Unlock成本高那么一点点，不过却在很大程度上保证了代码的整洁性。
// 大多数情况下对于并发程序来说，代码的整洁性比过度的优化更重要。如果可能的话尽量使用defer来将临界区扩展到函数的结束。
func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance3
}

// 考虑一下下面的Withdraw函数。
// 成功的时候，它会正确地减掉余额并返回true。但如果银行记录资金对交易来说不足，那么取款就会恢复余额，并返回false。
// 但是还有一点讨厌的副作用。当过多的取款操作同时执行时，balance可能会瞬时被减到0以下。
// 这可能会引起一个并发的取款被不合逻辑地拒绝。所以如果Bob尝试买一辆sports car时，Alice可能就没办法为她的早咖啡付款了。
// 这里的问题是取款不是一个原子操作：它包含了三个步骤，每一步都需要去获取并释放互斥锁，但任何一次锁都不会锁上整个取款流程。
// NOTE: not atomic!
func Withdraw(amount int) bool {
	Deposit3(-amount)
	if Balance3() < 0 {
		Deposit3(amount)
		return false // insufficient funds
	}
	return true
}

// 理想情况下，取款应该只在整个操作中获得一次互斥锁。下面这样的尝试是错误的：
// Deposit会调用mu.Lock()第二次去获取互斥锁，但因为mutex已经锁上了，而无法被重入(译注：go里没有重入锁，关于重入锁的概念，请参考java)
// --也就是说没法对一个已经锁上的mutex来再次上锁--这会导致程序死锁，没法继续执行下去，Withdraw会永远阻塞下去。
// NOTE: incorrect!
func Withdraw_Error(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	Deposit3(-amount)
	if Balance3() < 0 {
		Deposit3(amount)
		return false // insufficient funds
	}
	return true
}

// 一个通用的解决方案是将一个函数分离为多个函数，比如我们把Deposit分离成两个：
// 一个不导出的函数deposit，这个函数假设锁总是会被保持并去做实际的操作，
// 另一个是导出的函数Deposit，这个函数会调用deposit，但在调用前会先去获取锁。同理我们可以将Withdraw也表示成这种形式：
func Withdraw3(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance3 < 0 {
		deposit(amount)
		return false // insufficient funds
	}
	return true
}

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

// This function requires that the lock be held.
func deposit(amount int) { balance3 += amount }

// 封装(§6.6), 用限制一个程序中的意外交互的方式，可以使我们获得数据结构的不变性。
// 因为某种原因，封装还帮我们获得了并发的不变性。
// 当你使用mutex时，确保mutex和其保护的变量没有被导出(在go里也就是小写，且不要被大写字母开头的函数访问啦)，无论这些变量是包级的变量还是一个struct的字段。
