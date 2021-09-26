// Package tempconv performs Celsius and Fahrenheit temperature computations.

// Celsius 与 Fahrenheit 是不同的数据类型,不可以相互比较和混在一个表达式中,需要提供显示转换CToF()和FToC()

package tempconv

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
