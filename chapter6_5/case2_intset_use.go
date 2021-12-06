package chapter6_5

import "fmt"

func UseIntSet() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	fmt.Println(&x)         // "{1 9 42 144}" // IntSet指针类型有定义String()方法
	fmt.Println(x.String()) // "{1 9 42 144}" // IntSet对象会隐式插入&操作符,调用IntSet指针的String()方法
	fmt.Println(x)          // "{[4398046511618 0 65536]}" // IntSet对象没有实现String()方法,所以以默认方式打印.考虑到只是打印而不用修改原始变量,所以可以视情况将string()方法绑定到IntSet对象上.
}
