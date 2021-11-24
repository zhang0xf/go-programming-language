package chapter6_4

import "fmt"

// 当T是一个类型时，方法表达式可能会写作T.f或者(*T).f，会返回一个函数"值"，这种函数会将其第一个参数用作接收器，所以可以用通常的方式来对其进行调用：

func MethodExp() {
	p := Point{1, 2}
	q := Point{4, 6}

	distance := Point.Distance   // method expression 方法表达式(类型.方法)
	fmt.Println(distance(p, q))  // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)            // "{2 4}"
	fmt.Printf("%T\n", scale) // "func(*Point, float64)"
}
