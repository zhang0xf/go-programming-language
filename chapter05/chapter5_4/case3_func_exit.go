package chapter5_4

import (
	"fmt"
	"log"
	"os"
)

// 错误处理策略:
// 3.如果错误发生后，程序无法继续运行，我们就可以采用第三种策略：输出错误信息并结束程序。

func FunctionExit() {
	// 设置log的前缀信息,屏蔽时间信息(for the standard logger)
	log.SetPrefix("wait: ")
	log.SetFlags(0)

	var url string = ""
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		// 需要注意的是，这种策略只应在main中执行。对库函数而言，应仅向上传播错误，除非该错误意味着程序内部包含不一致性，即遇到了bug，才能在库函数中结束程序。
		os.Exit(1)
	}

	// 调用log.Fatalf可以更简洁的代码达到与上文相同的效果。
	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
}
