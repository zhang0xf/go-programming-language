package chapter3_5

import "fmt"

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func Comma() {
	// 字符串与字节slice([]byte)的转换
	s := "abc"
	b := []byte(s)
	s2 := string(b)

	fmt.Printf("%s\n", s2)
}
