package chapter1_2

import (
	"fmt"
	"os"
	"strings"
)

func PrintArgsByJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
