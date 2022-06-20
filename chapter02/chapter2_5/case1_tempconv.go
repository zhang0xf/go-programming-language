// Package tempconv performs Celsius and Fahrenheit temperature computations.

// Celsius 与 Fahrenheit 是不同的数据类型,不可以相互比较和混在一个表达式中,需要提供显示转换CToF()和FToC()

// 类型转换改变了语义(编译器:语义规则,语法分析)

// package tempconv
package chapter2_5

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func TypeTest() {

	// 内置运算
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	// fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch

	// == 与 <
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)          // "true"
	fmt.Println(f >= 0)          // "true"
	fmt.Println(c == Celsius(f)) // "true"!
	// fmt.Println(c == f)          // compile error: type mismatch

	// %g:浮点数
	// %s:字符串
	// %v:自然形式

	c2 := FToC(212.0)
	fmt.Println(c2.String()) // "100°C"
	fmt.Printf("%v\n", c2)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c2)   // "100°C"
	fmt.Println(c2)          // "100°C"
	fmt.Printf("%g\n", c2)   // "100"; does not call String
	fmt.Println(float64(c2)) // "100"; does not call String
}
