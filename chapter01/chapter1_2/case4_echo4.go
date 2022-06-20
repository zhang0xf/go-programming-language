package chapter1_2

import (
	"fmt"
	"os"
)

func PrintArgsDefault() {
	fmt.Println(os.Args[1:])
}
