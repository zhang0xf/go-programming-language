package chapter6_3

import "image/color"

// 一个struct类型也可能会有多个匿名字段。我们将ColoredPoint定义为下面这样：
// 然后这种类型的值便会拥有Point和RGBA类型的所有方法，以及直接定义在ColoredPoint中的方法。
// 当编译器解析一个选择器到方法时，比如p.ScaleBy，它会首先去找直接定义在这个类型里的ScaleBy方法，然后找被ColoredPoint的内嵌字段们引入的方法，然后去找Point和RGBA的内嵌字段引入的方法，然后一直递归向下找。如果选择器有二义性的话编译器会报错，比如你在同一级里有两个同名的方法。
type ColoredPoint3 struct {
	Point
	color.RGBA
}
