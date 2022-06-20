package chapter3_5

import (
	"fmt"
	"unicode/utf8"
)

func Utf8Sample() {
	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	n := 0
	for _, _ = range s {
		n++
	}
	fmt.Printf("n = %d\n", n)
	fmt.Printf("n = %d\n", utf8.RuneCountInString(s))
}
