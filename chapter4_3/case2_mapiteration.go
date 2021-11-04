package chapter4_3

import (
	"fmt"
	"sort"
)

func MapIteration() {
	ages := make(map[string]int)

	// map的迭代顺序每次都不一样
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// 必须显示地对key进行排序
	var names []string
	// names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

}

func MapNil() {
	// map上的大部分操作，包括查找、删除、len和range循环都可以安全工作在nil值的map上，它们的行为和一个空的map类似。
	var ages map[string]int
	fmt.Println(ages == nil)    // "true"
	fmt.Println(len(ages) == 0) // "true"

	// 但是向一个nil值的map存入元素将导致一个panic异常,在向map存数据前必须先创建map。
	ages["carol"] = 21 // panic: assignment to entry in nil map
}

func MapOk() {
	ages := make(map[string]int)

	// 如果元素类型是一个数字，你可能需要区分一个已经存在的0，和不存在而返回零值的0
	age, ok := ages["bob"]
	if !ok {
		/* "bob" is not a key in this map; age == 0. */
		fmt.Printf("%d", age)
	}

	if age, ok := ages["bob"]; !ok {
		/* ... */
		fmt.Printf("%d", age)
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
