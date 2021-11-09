package chapter4_4

import "fmt"

// Go语言有一个特性让我们只声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫匿名成员。
// 匿名成员并不是真的无法访问了,其中匿名成员Circle和Point都有自己的名字——就是命名的类型名字——但是这些名字在点操作符中是可选的。我们在访问子成员的时候可以忽略任何匿名成员部分。

func StructNoName() {
	type Circle struct {
		Point  // 匿名
		Radius int
	}

	type Wheel struct {
		Circle // 匿名
		Spokes int
	}

	var w Wheel
	w.X = 8      // equivalent to w.Circle.Point.X = 8
	w.Y = 8      // equivalent to w.Circle.Point.Y = 8
	w.Radius = 5 // equivalent to w.Circle.Radius = 5
	w.Spokes = 20

	// 结构体字面值与匿名成员
	// w = Wheel{8, 8, 5, 20}                       // compile error: unknown fields
	// w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields

	// 结构体字面值必须遵循形状类型声明时的结构
	// 方式1:
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	// 方式2:
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:8, Y:8}, Radius:5}, Spokes:20}

	w.X = 42

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:42, Y:8}, Radius:5}, Spokes:20}

	// 需要注意的是Printf函数中%v参数包含的#副词，它表示用和Go语言类似的语法打印值。对于结构体类型来说，将包含每个成员的名字。

	// 因为匿名成员也有一个隐式的名字，因此不能同时包含两个类型相同的匿名成员，这会导致名字冲突。

	// 匿名成员也有可见性的规则约束。在上面的例子中，Point和Circle匿名成员都是导出的。即使它们不导出（比如改成小写字母开头的point和circle），我们依然可以用简短形式访问匿名成员嵌套的成员
	// w.X = 8 // equivalent to w.circle.point.X = 8
	// 但是在包外部，因为circle和point没有导出，不能访问它们的成员，因此简短的匿名成员访问语法也是禁止的。
}
