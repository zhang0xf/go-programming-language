package chapter3_5

import (
	"fmt"
	"unicode"
)

// 使用rune的理由:大小相同
// utf-8是变长编码,每个字符所占用的字节不定

func RuneSample() {
	// "program" in Japanese katakana
	s := "プログラム"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"

	fmt.Println(string(r)) // "プログラム"

	fmt.Println(string(65))      // "A", not "65"
	fmt.Println(string(0x4eac))  // "京"
	fmt.Println(string(1234567)) // "?"

	// unicode包测试
	for i, v := range r {
		if unicode.IsDigit(v) {
			fmt.Printf("%d是数字\n", i)
		}
		if unicode.IsLetter(v) {
			fmt.Printf("%d是字母\n", i)
		}
	}
}
