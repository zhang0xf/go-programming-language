package chapter5_7

import "fmt"

// 可变参数
// 在声明可变参数函数时，需要在参数列表的最后一个参数类型之前加上省略符号“...”，这表示该函数会接收任意数量的该类型参数。

func Sum() {
	fmt.Println(sum())           // "0"
	fmt.Println(sum(3))          // "3"
	fmt.Println(sum(1, 2, 3, 4)) // "10"
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

// 在上面的代码中，调用者隐式的创建一个数组，并将原始参数复制到数组中，再把数组的一个切片作为参数传给被调用函数。
// 如果原始参数已经是切片类型，我们该如何传递给sum？只需在最后一个参数后加上省略符。

func Sum2() {
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
}

func f(...int) {}
func g([]int)  {}

// 虽然在可变参数函数内部，...int 型参数的行为看起来很像切片类型，但实际上，可变参数函数和以切片作为参数的函数是不同的。
func DiffBetween() {
	fmt.Printf("%T\n", f) // "func(...int)"
	fmt.Printf("%T\n", g) // "func([]int)"
}
