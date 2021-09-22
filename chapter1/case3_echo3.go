package chapter1

import (
	"fmt"
	"os"
	"strings"
)

func PrintArgsByJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
