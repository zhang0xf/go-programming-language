package chapter3_5

import "fmt"

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename1(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func BaseName1() {
	fmt.Println(basename1("a/b/c.go")) // "c"
	fmt.Println(basename1("c.d.go"))   // "c.d"
	fmt.Println(basename1("abc"))      // "abc"
}
