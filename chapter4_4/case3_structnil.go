package chapter4_4

// 结构体零值
// 结构体类型的零值是每个成员都是零值。
// 对于bytes.Buffer类型，结构体初始值就是一个随时可用的空缓存，还有在第9章将会讲到的sync.Mutex的零值也是有效的未锁定状态。

// 空结构体用法之一:
// 有些Go语言程序员用map来模拟set数据结构时，用空结构体来代替map中布尔类型的value,但是因为节约的空间有限，而且语法比较复杂，所以我们通常会避免这样的用法。

func StructNil() {
	seen := make(map[string]struct{}) // set of strings
	// ...
	s := ""
	if _, ok := seen[s]; !ok {
		seen[s] = struct{}{}
		// ...first time seeing s...
	}
}
