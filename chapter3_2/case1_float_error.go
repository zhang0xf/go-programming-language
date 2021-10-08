package chapter3_2

import (
	"fmt"
	"math"
)

// float有效位数只有23bit(精度),剩余用于指数和符号

func FloatError() {
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!
}

func FloatPrint() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}

// result:
// x = 0       e^x =    1.000
// x = 1       e^x =    2.718
// x = 2       e^x =    7.389
// x = 3       e^x =   20.086
// x = 4       e^x =   54.598
// x = 5       e^x =  148.413
// x = 6       e^x =  403.429
// x = 7       e^x = 1096.633

func FloatIllegalValue() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"
}

// NaN和任何数都不相等(包括它自己)
func FloatNaN() {
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"
}
