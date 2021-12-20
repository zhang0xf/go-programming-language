package chapter12_1

import "strconv"

// Go语言提供了一种机制，能够在运行时更新变量和检查它们的值、调用它们的方法和它们支持的内在操作，而不需要在编译时就知道这些变量的具体类型。这种机制被称为反射。
// 反射也可以让我们将类型本身作为第一类的值类型处理。

// 在本章，我们将探讨Go语言的反射特性，看看它可以给语言增加哪些表达力，以及在两个至关重要的API是如何使用反射机制的：
// 一个是fmt包提供的字符串格式化功能，
// 另一个是类似encoding/json和encoding/xml提供的针对特定协议的编解码功能。
// 然后，反射是一个复杂的内省技术，不应该随意使用，因此，尽管上面这些包内部都是用反射技术实现的，但是它们自己的API都没有公开反射相关的接口。

// 为何需要反射?

// 有时候我们需要编写一个函数能够处理一类并不满足普通公共接口的类型的值，也可能是因为它们并没有确定的表示方式，或者是在我们设计该函数的时候这些类型可能还不存在。
// 一个大家熟悉的例子是fmt.Fprintf函数提供的字符串格式化处理逻辑，它可以用来对任意类型的值格式化并打印，甚至支持用户自定义的类型。让我们也来尝试实现一个类似功能的函数。
// 为了简单起见，我们的函数只接收一个参数，然后返回和fmt.Sprint类似的格式化后的字符串。我们实现的函数名也叫Sprint。

// 我们首先用switch类型分支来测试输入参数是否实现了String方法，如果是的话就调用该方法。
// 然后继续增加类型测试分支，检查这个值的动态类型是否是string、int、bool等基础类型，并在每种情况下执行相应的格式化操作。

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	// ...similar cases for int16, uint32, and so on...
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array, chan, func, map, pointer, slice, struct
		return "???"
	}
}

// 但是我们如何处理其它类似[]float64、map[string][]string等类型呢？
// 我们当然可以添加更多的测试分支，但是这些组合类型的数目基本是无穷的。
// 还有如何处理类似url.Values这样的具名类型呢？
// 即使类型分支可以识别出底层的基础类型是map[string][]string，但是它并不匹配url.Values类型，因为它们是两种不同的类型，
// 而且switch类型分支也不可能包含每个类似url.Values的类型，这会导致对这些库的依赖。
// 没有办法来检查未知类型的表示方式，我们被卡住了。这就是我们为何需要反射的原因。
