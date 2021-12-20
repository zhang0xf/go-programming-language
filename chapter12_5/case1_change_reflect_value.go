package chapter12_5

import (
	"fmt"
	"reflect"
)

// 到目前为止，反射还只是程序中变量的另一种读取方式。然而，在本节中我们将重点讨论如何通过反射机制来修改变量。
// 回想一下，Go语言中类似x、x.f[1]和*p形式的表达式都可以表示变量，但是其它如x + 1和f(2)则不是变量。
// 一个变量就是一个可寻址的内存空间，里面存储了一个值，并且存储的值可以通过内存地址来更新。
// 对于reflect.Values也有类似的区别。有一些reflect.Values是可取地址的；其它一些则不可以。
// 考虑以下的声明语句：

func ReflectVariable() {
	x := 2                   // value   type    variable?
	a := reflect.ValueOf(2)  // 2       int     no
	b := reflect.ValueOf(x)  // 2       int     no
	c := reflect.ValueOf(&x) // &x      *int    no
	d := c.Elem()            // 2       int     yes (x)

	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", d)
}

// 其中a对应的变量不可取地址。因为a中的值仅仅是整数2的拷贝副本。
// b中的值也同样不可取地址。
// c中的值还是不可取地址，它只是一个指针&x的拷贝。
// 实际上，所有通过reflect.ValueOf(x)返回的reflect.Value都是不可取地址的。
