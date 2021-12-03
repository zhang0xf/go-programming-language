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

func StringSliceSort() {

	var names StringSlice = make(StringSlice, 0)
	names = append(names, "zhangfei2")
	names = append(names, "zhangfei1")
	names = append(names, "zhangfei3")
	names = append(names, "zhangfei6")
	names = append(names, "zhangfei5")
	names = append(names, "zhangfei4")

	sort.Sort(StringSlice(names))
	StringSlicePrint(names)
}

func StringSlicePrint(s StringSlice) {
	for _, str := range s {
		fmt.Println(str)
	}
}
