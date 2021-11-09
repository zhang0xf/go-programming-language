package chapter4_4

func StructExport() {
	// 企图隐式使用未导出成员的行为是不允许的
	// var _ = chapter4_4_1.T{a: 1, b: 2} // compile error: can't reference a, b
	// var _ = chapter4_4_1.T{1, 2}       // compile error: can't reference a, b
}
