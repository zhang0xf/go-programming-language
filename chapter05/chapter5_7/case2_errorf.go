package chapter5_7

import (
	"fmt"
	"os"
)

func Errorf() {
	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name) // "Line 12: undefined: count"
}

// interface{}表示函数的最后一个参数可以接收任意类型，我们会在第7章详细介绍。
func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
