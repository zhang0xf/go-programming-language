package chapter3_5

import (
	"fmt"
	"os"
)

func StringSample() {
	s := "hello, world"
	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
	// c := s[len(s)]          // panic: index out of range

	fmt.Println(s[0:5]) // "hello"

	fmt.Println(s[:5]) // "hello"
	fmt.Println(s[7:]) // "world"
	fmt.Println(s[:])  // "hello, world"

	fmt.Println("goodbye" + s[5:]) // "goodbye, world"

	s1 := "left foot"
	t1 := s
	s1 += ", right foot"
	fmt.Println(s1) // "left foot, right foot"
	fmt.Println(t1) // "left foot"
	// s1[0] = 'L'     // compile error: cannot assign to s[0]

	// 转义序列测试
	escapeString := "hello\n\"world"
	// escapeString2 := "hello\n
	// world"	// compile error
	fmt.Println(escapeString)

	// 原生字符串测试
	nativeString := `hello\n\'world`
	fmt.Println(nativeString)

	const GoUsage = `Go is a tool for managing Go source code.

	Usage:
    	go command [arguments]
	...`

	fmt.Println(GoUsage)

	f, err := os.Create("./data/stringSample.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Create stringSample.txt error: %v\n", err)
	}
	defer f.Close()
	var str = []byte(GoUsage)
	f.Write(str)

	var nativeString2 = `hello\n\'world` + "`" + `反引号的转义`
	fmt.Println(nativeString2)
}
