package chapter6_1_1

import (
	"exercise/chapter6_1"
	"fmt"
)

func OutGeometry() {
	perim := chapter6_1.Path{{X: 1, Y: 1}, {X: 5, Y: 1}, {X: 5, Y: 4}, {X: 1, Y: 1}}
	fmt.Println(chapter6_1.PathDistance(perim)) // "12", standalone function
	fmt.Println(perim.Distance())               // "12", method of geometry.Path
}
