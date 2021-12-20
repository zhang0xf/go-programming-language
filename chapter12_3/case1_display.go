package chapter12_3

import (
	"exercise/chapter7_9"
	"fmt"
	"reflect"
	"strconv"
)

// 接下来，让我们看看如何改善聚合数据类型的显示。
// 我们并不想完全克隆一个fmt.Sprint函数，我们只是构建一个用于调试用的Display函数：给定任意一个复杂类型 x，打印这个值对应的完整结构，同时标记每个元素的发现路径。
// 让我们从一个例子开始。

func DispLaySyntaxTree() {
	// e, _ := eval.Parse("sqrt(A / pi)")
	e, _ := chapter7_9.Parse("sqrt(A / pi)")
	// // 传入Display函数的参数是在7.9节一个表达式求值函数返回的语法树。
	Display("e", e)
}

// 你应该尽量避免在一个包的API中暴露涉及反射的接口。
// 我们将定义一个未导出的display函数用于递归处理工作，导出的是Display函数，它只是display函数简单的包装以接受interface{}类型的参数：

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

// 在display函数中，我们使用了前面定义的打印基础类型——基本类型、函数和chan等——元素值的formatAtom函数，
// 但是我们会使用reflect.Value的方法来递归显示复杂类型的每一个成员。
// 在递归下降过程中，path字符串，从最开始传入的起始值（这里是“e”），将逐步增长来表示是如何达到当前值（例如“e.args[0].value”）的。
func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

// 让我们针对不同类型分别讨论。

// Slice和数组：
// 两种的处理逻辑是一样的。
// Len方法返回slice或数组值中的元素个数，Index(i)获得索引i对应的元素，返回的也是一个reflect.Value；
// 如果索引i超出范围的话将导致panic异常，这与数组或slice类型内建的len(a)和a[i]操作类似。
// display针对序列中的每个元素递归调用自身处理，我们通过在递归处理时向path附加“[i]”来表示访问路径。
// 虽然reflect.Value类型带有很多方法，但是只有少数的方法能对任意值都安全调用。
// 例如，Index方法只能对Slice、数组或字符串类型的值调用，如果对其它类型调用则会导致panic异常。

// 结构体：
// NumField方法报告结构体中成员的数量，Field(i)以reflect.Value类型返回第i个成员的值。成员列表也包括通过匿名字段提升上来的成员。
// 为了在path添加“.f”来表示成员路径，我们必须获得结构体对应的reflect.Type类型信息，然后访问结构体第i个成员的名字。

// Maps:
// MapKeys方法返回一个reflect.Value类型的slice，每一个元素对应map的一个key。和往常一样，遍历map时顺序是随机的。
// MapIndex(key)返回map中key对应的value。
// 我们向path添加“[key]”来表示访问路径。（我们这里有一个未完成的工作。其实map的key的类型并不局限于formatAtom能完美处理的类型；数组、结构体和接口都可以作为map的key。）

// 指针：
// Elem方法返回指针指向的变量，依然是reflect.Value类型。
// 即使指针是nil，这个操作也是安全的，在这种情况下指针是Invalid类型，但是我们可以用IsNil方法来显式地测试一个空指针，这样我们可以打印更合适的信息。
// 我们在path前面添加“*”，并用括弧包含以避免歧义。

// 接口：
// 再一次，我们使用IsNil方法来测试接口是否是nil，如果不是，我们可以调用v.Elem()来获取接口对应的动态值，并且打印对应的类型和值。

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
