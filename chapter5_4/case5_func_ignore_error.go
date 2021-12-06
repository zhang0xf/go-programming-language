package chapter5_4

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 错误处理策略
// 5.直接忽略掉错误。

func FunctionIgnoreError() error {
	dir, err := ioutil.TempDir("", "scratch")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %v", err)
	}
	// ...use temp dir…
	os.RemoveAll(dir) // ignore errors; $TMPDIR is cleaned periodically
	return nil
}

// 尽管os.RemoveAll会失败，但上面的例子并没有做错误处理。这是因为操作系统会定期的清理临时目录。

// 函数错误处理总结:
// 检查某个子函数是否失败后，我们通常将处理失败的逻辑代码放在处理成功的代码之前。
// 如果某个错误会导致函数返回，那么成功时的逻辑代码不应放在else语句块中，而应直接放在函数体中。
