package chapter4_4

import "fmt"

// 结构体比较

type address struct {
	hostname string
	port     int
}

func StructCompare() {
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
	fmt.Println(p == q)                   // "false"

	// 可比较的结构体类型和其他可比较的类型一样，可以用于map的key类型。
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
}
