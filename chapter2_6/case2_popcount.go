// package popcount

package chapter2_6

// pc[i] is the population count of i.
var pc [256]byte

// 生成表格:p[i]存放i这个数中有多少个1

// 为什么是256? 因为2 ^ 8 = 256,涵盖了byte(8bit)的所有情况

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func init() {
	// 可以有多个init()
}

// pc[i] is the population count of i.
// var pc [256]byte = func() (pc [256]byte) {
// 	for i := range pc {
// 		pc[i] = pc[i/2] + byte(i&1)
// 	}
// 	return
// }()
// init可以使用匿名函数替代,最后的括号表示函数调用

// PopCount returns the population count (number of set bits) of x.
// 匹配表格
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
