package chapter7_10

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// 类型断言是一个使用在接口值(动态类型和动态值)上的操作。(注:w是一个接口类型)
// 语法上它看起来像x.(T)被称为断言类型，这里x表示一个接口的类型和T表示一个类型。
// 一个类型断言检查它操作对象的动态类型是否和断言的类型匹配。

// 如果断言的类型T是一个具体类型，然后类型断言检查x的动态类型是否和T相同。
// 如果这个检查成功了，类型断言的结果是x的动态值，当然它的类型是T。
// 如果检查失败，接下来这个操作会抛出panic。
func Assertion() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)            // success: f == os.Stdout 断言的类型是:具体的类型
	fmt.Printf("%T\n", f)        // *os.File
	f1 := w.(io.ReadWriteCloser) // 断言的类型是:接口类型
	fmt.Printf("%T\n", f1)       // *os.File
	c := w.(*bytes.Buffer)       // panic: interface holds *os.File, not *bytes.Buffer
	fmt.Printf("%T\n", c)
}

// 如果相反地断言的类型T是一个接口类型，然后类型断言检查是否x的动态类型满足T。
// 对一个接口类型的类型断言改变了类型的表述方式，改变了可以获取的方法集合（通常更大），但是它保留了接口值内部的动态类型和值的部分。
// 在下面的第一个类型断言后，w和rw都持有os.Stdout，因此它们都有一个动态类型*os.File，但是变量w是一个io.Writer类型，只对外公开了文件的Write方法，而rw变量还公开了它的Read方法。
func Assertion2() {
	var w io.Writer
	w = os.Stdout
	fmt.Printf("%T\n", w)   // *os.File
	rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	fmt.Printf("%T\n", rw)  // *os.File
	// w = new(ByteCounter)
	rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method
}

// 如果断言操作的对象是一个nil接口值，那么不论被断言的类型是什么这个类型断言都会失败。
// 我们几乎不需要对一个更少限制性的接口类型（更少的方法集合）做断言，因为它表现的就像是赋值操作一样，除了对于nil接口值的情况。?
// 举例:w是一个write接口,将w断言成ReadWriter接口,可行.因为动态类型是*os.File,动态值是os.Stdout,不过此行为表现的像赋值.
func Assertion3() {
	var w io.Writer
	w = os.Stdout
	fmt.Printf("%T\n", w)   // *os.File
	rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	w = rw                  // io.ReadWriter is assignable to io.Writer
	w = rw.(io.Writer)      // fails only if rw == nil
	fmt.Printf("%T\n", w)   // *os.File
	fmt.Printf("%T\n", rw)  // *os.File
}

// 经常地，对一个接口值的动态类型我们是不确定的，并且我们更愿意去检验它是否是一些特定的类型。
// 如果这个操作失败了，那么ok就是false值，第一个结果等于被断言类型的零值，在这个例子中就是一个nil的*bytes.Buffer类型。
func Assertion4() {
	var w io.Writer = os.Stdout
	f, ok := w.(*os.File) // success:  ok, f == os.Stdout
	if ok {
		fmt.Printf("%T", f)
	}
	b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil
	if ok {
		fmt.Printf("%T", b)
	}
}
