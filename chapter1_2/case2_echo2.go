package chapter1_2

import (
	"fmt"
	"os"
)

func PrintArgsByRange() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
