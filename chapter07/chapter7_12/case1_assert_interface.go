package chapter7_12

import "io"

// 通过类型断言查询接口

// 使用场景(略)
// 我们不能对任意io.Writer类型的变量w，假设它也拥有WriteString方法。
// 但是我们可以定义一个只有这个方法的新接口并且使用类型断言来检测是否w的动态类型满足这个新接口。

// writeString writes s to w.
// If w has a WriteString method, it is invoked instead of w.Write.
func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	if sw, ok := w.(stringWriter); ok { // 类型断言:检测是否有可用的WriteString方法(实践中,接口类型很少意外巧合地被实现)
		return sw.WriteString(s) // avoid a copy
	}
	return w.Write([]byte(s)) // allocate temporary copy
}

func writeHeader(w io.Writer, contentType string) error {
	if _, err := writeString(w, "Content-Type: "); err != nil {
		return err
	}
	if _, err := writeString(w, contentType); err != nil {
		return err
	}
	// ...
	return nil
}

// 为了避免重复定义，我们将这个检查移入到一个实用工具函数writeString中，但是它太有用了以致于标准库将它作为io.WriteString函数提供。
// 上面的writeString函数使用一个类型断言来获知一个普遍接口类型的值是否满足一个更加具体的接口类型；并且如果满足，它会使用这个更具体接口的行为。
// 这也是fmt.Fprintf函数怎么从其它所有值中区分满足error或者fmt.Stringer接口的值。
