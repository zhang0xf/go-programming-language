package chapter13_2

import (
	"fmt"
	"unsafe"
)

// 大多数指针类型会写成*T，表示是“一个指向T类型变量的指针”。
// unsafe.Pointer是特别定义的一种指针类型（译注：类似C语言中的void*类型的指针），它可以包含任意类型变量的地址。
// 当然，我们不可以直接通过*p来获取unsafe.Pointer指针指向的真实变量的值，因为我们并不知道变量的具体类型。
// 和普通指针一样，unsafe.Pointer指针也是可以比较的，并且支持和nil常量比较判断是否为空指针。

// 一个普通的*T类型指针可以被转化为unsafe.Pointer类型指针，并且一个unsafe.Pointer类型指针也可以被转回普通的指针，被转回普通的指针类型并不需要和原始的*T类型相同。
// 通过将*float64类型指针转化为*uint64类型指针，我们可以查看一个浮点数变量的位模式。

func Float64bits(f float64) uint64 { return *(*uint64)(unsafe.Pointer(&f)) }

func UnsafePointer() {
	fmt.Printf("%#016x\n", 1.0)              // "00000x1.0000p+00"
	fmt.Printf("%#016x\n", Float64bits(1.0)) // "0x3ff0000000000000"
}

// 一个unsafe.Pointer指针也可以被转化为uintptr类型，然后保存到指针型数值变量中，然后用以做必要的指针数值运算。
// （第三章内容，uintptr是一个无符号的整型数，足以保存一个地址）这种转换虽然也是可逆的，但是将uintptr转为unsafe.Pointer指针可能会破坏类型系统，因为并不是所有的数字都是有效的内存地址。

// 许多将unsafe.Pointer指针转为原生数字，然后再转回为unsafe.Pointer类型指针的操作也是不安全的。
// 比如下面的例子需要将变量x的地址加上b字段地址偏移量转化为*int16类型指针，然后通过该指针更新x.b：

var x struct {
	a bool
	b int16
	c []int
}

func UnsafePointer2() {
	// 和 pb := &x.b 等价
	pb := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"
}

// 上面的写法尽管很繁琐，但在这里并不是一件坏事，因为这些功能应该很谨慎地使用。
