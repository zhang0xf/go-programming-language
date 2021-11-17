package chapter5_6

import (
	"fmt"
	"sort"
)

// prereqs记录了每个课程的前置课程
// 因为prereqs映射的是切片而不是更复杂的map，所以数据的遍历次序是固定的
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// 这类问题被称作拓扑排序。
// 图中的顶点表示课程，边表示课程间的依赖关系。显然，图中应该无环
// 下面的代码用深度优先搜索了整张图，获得了符合要求的课程序列。
func TopoSort() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// 当匿名函数需要被递归调用时，我们必须首先声明一个变量
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
