// // Package bank provides a concurrency-safe bank with one account.
package chapter9_1

// 下面是一个重写了的银行的例子，这个例子中balance变量被限制在了monitor goroutine中，名为teller：
// 即使当一个变量无法在其整个生命周期内被绑定到一个独立的goroutine，绑定依然是并发问题的一个解决方案。

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit2(amount int) { deposits <- amount }
func Balance2() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
