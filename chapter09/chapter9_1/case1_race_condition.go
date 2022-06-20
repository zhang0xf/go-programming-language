package chapter9_1

import (
	"fmt"
	"image"
)

// 我们会避免并发访问大多数的类型，无论是将变量局限在单一的一个goroutine内，还是用互斥条件维持更高级别的不变性，都是为了这个目的。
// 相反，包级别的导出函数一般情况下都是并发安全的。由于package级的变量没法被限制在单一的gorouine，所以修改这些变量“必须”使用互斥条件。
// 一个函数在并发调用时没法工作的原因太多了，比如死锁(deadlock)、活锁(livelock)和饿死(resource starvation)。
// 竞争条件指的是程序在多个goroutine交叉执行操作时，没有给出正确的结果。
// 竞争条件是很恶劣的一种场景，因为这种问题会一直潜伏在你的程序里，然后在非常少见的时候蹦出来，或许只是会在很大的负载时才会发生，又或许是会在使用了某一个编译器、某一种平台或者某一种架构的时候才会出现。
// 这些使得竞争条件带来的问题非常难以复现而且难以分析诊断。

// 传统上经常用经济损失来为竞争条件做比喻，所以我们来看一个简单的银行账户程序。

var balance int

func Deposit(amount int) { balance = balance + amount }
func Balance() int       { return balance }

// 然而，当我们并发地而不是顺序地调用这些函数的话，Balance就再也没办法保证结果正确了。

func RaceCondition() {
	// Alice:
	go func() {
		Deposit(200)                // A1
		fmt.Println("=", Balance()) // A2
	}()

	// Bob:
	go Deposit(100) // B
}

// 因为Alice的存款操作A1实际上是两个操作的一个序列，读取然后写；可以称之为A1r和A1w。
// 这个程序包含了一个特定的竞争条件，叫作数据竞争。无论任何时候，只要有两个goroutine并发访问同一变量，且至少其中的一个是写操作的时候就会发生数据竞争。

// 如果数据竞争的对象是一个比一个机器字(译注：32位机器上一个字=4个字节)更大的类型时，事情就变得更麻烦了，比如interface，string或者slice类型都是如此。下面的代码会并发地更新两个不同长度的slice：
// 最后一个语句中的x的值是未定义的；其可能是nil，或者也可能是一个长度为10的slice，也可能是一个长度为1,000,000的slice。
// 。但是回忆一下slice的三个组成部分：指针(pointer)、长度(length)和容量(capacity)。如果指针是从第一个make调用来，而长度从第二个make来，x就变成了一个混合体，一个自称长度为1,000,000但实际上内部只有10个元素的slice。
func RaceCondition2() {
	var x []int
	go func() { x = make([]int, 10) }()
	go func() { x = make([]int, 1000000) }()
	x[999999] = 1 // NOTE: undefined behavior; memory corruption possible!
}

// 一个好的经验法则是根本就没有什么所谓的良性数据竞争。所以我们一定要避免数据竞争，那么在我们的程序中要如何做到呢？
// 数据竞争的定义：数据竞争会在两个以上的goroutine并发访问相同的变量且至少其中一个为写操作时发生。根据上述定义，有三种方式可以避免数据竞争：
// 第一种方法是不要去写变量。
// 考虑一下下面的map，会被“懒”填充，也就是说在每个key被第一次请求到的时候才会去填值。
// 如果Icon是被顺序调用的话，这个程序会工作很正常，但如果Icon被并发调用，那么对于这个map来说就会存在数据竞争。(有写)
var icons = make(map[string]image.Image)

func loadIcon(name string) image.Image {
	// ...
	return nil
}

// NOTE: not concurrency-safe!
func Icon(name string) image.Image {
	icon, ok := icons[name]
	if !ok {
		icon = loadIcon(name)
		icons[name] = icon
	}
	return icon
}

// 反之，如果我们在创建goroutine之前的初始化阶段，就初始化了map中的所有条目并且再也不去修改它们，那么任意数量的goroutine并发访问Icon都是安全的，因为每一个goroutine都只是去读取而已。
var icons2 = map[string]image.Image{
	"spades.png":   loadIcon("spades.png"),
	"hearts.png":   loadIcon("hearts.png"),
	"diamonds.png": loadIcon("diamonds.png"),
	"clubs.png":    loadIcon("clubs.png"),
}

// Concurrency-safe.
func Icon2(name string) image.Image { return icons2[name] } // ( 只读)

// 上面的例子里icons变量在包初始化阶段就已经被赋值了，包的初始化是在程序main函数开始执行之前就完成了的。

// 第二种避免数据竞争的方法是，避免从多个goroutine访问变量。这也是前一章中大多数程序所采用的方法。
// 由于其它的goroutine不能够直接访问变量，它们只能使用一个channel来发送请求给指定的goroutine来查询更新变量。
// 这也就是Go的口头禅“不要使用共享数据来通信；使用通信来共享数据”。
// 一个提供对一个指定的变量通过channel来请求的goroutine叫做这个变量的monitor（监控）goroutine。例如broadcaster goroutine会监控clients map的全部访问。
// 即使当一个变量无法在其整个生命周期内被绑定到一个独立的goroutine，绑定依然是并发问题的一个解决方案。
// 例如在一条流水线上的goroutine之间共享变量是很普遍的行为，在这两者间会通过channel来传输地址信息。
// 如果流水线的每一个阶段都能够避免在将变量传送到下一阶段后再去访问它，那么对这个变量的所有访问就是线性的。
// 其效果是变量会被绑定到流水线的一个阶段，传送完之后被绑定到下一个，以此类推。这种规则有时被称为串行绑定。
// 下面的例子中，Cakes会被严格地顺序访问，先是baker gorouine，然后是icer gorouine：

type Cake struct{ state string }

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake // baker never touches this cake again
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake // icer never touches this cake again
	}
}

func CakeLine() {
	var ICED chan *Cake
	var COOKED chan *Cake // 双向

	baker(COOKED)      // 单向
	icer(ICED, COOKED) // 单向
}

// 第三种避免数据竞争的方法是允许很多goroutine去访问变量，但是在同一个时刻最多只有一个goroutine在访问。这种方式被称为“互斥”，在下一节来讨论这个主题。
