package chapter5_6

import (
	"strings"
)

// 拥有函数名的函数只能在包级语法块中被声明，通过函数字面量（function literal），我们可绕过这一限制，在任何表达式中表示一个函数值。
// 函数字面量的语法和函数声明相似，区别在于func关键字后没有函数名。
// 函数值字面量是一种表达式，它的值被称为匿名函数（anonymous function）。
// 函数字面量允许我们在使用函数时，再定义它。

func AnonymousFunc() {
	strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000")
}
