package chapter9_4

import "fmt"

// 你可能比较纠结为什么Balance方法需要用到互斥条件，无论是基于channel还是基于互斥量。
// 毕竟和存款不一样，它只由一个简单的操作组成，所以不会碰到其它goroutine在其执行“期间”执行其它逻辑的风险。
// 这里使用mutex有两方面考虑。第一Balance不会在其它操作比如Withdraw“中间”执行。
// 第二（更重要的）是“同步”不仅仅是一堆goroutine执行顺序的问题，同样也会涉及到内存的问题。

// 在现代计算机中可能会有一堆处理器，每一个都会有其本地缓存(local cache)。
// 为了效率，对内存的写入一般会在每一个处理器中缓冲，并在必要时一起flush到主存。
// 这种情况下这些数据可能会以与当初goroutine写入顺序不同的顺序被提交到主存。
// 像channel通信或者互斥量操作这样的原语会使处理器将其聚集的写入flush并commit，这样goroutine在某个时间点上的执行结果才能被其它处理器上运行的goroutine得到。

// 考虑一下下面代码片段的可能输出：
func MemorySync() {
	var x, y int

	go func() {
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()

	go func() {
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()
}

// 我们可能希望它能够打印出下面这四种结果中的一种，相当于几种不同的交错执行时的情况：
// y:0 x:1
// x:0 y:1
// x:1 y:1
// y:1 x:1

// 然而实际运行时还是有些情况让我们有点惊讶：
// x:0 y:0
// y:0 x:0

// 根据所使用的编译器，CPU，或者其它很多影响因子，这两种情况也是有可能发生的。那么这两种情况要怎么解释呢？
// 在一个独立的goroutine中，每一个语句的执行顺序是可以被保证的，也就是说goroutine内顺序是连贯的。
// 但是在不使用channel且不使用mutex这样的显式同步操作时，我们就没法保证事件在不同的goroutine中看到的执行顺序是一致的了。
// 尽管goroutine A中一定需要观察到x=1执行成功之后才会去读取y，但它没法确保自己观察得到goroutine B中对y的写入，所以A还可能会打印出y的一个旧版的值。
// 尽管去理解并发的一种尝试是去将其运行理解为不同goroutine语句的交错执行，但看看上面的例子，这已经不是现代的编译器和cpu的工作方式了。
// 因为赋值和打印指向不同的变量，编译器可能会断定两条语句的顺序不会影响执行结果，并且会交换两个语句的执行顺序。
// 如果两个goroutine在不同的CPU上执行，每一个核心有自己的缓存，这样一个goroutine的写入对于其它goroutine的Print，在主存同步之前就是不可见的了。

// 所有并发的问题都可以用一致的、简单的既定的模式来规避。所以可能的话，将变量限定在goroutine内部；如果是多个goroutine都需要访问的变量，使用互斥条件来访问。
