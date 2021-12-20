package chapter12_2

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// 我们使用 reflect.Value 的 Kind 方法来替代之前的类型 switch.
// 虽然还是有无穷多的类型, 但是它们的kinds类型却是有限的:
// Bool, String 和 所有数字类型的基础类型;
// Array 和 Struct 对应的聚合类型;
// Chan, Func, Ptr, Slice, 和 Map 对应的引用类型;
// interface 类型;
// 还有表示空值的 Invalid 类型.(空的 reflect.Value 的 kind 即为 Invalid.)

// Any formats any value as a string.
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

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

// 到目前为止, 我们的函数将每个值视作一个不可分割没有内部结构的物品, 因此它叫 formatAtom.
// 对于聚合类型(结构体和数组)和接口，只是打印值的类型,
// 对于引用类型(channels, functions, pointers, slices, 和 maps), 打印类型和十六进制的引用地址.
// 虽然还不够理想, 但是依然是一个重大的进步, 并且 Kind 只关心底层表示, format.Any 也支持具名类型. 例如:

func FormatAny() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(Any(x))                  // "1"
	fmt.Println(Any(d))                  // "1"
	fmt.Println(Any([]int64{x}))         // "[]int64 0x8202b87b0"
	fmt.Println(Any([]time.Duration{d})) // "[]time.Duration 0x8202b87e0"
}
