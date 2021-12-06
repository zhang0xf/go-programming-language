package chapter7_5

import (
	"bytes"
	"fmt"
	"io"
)

// 一个不包含任何值的nil接口值和一个刚好包含nil指针的接口值是不同的。这个细微区别产生了一个容易绊倒每个Go程序员的陷阱。

const debug = true

func InterfaceValueAndNil() {
	var buf *bytes.Buffer // struct
	// var buf io.Writer // interface 正确写法
	fmt.Printf("%T\n", buf)
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	fmt.Printf("%T\n", out)
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n")) // panic: nil pointer dereference(when debug = false)
	}
}

// 我们可能会预计当把变量debug设置为false时可以禁止对输出的收集，但是实际上在out.Write方法调用时程序发生了panic：
// 当InterfaceValueAndNil()函数调用函数f时，它给f函数的out参数赋了一个*bytes.Buffer的空指针，所以out的动态值是nil。
// 然而，它的动态类型是*bytes.Buffer，意思就是out变量是一个包含空指针值的非空接口，所以防御性检查out!=nil的结果依然是true。
// 注释:使用一个*bytes.Buffer类型的变量去赋值out,会改变其动态类型为*bytes.Buffer,接口值为对应的值(上例为nil),与直接的out = nil赋值方式不同! 同样,out != nil需要将out的动态类型和接口值两项与nil比较.
// 动态分配机制依然决定(*bytes.Buffer).Write的方法会被调用，但是这次的接收者的值是nil。
// 对于一些如*os.File的类型，nil是一个有效的接收者(§6.2.1)，但是*bytes.Buffer类型不在这些种类中。这个方法会被调用，但是当它尝试去获取缓冲区时会发生panic。
// 解决方案就是将InterfaceValueAndNil()函数中的变量buf的类型改为io.Writer，因此可以避免一开始就将一个不完整的值赋值给这个接口
