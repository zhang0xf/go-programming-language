package chapter6_4

import (
	"fmt"
	"math"
)

// 我们经常选择一个方法，并且在同一个表达式里执行，比如常见的p.Distance()形式，实际上将其分成两步来执行也是可能的。(注:将p.Distance()分两步执行)
// p.Distance叫作“选择器”，选择器会返回一个方法"值"->一个将方法(Point.Distance)绑定到特定接收器变量的函数。
// 这个函数可以不通过指定其接收器即可被调用；即调用时不需要指定接收器，只要传入函数的参数即可：

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func MethodValue() {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance        // method value 方法值
	fmt.Println(distanceFromP(q))      // "5"
	var origin Point                   // {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.23606797749979", sqrt(5)

	scaleP := p.ScaleBy // method value 方法值
	scaleP(2)           // p becomes (2, 4)
	scaleP(3)           //    then (6, 12)
	scaleP(10)          //    then (60, 120)
}
