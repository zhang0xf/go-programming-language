package chapter7_2

// 接口类型

// 接口类型具体描述了一系列方法的集合，一个实现了这些方法的具体类型是这个接口类型的实例。

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// 到现在你可能注意到了很多Go语言中单方法接口的命名习惯

// 我们发现有些新的接口类型通过组合已有的接口来定义。

type ReadWriter interface {
	Writer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// 上面用到的语法和结构内嵌相似，我们可以用这种方式以一个简写命名一个接口，而不用声明它所有的方法。这种方式称为接口内嵌。
// 尽管略失简洁，我们可以像下面这样，不使用内嵌来声明io.ReadWriter接口:

type ReadWriter2 interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

// 或者甚至使用一种混合的风格：

type ReadWriter3 interface {
	Read(p []byte) (n int, err error)
	Writer
}
