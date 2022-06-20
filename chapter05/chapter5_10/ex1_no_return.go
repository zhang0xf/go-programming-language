package chapter5_10

import "fmt"

// 练习5.19：使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。

func ReturnN() {
	a := returnN()
	fmt.Println(a)
}

func returnN() (result int) {
	defer func() {
		if p := recover(); p != nil {
			result = p.(int)
		}
	}()
	panic(3)
}
