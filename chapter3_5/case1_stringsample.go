package chapter3_5

import "fmt"

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

}
