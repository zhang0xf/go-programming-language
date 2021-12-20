package chapter12_2

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

// 反射是由 reflect 包提供的。它定义了两个重要的类型, Type 和 Value.

// 一个 Type 表示一个Go类型.它是一个接口,有许多方法来区分类型以及检查它们的组成部分,
// 例如一个结构体的成员或一个函数的参数等. 唯一能反映 reflect.Type 实现的是接口的类型描述信息(§7.5), 也正是这个实体标识了接口值的动态类型.

// 函数 reflect.TypeOf 接受任意的 interface{} 类型, 并以reflect.Type形式返回其动态类型:

func ReflectType() {
	t := reflect.TypeOf(3)  // a reflect.Type
	fmt.Println(t.String()) // "int"
	fmt.Println(t)          // "int"
}

// 其中 TypeOf(3) 调用将值 3 传给 interface{} 参数.
// 回到 7.5节 的将一个具体的值转为接口类型会有一个隐式的接口转换操作, 它会创建一个包含两个信息的接口值: 操作数的动态类型(这里是int)和它的动态的值(这里是3).

// 因为 reflect.TypeOf 返回的是一个动态类型的接口值, 它总是返回具体的类型.
// 因此, 下面的代码将打印 "*os.File" 而不是 "io.Writer". 稍后, 我们将看到能够表达接口类型的 reflect.Type.

func ReflectType2() {
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // "*os.File"
}

// 要注意的是 reflect.Type 接口是满足 fmt.Stringer 接口的. 因为打印一个接口的动态类型对于调试和日志是有帮助的,
// fmt.Printf 提供了一个缩写 %T 参数, 内部使用 reflect.TypeOf 来输出:

func ReflectType3() {
	fmt.Printf("%T\n", 3) // "int"
}

// reflect 包中另一个重要的类型是 Value.一个 reflect.Value 可以装载任意类型的值.
// 函数 reflect.ValueOf 接受任意的 interface{} 类型, 并返回一个装载着其动态值的 reflect.Value.
// 和 reflect.TypeOf 类似, reflect.ValueOf 返回的结果也是具体的类型, 但是 reflect.Value 也可以持有一个接口值.

func ReflectValue() {
	v := reflect.ValueOf(3) // a reflect.Value
	fmt.Println(v)          // "3"
	fmt.Printf("%v\n", v)   // "3"
	fmt.Println(v.String()) // NOTE: "<int Value>"
}

// 和 reflect.Type 类似, reflect.Value 也满足 fmt.Stringer 接口, 但是除非 Value 持有的是字符串, 否则 String 方法只返回其类型.
// 而使用 fmt 包的 %v 标志参数会对 reflect.Values 特殊处理.

// 对 Value 调用 Type 方法将返回具体类型所对应的 reflect.Type:

func ReflectValue2() {
	v := reflect.ValueOf(3) // a reflect.Value
	t := v.Type()           // a reflect.Type
	fmt.Println(t.String()) // "int"
}

// reflect.ValueOf 的逆操作是 reflect.Value.Interface 方法. 它返回一个 interface{} 类型，装载着与 reflect.Value 相同的具体值:

func ReflectValue3() {
	v := reflect.ValueOf(3) // a reflect.Value
	x := v.Interface()      // an interface{}
	i := x.(int)            // an int
	fmt.Printf("%d\n", i)   // "3"
}

// reflect.Value 和 interface{} 都能装载任意的值.
// 所不同的是, 一个空的接口隐藏了值内部的表示方式和所有方法, 因此只有我们知道具体的动态类型才能使用类型断言来访问内部的值(就像上面那样), 内部值我们没法访问.
// 相比之下, 一个 Value 则有很多方法来检查其内容, 无论它的具体类型是什么. 让我们再次尝试实现我们的格式化函数 format.Any.(见case2_format.go)
