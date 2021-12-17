package chapter10_7

// go list命令可以查询可用包的信息。其最简单的形式，可以测试包是否在工作区并打印它的导入路径：

// $ go list github.com/go-sql-driver/mysql
// github.com/go-sql-driver/mysql

// go list命令的参数还可以用"..."表示匹配任意的包的导入路径。我们可以用它来列出工作区中的所有包：

// $ go list ...
// archive/tar
// archive/zip
// bufio
// bytes
// cmd/addr2line
// cmd/api
// ...many more...

// 或者是特定子目录下的所有包：

// $ go list gopl.io/ch3/...
// gopl.io/ch3/basename1
// gopl.io/ch3/basename2
// gopl.io/ch3/comma
// gopl.io/ch3/mandelbrot
// gopl.io/ch3/netflag
// gopl.io/ch3/printints
// gopl.io/ch3/surface

// 或者是和某个主题相关的所有包:

// $ go list ...xml...
// encoding/xml
// gopl.io/ch7/xmlselect

// go list命令还可以获取每个包完整的元信息，而不仅仅只是导入路径，这些元信息可以以不同格式提供给用户。其中-json命令行参数表示用JSON格式打印每个包的元信息。

// $ go list -json hash
// {
//     "Dir": "/home/gopher/go/src/hash",
//     "ImportPath": "hash",
//     "Name": "hash",
//     "Doc": "Package hash provides interfaces for hash functions.",
//     "Target": "/home/gopher/go/pkg/darwin_amd64/hash.a",
//     "Goroot": true,
//     "Standard": true,
//     "Root": "/home/gopher/go",
//     "GoFiles": [
//             "hash.go"
//     ],
//     "Imports": [
//         "io"
//     ],
//     "Deps": [
//         "errors",
//         "io",
//         "runtime",
//         "sync",
//         "sync/atomic",
//         "unsafe"
//     ]
// }

// 命令行参数-f则允许用户使用text/template包（§4.6）的模板语言定义输出文本的格式。
// 下面的命令将打印strconv包的依赖的包，然后用join模板函数将结果链接为一行，连接时每个结果之间用一个空格分隔：

// $ go list -f '{{join .Deps " "}}' strconv
// $ go list -f "{{join .Deps \" \"}}" strconv (windows)
// errors math runtime unicode/utf8 unsafe

// 下面的命令打印compress子目录下所有包的导入包列表：

// $ go list -f '{{.ImportPath}} -> {{join .Imports " "}}' compress/...
// $ go list -f "{{.ImportPath}} -> {{join .Imports \" \"}}" compress/... (windows)
// compress/bzip2 -> bufio io sort
// compress/flate -> bufio fmt io math sort strconv
// compress/gzip -> bufio compress/flate errors fmt hash hash/crc32 io time
// compress/lzw -> bufio errors fmt io
// compress/zlib -> bufio compress/flate errors fmt hash hash/adler32 io
