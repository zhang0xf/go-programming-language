package chapter8_1

// 在Go语言中，每一个并发的执行单元叫作一个goroutine。
// 如果你使用过操作系统或者其它语言提供的线程，那么你可以简单地把goroutine类比作一个线程，这样你就可以写出一些正确的程序了。
// goroutine和线程的本质区别会在9.8节中讲。

// 当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。
// 新的goroutine会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。
// go语句会使其语句中的函数在一个新创建的goroutine中运行。而go语句本身会迅速地完成。

func f() {} // call f(); wait for it to return

func Goroutines() {
	go f() // create a new goroutine that calls f(); don't wait
}
