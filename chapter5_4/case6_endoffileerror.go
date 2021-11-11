package chapter5_4

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 文件结尾错误(EOF)

// ...这会导致调用者必须分别处理由文件结束引起的各种错误。基于这样的原因，io包保证任何由文件结束引起的读取失败都返回同一个错误——io.EOF
// 4.3的chartcount程序展示了更加复杂的代码

func EndOfFileError() error {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break // finished reading
		}
		if err != nil {
			return fmt.Errorf("read failed:%v", err)
		}
		// ...use r…
		fmt.Printf("%T", r)
	}
	return nil
}
