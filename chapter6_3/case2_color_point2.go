package chapter6_3

import (
	"fmt"
	"image/color"
)

// 在类型中内嵌的匿名字段也可能是一个命名类型的指针，这种情况下字段和方法会被间接地引入到当前的类型中
// 添加这一层间接关系让我们可以共享通用的结构并动态地改变对象之间的关系
type ColoredPoint2 struct {
	*Point
	Color color.RGBA
}

// 匿名字段指针
func ColorPoint3() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint2{&Point{1, 1}, red}
	q := ColoredPoint2{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point)) // "5"
	q.Point = p.Point                 // p and q now share the same Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
}
