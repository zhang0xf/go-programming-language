package chapter4_2

import "fmt"

// 压栈
// stack = append(stack, v) // push v

// 栈顶元素
// top := stack[len(stack)-1] // top of stack

// 出栈
// stack = stack[:len(stack)-1] // pop

// 移除栈中间元素(保持有序)
func remove1(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func Remove1() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove1(s, 2)) // "[5 6 8 9]"
}

// 移除栈中间元素(不用保持有序)
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func Remove2() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove2(s, 2)) // "[5 6 9 8]
}
