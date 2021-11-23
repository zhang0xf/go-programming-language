package chapter5_9

import "fmt"

// 当f(0)被调用时，发生panic异常，之前被延迟执行的3个fmt.Printf被调用。
func Panic2() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
