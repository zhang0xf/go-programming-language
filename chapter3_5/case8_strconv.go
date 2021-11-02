package chapter3_5

import (
	"fmt"
	"strconv"
)

func StrConv() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"

	// 使用二进制格式化x
	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"

	// fmt有时比FormatInt()更方便
	// s := fmt.Sprintf("x=%b", x) // "x=1111011"

	// x1, err := strconv.Atoi("123")             // x is an int
	// y1, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
}
