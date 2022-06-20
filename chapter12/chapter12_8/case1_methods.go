package chapter12_8

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// 我们的最后一个例子是使用reflect.Type来打印任意值的类型和枚举它的方法：

// reflect.Type和reflect.Value都提供了一个Method方法。
// 每次t.Method(i)调用将一个reflect.Method的实例，对应一个用于描述一个方法的名称和类型的结构体。
// 每次v.Method(i)方法调用都返回一个reflect.Value以表示对应的值（§6.4），也就是一个方法是帮到它的接收者的。
// 使用reflect.Value.Call方法（我们这里没有演示），将可以调用一个Func类型的Value，但是这个例子中只用到了它的类型。

// Print prints the method set of the value x.
func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		name := t.Method(i).Name
		methTypeString := strings.TrimPrefix(methType.String(), "func")
		fmt.Printf("func (%s) %s%s\n", t, name, methTypeString)
	}
}

func PrintMethods() {
	Print(time.Hour)
	Print(new(strings.Replacer))
}
