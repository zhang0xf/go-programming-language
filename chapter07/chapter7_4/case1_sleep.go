package chapter7_4

import (
	"flag"
	"fmt"
	"time"
)

// usage : ./exercise
//         ./exercise -period 50ms
//         ./exercise -period 2m30s

var period = flag.Duration("period", 1*time.Second, "sleep period")

func Sleep() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

// 默认情况下，休眠周期是一秒，但是可以通过 -period 这个命令行标记来控制。

// 因为时间周期标记值非常的有用，所以这个特性被构建到了flag包中；
// 但是我们为我们自己的数据类型定义新的标记符号是简单容易的。我们只需要定义一个实现flag.Value接口的类型
