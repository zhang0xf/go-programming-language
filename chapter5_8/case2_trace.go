package chapter5_8

import (
	"log"
	"time"
)

// 调试复杂程序时，defer机制也常被用于记录何时进入和退出函数。

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses(括弧表示函数调用!trace的返回值为一个函数,注:延迟调用的是返回的这个函数)
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
