package chapter11_2

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// 对于测试包go test是一个有用的工具，但是稍加努力我们也可以用它来测试可执行程序。
// 如果一个包的名字是 main，那么在构建时会生成一个可执行程序，不过main包可以作为一个包被测试器代码导入。

// 让我们为2.3.2节的echo程序编写一个测试。我们先将程序拆分为两个函数：
// echo函数完成真正的工作，
// main(Echo)函数用于处理命令行输入参数和echo可能返回的错误。

var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "separator")
)

var out io.Writer = os.Stdout // modified during testing

func Echo() {
	flag.Parse()
	if err := echo(!*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func echo(newline bool, sep string, args []string) error {
	fmt.Fprint(out, strings.Join(args, sep))
	if newline {
		fmt.Fprintln(out)
	}
	return nil
}
