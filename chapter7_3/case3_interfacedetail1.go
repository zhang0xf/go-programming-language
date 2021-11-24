package chapter7_3

import (
	"io"
	"os"
)

// 就像信封封装和隐藏起信件来一样，接口类型封装和隐藏具体类型和它的值。即使具体类型有其它的方法，也只有接口类型暴露出来的方法会被调用到：
func InterfaceDetail1() {
	os.Stdout.Write([]byte("hello")) // OK: *os.File has Write method
	os.Stdout.Close()                // OK: *os.File has Close method

	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello")) // OK: io.Writer has Write method
	// w.Close()                // compile error: io.Writer lacks Close method
}
