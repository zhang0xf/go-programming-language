package chapter5_1

import (
	"fmt"
	"math"
)

// 函数

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// 不必为每个形参都写出参数类型
func f1(i, j, k int, s, t string)                { /* ... */ }
func f2(i int, j int, k int, s string, t string) { /* ... */ }

// _符号,可以强调某个参数未被使用。
// 返回值也可以像形式参数一样被命名。在这种情况下，每个返回值被声明成一个局部变量，并根据该返回值的类型，将其初始化为0。
func add(x int, y int) int   { return x + y }
func add2(x, y, z int) int   { return x + y + z }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

func Function() {

	// 在函数调用时，Go语言没有默认参数值
	// 函数的形参和有名返回值作为函数最外层的局部变量，被存储在相同的词法块中。??
	// 实参通过值的方式传递，因此函数的形参是实参的拷贝。
	// 如果实参包括引用类型，如指针，slice(切片)、map、function、channel等类型，实参可能会由于函数的间接引用被修改。
	fmt.Println(hypot(3, 4)) // "5"

	fmt.Printf("%T\n", add)   // "func(int, int) int"
	fmt.Printf("%T\n", add2)  // "func(int, int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"
}

// 你可能会偶尔遇到没有函数体的函数声明，这表示该函数不是以Go实现的。这样的声明定义了函数标识符。
// func Sin(x float64) float64 //implemented in assembly language
