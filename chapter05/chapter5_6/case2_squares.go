package chapter5_6

import (
	"fmt"
	"log"
)

// 更为重要的是，通过匿名函数这种方式定义的函数可以访问完整的词法环境（lexical environment），这意味着在函数中定义的内部函数可以引用该函数的变量

// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
	var x int
	log.Printf("x = %d\n", x)
	return func() int {
		x++
		return x * x
	}
}

func Squares() {
	log.SetFlags(0)
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}

// squares的例子证明，函数值不仅仅是一串代码，还记录了状态。
// 在squares中定义的匿名内部函数可以访问和更新squares中的局部变量，这意味着匿名函数和squares中，存在变量引用。
// 这就是函数值属于引用类型和函数值不可比较的原因。
// Go使用闭包（closures）技术实现函数值，Go程序员也把函数值叫做闭包。(?)
// 通过这个例子，我们看到变量的生命周期不由它的作用域决定：squares返回后，变量x仍然隐式的存在于f中。
