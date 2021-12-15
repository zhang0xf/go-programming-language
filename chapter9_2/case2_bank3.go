package chapter9_2

import "sync"

var (
	mu       sync.Mutex // guards balance
	balance3 int
)

func Deposit3(amount int) {
	mu.Lock()
	balance3 = balance3 + amount
	mu.Unlock()
}

func Balance3() int {
	mu.Lock()
	b := balance3
	mu.Unlock()
	return b
}
