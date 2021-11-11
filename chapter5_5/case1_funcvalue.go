package chapter5_5

import (
	"fmt"
	"strings"
)

// 在Go中，函数被看作第一类值（first-class values）：
// 函数像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。对函数值（function value）的调用类似函数调用。

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

// 函数值
func FunctionValue() {

	f := square
	fmt.Println(f(3)) // "9"

	f = negative
	fmt.Println(f(3))     // "-3"
	fmt.Printf("%T\n", f) // "func(int) int"

	// f = product // compile error: can't assign func(int, int) int to func(int) int
}

// 函数零值
// 函数类型的零值是nil。
func FunctionNilValue() {
	var f func(int) int
	f(3) // 此处f的值为nil, 会引起panic错误

	// 函数值可以与nil比较
	// 但是函数值之间是不可比较的，也不能用函数值作为map的key。
	if f != nil {
		f(3)
	}
}

func add1(r rune) rune { return r + 1 }

// 函数值使得我们不仅仅可以通过数据来参数化函数，亦可通过行为。(?)
// strings.Map对字符串中的每个字符调用add1函数，并将每个add1函数的返回值组成一个新的字符串返回给调用者。
func FunctionValueSkill() {
	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
}
