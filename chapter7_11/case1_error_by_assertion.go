package chapter7_11

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

// 基于类型断言区别错误类型

// I/O可以因为任何数量的原因失败，但是有三种经常的错误必须进行不同的处理：文件已经存在（对于创建操作），找不到文件（对于读取操作），和权限拒绝。
// os包中提供了三个帮助函数来对给定的错误值表示的失败进行分类：
// func IsExist(err error) bool
// func IsNotExist(err error) bool
// func IsPermission(err error) bool

// 一个可靠的方式是使用一个专门的类型来描述结构化的错误。
// os包中定义了一个PathError类型来描述在文件路径操作中涉及到的失败，像Open或者Delete操作；并且定义了一个叫LinkError的变体来描述涉及到两个文件路径的操作，像Symlink和Rename。
// 源码片段:
// PathError records an error and the operation and file path that caused it.
type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

// 文件路径错误举例:
func FileOpenError() {
	_, err := os.Open("/no/such/file")
	fmt.Println(err) // "open /no/such/file: No such file or directory"
	fmt.Printf("%#v\n", err)
	// Output:
	// &os.PathError{Op:"open", Path:"/no/such/file", Err:0x2}
}

// 源码片段:
var ErrNotExist = errors.New("file does not exist")

// IsNotExist returns a boolean indicating whether the error is known to
// report that a file or directory does not exist. It is satisfied by
// ErrNotExist as well as some syscall errors.
func IsNotExist(err error) bool {
	if pe, ok := err.(*PathError); ok { // 断言错误类型
		err = pe.Err
	}
	return err == syscall.ENOENT || err == ErrNotExist
}

// IsNotExist()应用:
// 它会报出是否一个错误和syscall.ENOENT(§7.8)或者和有名的错误os.ErrNotExist相等(可以在§5.4.2中找到io.EOF）；或者是一个*PathError
func IsNotExistError() {
	_, err := os.Open("/no/such/file")
	fmt.Println(os.IsNotExist(err)) // "true"
}

// 如果错误消息结合成一个更大的字符串，当然PathError的结构就不再为人所知，例如通过一个对fmt.Errorf函数的调用。
// 区别错误通常必须在失败操作后，错误传回调用者前进行。
