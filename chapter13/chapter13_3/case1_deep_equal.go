package chapter13_3

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// 来自reflect包的DeepEqual函数可以对两个值进行深度相等判断。
// DeepEqual函数使用内建的==比较操作符对基础类型进行相等判断，对于复合类型则递归该变量的每个基础类型然后做类似的比较判断。
// 因为它可以工作在任意的类型上，甚至对于一些不支持==操作运算符的类型也可以工作，因此在一些测试代码中广泛地使用该函数。
// 比如下面的代码是用DeepEqual函数比较两个字符串数组是否相等。

func TestSplit(t *testing.T) {
	got := strings.Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) { /* ... */
	}
}

// 尽管DeepEqual函数很方便，而且可以支持任意的数据类型，但是它也有不足之处。
// 例如，它将一个nil值的map和非nil值但是空的map视作不相等，同样nil值的slice 和非nil但是空的slice也视作不相等。

func DeepEqualNilAndEmpty() {
	var a, b []string = nil, []string{}
	fmt.Println(reflect.DeepEqual(a, b)) // "false"

	var c, d map[string]int = nil, make(map[string]int)
	fmt.Println(reflect.DeepEqual(c, d)) // "false"
}
