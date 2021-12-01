package chapter7_4

import (
	"flag"
	"fmt"
)

// usgae : ./exercise
//         ./exercise -temp 18C
//         ./exercise -temp 212F

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func TempFlag() {
	flag.Parse()
	fmt.Println(*temp)
}
