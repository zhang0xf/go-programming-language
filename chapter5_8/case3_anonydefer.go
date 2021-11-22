package chapter5_8

import (
	"fmt"
	"os"
)

// defer语句中的函数会在return语句更新返回值变量后再执行，又因为在函数中定义的匿名函数可以访问该函数包括返回值变量在内的所有变量，所以，对匿名函数采用defer机制，可以使其观察函数的返回值。

func Double() {
	_ = double2(4)
	// Output:
	// "double(4) = 8"
}

func double1(x int) int {
	return x + x
}

func double2(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

// 但对于有许多return语句的函数而言，这个技巧很有用。被延迟执行的匿名函数甚至可以修改函数返回给调用者的返回值：

func triple(x int) (result int) {
	defer func() { result += x }()
	return double1(x)
}

func Triple() {
	fmt.Println(triple(4)) // "12"
}

// 在循环体中的defer语句需要特别注意，因为只有在函数执行完毕后，这些被延迟的函数才会执行。
// 下面的代码会导致系统的文件描述符耗尽，因为在所有文件都被处理之前，没有文件会被关闭。

func LoopDefer() error {
	var filenames []string
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close() // NOTE: risky; could run out of file descriptors
		// ...process f…
	}
	return nil
}

// 一种解决方法是将循环体中的defer语句移至另外一个函数。在每次循环时，调用这个函数。

func LoopDefer2() error {
	var filenames []string
	for _, filename := range filenames {
		if err := doFile(filename); err != nil {
			return err
		}
	}
	return nil
}

func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ...process f…
	return nil
}
