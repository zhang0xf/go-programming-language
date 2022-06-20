package chapter13_4

// 通过cgo调用C代码
// Go程序可能会遇到要访问C语言的某些硬件驱动函数的场景，或者是从一个C++语言实现的嵌入式数据库查询记录的场景，或者是使用Fortran语言实现的一些线性代数库的场景。
// C语言作为一个通用语言，很多库会选择提供一个C兼容的API，然后用其他不同的编程语言实现。

// 在本节中，我们将构建一个简易的数据压缩程序，使用了一个Go语言自带的叫cgo的用于支援C语言函数调用的工具。
// 这类工具一般被称为 foreign-function interfaces （简称ffi）, 并且在类似工具中cgo也不是唯一的。
// SWIG（ http://swig.org ）是另一个类似的且被广泛使用的工具，SWIG提供了很多复杂特性以支援C++的特性，但SWIG并不是我们要讨论的主题。

// 在标准库的compress/...子包有很多流行的压缩算法的编码和解码实现，包括流行的LZW压缩算法（Unix的compress命令用的算法）和DEFLATE压缩算法（GNU gzip命令用的算法）。
// 这些包的API的细节虽然有些差异，但是它们都提供了针对 io.Writer类型输出的压缩接口和提供了针对io.Reader类型输入的解压缩接口。例如：

// package gzip // compress/gzip
// func NewWriter(w io.Writer) io.WriteCloser
// func NewReader(r io.Reader) (io.ReadCloser, error)

// bzip2压缩算法，是基于优雅的Burrows-Wheeler变换算法，运行速度比gzip要慢，但是可以提供更高的压缩比。
// 标准库的compress/bzip2包目前还没有提供bzip2压缩算法的实现。
// 完全从头开始实现一个压缩算法是一件繁琐的工作，而且 http://bzip.org 已经有现成的libbzip2的开源实现，不仅文档齐全而且性能又好。

// 如果是比较小的C语言库，我们完全可以用纯Go语言重新实现一遍。
// 如果我们对性能也没有特殊要求的话，我们还可以用os/exec包的方法将C编写的应用程序作为一个子进程运行。
// 只有当你需要使用复杂而且性能更高的底层C接口时，就是使用cgo的场景了。下面我们将通过一个例子讲述cgo的具体用法。(见case2_bzip.go)

// 要使用libbzip2，我们需要先构建一个bz_stream结构体，用于保持输入和输出缓存。
// 。然后有三个函数：
// BZ2_bzCompressInit用于初始化缓存，
// BZ2_bzCompress用于将输入缓存的数据压缩到输出缓存，
// BZ2_bzCompressEnd用于释放不需要的缓存。
// 我们可以在Go代码中直接调用BZ2_bzCompressInit和BZ2_bzCompressEnd，但是对于BZ2_bzCompress，我们将定义一个C语言的包装函数，用它完成真正的工作。
// (见case2_bzip.c)

// 现在让我们转到Go语言部分，(见case2_bzip.go)。其中import "C"的语句是比较特别的。其实并没有一个叫C的包，但是这行语句会让Go编译程序在编译之前先运行cgo工具。
// 在预处理过程中，cgo工具生成一个临时包用于包含所有在Go语言中访问的C语言的函数或类型。例如C.bz_stream和C.BZ2_bzCompressInit。
// 在cgo注释中还可以包含#cgo指令，用于给C语言工具链指定特殊的参数。
// 例如CFLAGS和LDFLAGS分别对应传给C语言编译器的编译参数和链接器参数，使它们可以从特定目录找到bzlib.h头文件和libbz2.a库文件。
// 这个例子假设你已经在/usr目录成功安装了bzip2库。如果bzip2库是安装在不同的位置，你需要更新这些参数

// 我们演示了如何将一个C语言库链接到Go语言程序。
// 相反, 将Go编译为静态库然后链接到C程序，或者将Go程序编译为动态库然后在C程序中动态加载也都是可行的。
// 这里我们只展示的cgo很小的一些方面，更多的关于内存管理、指针、回调函数、中断信号处理、字符串、errno处理、终结器，以及goroutines和系统线程的关系等，有很多细节可以讨论。
// 特别是如何将Go语言的指针传入C函数的规则也是异常复杂的，部分的原因在13.2节有讨论到，但是在Go1.5中还没有被明确。
// 如果要进一步阅读，可以从 https://golang.org/cmd/cgo 开始。
