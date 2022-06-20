package chapter7_6

import (
	"fmt"
	"sort"
)

// 一个内置的排序算法需要知道三个东西：序列的长度，表示两个元素比较的结果，一种交换两个元素的方式；这就是sort.Interface的三个方法
// 为了对序列进行排序，我们需要定义一个实现了这三个方法的类型，然后对这个类型的一个实例应用sort.Sort函数。

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// 对字符串切片的排序是很常用的需要，所以sort包提供了StringSlice类型，也提供了Strings函数能让上面这些调用简化成sort.Strings(names)。

func StringSliceSort() {
	names := []string{
		"zhangfei2",
		"zhangfei1",
		"zhangfei3",
		"zhangfei6",
		"zhangfei5",
		"zhangfei4",
	}

	sort.Sort(StringSlice(names))
	// sort.Strings(names)
	StringSlicePrint(names)
}

func StringSlicePrint(s StringSlice) {
	for _, str := range s {
		fmt.Println(str)
	}
}
