package chapter3_1

import "fmt"

// 如果选择uint,那么对于下面的逆序循环是灾难性的,因为i >= 0将永远为真!

func ReverseLoop() {
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}
}

// panic
