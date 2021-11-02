package chapter3_5

import (
	"fmt"
	"strings"
)

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func BaseName2() {
	fmt.Println(basename2("a/b/c.go")) // "c"
	fmt.Println(basename2("c.d.go"))   // "c.d"
	fmt.Println(basename2("abc"))      // "abc"
}
