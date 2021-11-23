package chapter6_2

import (
	"fmt"
	"math"
)

// 基于指针对象的方法:避免大对象的值拷贝

// 我们用来更新接收器的对象的方法，当这个接受者变量本身比较大时，我们就可以用其指针而不是对象来声明方法

type Point struct{ X, Y float64 }

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 这个方法的名字是(*Point).ScaleBy。这里的括号是必须的；没有括号的话这个表达式可能会被理解为*(Point.ScaleBy)。
// 在现实的程序里，一般会约定如果Point这个类有一个指针作为接收器的方法，那么所有Point的方法都必须有一个指针接收器，即使是那些并不需要这个指针接收器的函数。

// 只有类型(Point)和指向他们的指针(*Point)，才可能是出现在接收器声明里的两种接收器。
// 在声明方法时，如果一个类型名本身是一个指针的话，是不允许其出现在接收器中的，比如下面这个例子：

type P *int

// func (P) f() { /* ... */ } // compile error: invalid receiver type

// 调用
func CallFuncByPointer() {
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r) // "{2, 4}"

	// 或者

	p := Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p) // "{2, 4}"

	// 或者

	p1 := Point{1, 2}
	(&p1).ScaleBy(2)
	fmt.Println(p1) // "{2, 4}"

	// 或者
	p.ScaleBy(2)
	// 编译器会隐式地帮我们用&p去调用ScaleBy这个方法。
	// 这种简写方法只适用于“变量”，包括struct里的字段比如p.X，以及array和slice内的元素比如perim[0]。
	// 我们不能通过一个无法取到地址的接收器来调用指针方法，比如临时变量的内存地址就无法获取得到：
	// Point{1, 2}.ScaleBy(2) // compile error: can't take address of Point literal
}

// 调用
func CallFuncByObjs() {
	var pptr *Point
	var q Point

	pptr.Distance(q)
	(*pptr).Distance(q)
	// 我们可以通过地址来找到这个变量，只要用解引用符号*来取到该变量即可。编译器在这里也会给我们隐式地插入*这个操作符

}

// 调用
func CallFuncByBoth() {
	var pptr *Point
	var q Point
	p := Point{1, 2}

	// 接收器的实际参数和其形式参数是相同
	Point{1, 2}.Distance(q) //  Point 临时变量
	pptr.ScaleBy(2)         // *Point 指针对象

	// 接收器实参是类型T，但接收器形参是类型*T
	p.ScaleBy(2) // implicit (&p) 隐式转换

	// 接收器实参是类型*T，形参是类型T
	pptr.Distance(q) // implicit (*pptr) 隐式转换
}

// 译注： 作者这里说的比较绕，其实有两点：

// 1.不管你的method的receiver是指针类型还是非指针类型，都是可以通过指针/非指针类型进行调用的，编译器会帮你做类型转换。
// 2.在声明一个method的receiver该是指针还是非指针类型时，你需要考虑两方面的因素:
//   第一方面是这个对象本身是不是特别大，如果声明为非指针变量时，调用会产生一次拷贝；
//   第二方面是如果你用指针类型作为receiver，那么你一定要注意，这种指针类型指向的始终是一块内存地址，就算你对其进行了拷贝。熟悉C或者C++的人这里应该很快能明白。
