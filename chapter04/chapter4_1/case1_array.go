package chapter4_1

import "fmt"

func Array() {
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// 零值
	var q [3]int = [3]int{1, 2, 3}
	fmt.Println(q[2]) // "3"
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	// 省略号
	q1 := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q1) // "[3]int"

	// 长度是类型的组成部分
	q2 := [3]int{1, 2, 3}
	// q2 = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int
	fmt.Printf("%T\n", q2)

	type Currency int

	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

	fmt.Println(RMB, symbol[RMB]) // "3 ￥"

	// 初始化(100个元素)
	r1 := [...]int{99: -1}
	fmt.Printf("%T\n", r1)

	// 数组比较
	a2 := [2]int{1, 2}
	b2 := [...]int{1, 2}
	c2 := [2]int{1, 3}
	fmt.Println(a2 == b2, a2 == c2, b2 == c2) // "true false false"
	d2 := [3]int{1, 2}
	// fmt.Println(a2 == d2) // compile error: cannot compare [2]int == [3]int
	fmt.Printf("%T\n", d2)
}
