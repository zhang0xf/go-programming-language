package chapter4_2

import "fmt"

func AppendRune() {
	var runes []rune
	// []rune("hello,世界")
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"
}
