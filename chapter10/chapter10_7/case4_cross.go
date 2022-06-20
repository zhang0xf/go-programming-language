package chapter10_7

import (
	"fmt"
	"runtime"
)

// 下面交叉编译的程序将输出它在编译时的操作系统和CPU类型：

func Cross() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
