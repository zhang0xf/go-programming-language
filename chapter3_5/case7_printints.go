package chapter3_5

import (
	"bytes"
	"fmt"
)

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	// 添加ASCII时,推荐使用WriteByte()
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')

	// 添加任意字符的utf8编码时,最好使用WriteRune()
	s := "你好,世界"
	r := []rune(s)
	for _, v := range r {
		buf.WriteRune(v)
	}

	return buf.String()
}

func PrintInts() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}
