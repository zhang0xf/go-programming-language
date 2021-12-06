package chapter5_4

import (
	"fmt"
	"log"
	"os"
)

// 错误处理策略:
// 4.有时，我们只需要输出错误信息就足够了，不需要中断程序的运行。

func FunctionLog() {
	// log包提供函数
	// log包中的所有函数会为没有换行符的字符串增加换行符。
	if err := Ping(); err != nil {
		log.Printf("ping failed: %v; networking disabled", err)
	}

	// 标准错误流输出错误信息
	if err := Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", err)
	}
}

func Ping() error {
	return nil
}
