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
// 但是对于d，它是c的解引用方式生成的，指向另一个变量，因此是可取地址的。
// 我们可以通过调用reflect.ValueOf(&x).Elem()，来获取任意变量x对应的可取地址的Value。
// 我们可以通过调用reflect.Value的CanAddr方法来判断其是否可以被取地址：

func ReflectValueCanAddr() {
	x := 2                   // value   type    variable?
	a := reflect.ValueOf(2)  // 2       int     no
	b := reflect.ValueOf(x)  // 2       int     no
	c := reflect.ValueOf(&x) // &x      *int    no
	d := c.Elem()            // 2       int     yes (x)

	fmt.Println(a.CanAddr()) // "false"
	fmt.Println(b.CanAddr()) // "false"
	fmt.Println(c.CanAddr()) // "false"
	fmt.Println(d.CanAddr()) // "true"
}

// 每当我们通过指针间接地获取的reflect.Value都是可取地址的，即使开始的是一个不可取地址的Value。
// 在反射机制中，所有关于是否支持取地址的规则都是类似的。
// 例如，slice的索引表达式e[i]将隐式地包含一个指针，它就是可取地址的，即使开始的e表达式不支持也没有关系。
// 以此类推，reflect.ValueOf(e).Index(i)对应的值也是可取地址的，即使原始的reflect.ValueOf(e)不支持也没有关系。

// 要从变量对应的可取地址的reflect.Value来访问变量需要三个步骤。
// 第一步是调用Addr()方法，它返回一个Value，里面保存了指向变量的指针。
// 然后是在Value上调用Interface()方法，也就是返回一个interface{}，里面包含指向变量的指针。
// 最后，如果我们知道变量的类型，我们可以使用类型的断言机制将得到的interface{}类型的接口强制转为普通的类型指针。
// 这样我们就可以通过这个普通指针来更新变量了：
