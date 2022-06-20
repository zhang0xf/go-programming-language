package chapter10_5

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // register PNG decoder
	"io"
	"os"
)

// usage : go build chapter3_3.DrawMandelbrot()
//         ./exercise > ./data/test.PNG
//         go build chapter10_5.JPEG()
//         ./exercise < ./data/test.PNG > ./data/test.jpeg

// 标准库的image图像包包含了一个Decode函数，用于从io.Reader接口读取数据并解码图像，它调用底层注册的图像解码器来完成任务，然后返回image.Image类型的图像。
// 使用image.Decode很容易编写一个图像格式的转换工具，读取一种格式的图像，然后编码为另一种图像格式：
// 如果我们将gopl.io/ch3/mandelbrot（§3.3）的输出导入到这个程序的标准输入，它将解码输入的PNG格式图像，然后转换为JPEG格式的图像输出（图3.3）。
// 要注意image/png包的匿名导入语句。如果没有这一行语句，程序依然可以编译和运行，但是它将不能正确识别和解码PNG格式的图像：

// $ go build gopl.io/ch3/mandelbrot
// $ go build gopl.io/ch10/jpeg
// $ ./mandelbrot | ./jpeg >mandelbrot.jpg
// Input format = png

func JPEG() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

// 标准库还提供了GIF、PNG和JPEG等格式图像的解码器，用户也可以提供自己的解码器，但是为了保持程序体积较小，很多解码器并没有被全部包含，除非是明确需要支持的格式。
// image.Decode函数在解码时会依次查询支持的格式列表。每个格式驱动列表的每个入口指定了四件事情：
// 格式的名称；
// 一个用于描述这种图像数据开头部分模式的字符串，用于解码器检测识别；
// 一个Decode函数用于完成解码图像工作；
// 一个DecodeConfig函数用于解码图像的大小和颜色空间的信息。
// 每个驱动入口是通过调用image.RegisterFormat函数注册，一般是在每个格式包的init初始化函数中调用，例如image/png包是这样注册的：

// package png // image/png

// func Decode(r io.Reader) (image.Image, error)
// func DecodeConfig(r io.Reader) (image.Config, error)

// func init() {
//     const pngHeader = "\x89PNG\r\n\x1a\n"
//     image.RegisterFormat("png", pngHeader, Decode, DecodeConfig)
// }

// 最终的效果是，主程序只需要匿名导入特定图像驱动包就可以用image.Decode解码对应格式的图像了。
// 数据库包database/sql也是采用了类似的技术，让用户可以根据自己需要选择导入必要的数据库驱动。例如：

// import (
//     "database/sql"
//     _ "github.com/lib/pq"              // enable support for Postgres
//     _ "github.com/go-sql-driver/mysql" // enable support for MySQL
// )

// db, err = sql.Open("postgres", dbname) // OK
// db, err = sql.Open("mysql", dbname)    // OK
// db, err = sql.Open("sqlite3", dbname)  // returns error: unknown driver "sqlite3"
