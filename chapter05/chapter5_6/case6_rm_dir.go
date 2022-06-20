package chapter5_6

import "os"

// 循环变量与函数值(匿名函数)

func RemoveDir() {
	var rmdirs []func()
	for _, d := range tempDirs() {
		dir := d               // NOTE: necessary! // declares inner dir, initialized to outer dir
		os.MkdirAll(dir, 0755) // creates parent directories too
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}
	// ...do some work…
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
}

func RemoveDir2() {
	var rmdirs []func()
	for _, dir := range tempDirs() {
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir) // NOTE: incorrect!
		})
	}
}

// 问题的原因在于循环变量的作用域。
// 在上面的程序中，for循环语句引入了新的词法块，循环变量dir在这个词法块中被声明。
// 在该循环中生成的所有函数值都共享相同的循环变量。需要注意，函数值中记录的是循环变量的内存地址，而不是循环变量某一时刻的值。
// 以dir为例，后续的迭代会不断更新dir的值，当删除操作执行时，for循环已完成，dir中存储的值等于最后一次迭代的值。这意味着，每次对os.RemoveAll的调用删除的都是相同的目录。

// 这个问题不仅存在基于range的循环，在下面的例子中，对循环变量i的使用也存在同样的问题：
func RemoveDir3() {
	var rmdirs []func()
	dirs := tempDirs()
	for i := 0; i < len(dirs); i++ {
		os.MkdirAll(dirs[i], 0755) // OK
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dirs[i]) // NOTE: incorrect!
		})
	}
}

func tempDirs() []string {
	return nil
}
