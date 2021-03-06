package chapter1_2

import (
	"fmt"
	"os"
)

func PrintArgsByFor() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
