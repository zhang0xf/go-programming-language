package chapter7_3

import (
	"fmt"
)

type IntSet struct {
	/* ... */
}

func (*IntSet) String() string

func InterfaceDetail() {
	// IntSet类型的String方法的接收者是一个指针类型，所以我们不能在一个不能寻址的IntSet值上调用这个方法：
	// var _ = IntSet{}.String() // compile error: String requires *IntSet receiver

	// 但是我们可以在一个IntSet值上调用这个方法：
	var s IntSet
	var _ = s.String() // OK: s is a variable and &s has a String method

	// 然而，由于只有*IntSet类型有String方法，所以也只有*IntSet类型实现了fmt.Stringer接口：
	var _ fmt.Stringer = &s // OK
	// var _ fmt.Stringer = s  // compile error: IntSet lacks String method
}
