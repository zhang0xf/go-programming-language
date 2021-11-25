package chapter7_3

import (
	"bytes"
	"fmt"
	"io"
)

// 因为接口与实现只依赖于判断两个类型的方法，所以没有必要定义一个具体类型和它实现的接口之间的关系。
// 也就是说，尝试文档化和断言这种关系几乎没有用，所以并没有通过程序强制定义。

func InterfaceDetail3() {
	// 下面的定义在编译期断言一个*bytes.Buffer的值实现了io.Writer接口类型:
	// *bytes.Buffer must satisfy io.Writer
	var w io.Writer = new(bytes.Buffer)

	// 因为任意*bytes.Buffer的值，甚至包括nil通过(*bytes.Buffer)(nil)进行显示的转换都实现了这个接口，所以我们不必分配一个新的变量。
	// 并且因为我们绝不会引用变量w，我们可以使用空标识符来进行代替。
	// *bytes.Buffer must satisfy io.Writer
	var _ io.Writer = (*bytes.Buffer)(nil)
	fmt.Printf("%T", w)
}
