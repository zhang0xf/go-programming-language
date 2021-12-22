package chapter13_3

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 我们希望在这里实现一个自己的Equal函数，用于比较类型的值。
// 和DeepEqual函数类似的地方是它也是基于slice和map的每个元素进行递归比较，不同之处是它将nil值的slice（map类似）和非nil值但是空的slice视作相等的值。
// 基础部分的比较可以基于reflect包完成，和12.3章的Display函数的实现方法类似。
// 同样，我们也定义了一个内部函数equal，用于内部的递归比较。
// 对于每一对需要比较的x和y，equal函数首先检测它们是否都有效（或都无效），然后检测它们是否是相同的类型。
// 剩下的部分是一个巨大的switch分支，用于相同基础类型的元素比较。因为页面空间的限制，我们省略了一些相似的分支。
// 为了确保算法对于有环的数据结构也能正常退出，我们必须记录每次已经比较的变量，从而避免进入第二次的比较。
// Equal函数分配了一组用于比较的结构体，包含每对比较对象的地址（unsafe.Pointer形式保存）和类型。
// 我们要记录类型的原因是，有些不同的变量可能对应相同的地址。例如，如果x和y都是数组类型，那么x和x[0]将对应相同的地址，y和y[0]也是对应相同的地址，这可以用于区分x与y之间的比较或x[0]与y[0]之间的比较是否进行过了。

func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}

	// ...cycle check omitted (shown later)...

	// cycle check
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // identical references
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Int:
		return x.Int() == y.Int()
	case reflect.Bool:
		return x.Bool() == y.Bool()
	case reflect.String:
		return x.String() == y.String()

	// ...numeric cases omitted for brevity...

	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()
	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)
	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

		// ...struct and map cases omitted for brevity...
	}

	panic("unreachable")
}

// 和前面的建议一样，我们并不公开reflect包相关的接口，所以导出的函数需要在内部自己将变量转为reflect.Value类型。
// Equal reports whether x and y are deeply equal.
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

type comparison struct {
	x, y unsafe.Pointer
	reflect.Type
}

func UseEqual() {
	fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3}))   // "true"
	fmt.Println(Equal([]string{"foo"}, []string{"bar"})) // "false"
	fmt.Println(Equal([]string(nil), []string{}))        // "true"
	// fmt.Println(Equal(map[string]int(nil), map[string]int{})) // "true"
}

// Equal函数甚至可以处理类似12.3章中导致Display陷入死循环的带有环的数据。
// 目前,equal并没有实现struct分支,所以会panic
func UseEqualwithCircular() {
	type link struct {
		value string
		tail  *link
	}

	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	fmt.Println(Equal(a, a)) // "true"
	fmt.Println(Equal(b, b)) // "true"
	fmt.Println(Equal(c, c)) // "true"
	fmt.Println(Equal(a, b)) // "false"
	fmt.Println(Equal(a, c)) // "false"
}
