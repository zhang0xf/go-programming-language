package chapter1

import (
	"fmt"
	"os"
)

func PrintArgsDefault() {
	fmt.Println(os.Args[1:])
}
