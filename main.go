package main

import (
	"exercise/chapter2"
	"fmt"
)

func main() {
	// hello world
	fmt.Println("Hello, World!")

	// chapter1
	// chapter1.PrintArgsByFor()
	// chapter1.PrintArgsByRange()
	// chapter1.PrintArgsByJoin()
	// chapter1.PrintArgsDefault()

	// chapter2
	// chapter2.DupByStdin()
	// chapter2.DupByFilesInStreamMode()
	chapter2.DupByFilesOneTimeRead()
}
