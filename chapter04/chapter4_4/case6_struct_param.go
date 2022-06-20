package chapter4_4

import "fmt"

// 结构体可以作为函数的参数和返回值。

func StructParam() {
	fmt.Println(Scale(Point{1, 2}, 5)) // "{5 10}"
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

// 较大结构体使用指针传入
func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

// 参数是"值"传入,需要使用指针以改变原始变量
func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}
