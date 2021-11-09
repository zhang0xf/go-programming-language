package chapter4_4

// 结构体嵌入

type Circle struct {
	X, Y, Radius int
}

type Wheel struct {
	X, Y, Radius, Spokes int
}

func StructEmbed() {
	var w1 Wheel
	w1.X = 8
	w1.Y = 8
	w1.Radius = 5
	w1.Spokes = 20

	// 将相同的属性独立出来,使结构体清晰,但是访问成员变得繁琐
	type Point struct {
		X, Y int
	}

	type Circle struct {
		Center Point
		Radius int
	}

	type Wheel struct {
		Circle Circle
		Spokes int
	}

	var w2 Wheel
	w2.Circle.Center.X = 8
	w2.Circle.Center.Y = 8
	w2.Circle.Radius = 5
	w2.Spokes = 20
}
