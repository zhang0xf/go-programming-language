package chapter3_6

import (
	"fmt"
	"math"
	"time"
)

// 常量作数组长度
const IPv4Len = 4

type IP [IPv4Len]byte

// parseIPv4 parses an IPv4 address (d.d.d.d).
func parseIPv4(s string) IP {
	var p [IPv4Len]byte
	// ...
	return p
}

func Const() {
	// 常量定义
	const pi1 = 3.14159 // approximately; math.Pi is a better approximation

	const (
		e   = 2.71828182845904523536028747135266249775724709369995957496696763
		pi2 = 3.14159265358979323846264338327950288419716939937510582097494459
	)

	// 常量的类型和初始化
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

	// 常量初始化
	const (
		a = 1
		b
		c = 2
		d
	)

	fmt.Println(a, b, c, d) // "1 1 2 2"

	// 无类型常量
	// 提供更高的运算精度，而且可以直接用于更多的表达式而不需要显式的类型转换
	// ZiB和YiB的值已经超出任何Go语言中整数类型能表达的范围，但是它们依然是合法的常量
	// YiB/ZiB是在编译期计算出来的，并且结果常量是1024，是Go语言int变量能有效表示的
	fmt.Println(YiB / ZiB) // "1024"

	// 常量除法表达式可能对应不同的结果:
	var f1 float64 = 212
	fmt.Println((f1 - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
	fmt.Println(5 / 9 * (f1 - 32))     // "0";   5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f1 - 32)) // "100"; 5.0/9.0 is an untyped float

	// 隐式转换和显示转换
	var f2 float64 = 3 + 0i // untyped complex -> float64
	f2 = 2                  // untyped integer -> float64
	f2 = 1e123              // untyped floating-point -> float64
	f2 = 'a'                // untyped rune -> float64
	fmt.Println(f2)

	var f3 float64 = float64(3 + 0i)
	f3 = float64(2)
	f3 = float64(1e123)
	f3 = float64('a')
	fmt.Println(f3)

	const Pi64 float64 = math.Pi

	var x1 float32 = math.Pi
	var y1 float64 = math.Pi
	var z1 complex128 = math.Pi
	fmt.Println(x1)
	fmt.Println(y1)
	fmt.Println(z1)

	var x2 float32 = float32(Pi64)
	var y2 float64 = Pi64
	var z2 complex128 = complex128(Pi64)
	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Println(z2)

	// 隐式决定类型
	i4 := 0      // untyped integer;        implicit int(0)
	r4 := '\000' // untyped rune;           implicit rune('\000')
	f4 := 0.0    // untyped floating-point; implicit float64(0.0)
	c4 := 0i     // untyped complex;        implicit complex128(0i)
	fmt.Println(i4)
	fmt.Println(r4)
	fmt.Println(f4)
	fmt.Println(c4)
	fmt.Printf("%T\n", 0)      // "int"
	fmt.Printf("%T\n", 0.0)    // "float64"
	fmt.Printf("%T\n", 0i)     // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)

	// 显示转换类型
	var i5 = int8(0)
	var i6 int8 = 0
	fmt.Println(i5)
	fmt.Println(i6)
}
