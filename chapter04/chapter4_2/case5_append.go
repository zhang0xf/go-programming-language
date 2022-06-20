package chapter4_2

import "fmt"

// 内置的append函数则可以追加多个元素，甚至追加一个slice
func Append() {
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) // append the slice x
	fmt.Println(x)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"
}
