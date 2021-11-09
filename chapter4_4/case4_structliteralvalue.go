package chapter4_4

import (
	"fmt"
	"image/gif"
)

type Point struct{ X, Y int }

// 结构体字面值

func StrutLiteralValue() {
	p := Point{1, 2}
	fmt.Printf("%T", p)

	// 其实更常用的是第二种写法，以成员名字和相应的值来初始化，可以包含部分或全部的成员
	nframes := 1
	anim := gif.GIF{LoopCount: nframes}
	fmt.Printf("%T", anim)

	// 因为结构体通常通过指针处理，可以用下面的写法来创建并初始化一个结构体变量，并返回结构体的地址
	pp := &Point{1, 2}
	fmt.Printf("%T", pp)
}
