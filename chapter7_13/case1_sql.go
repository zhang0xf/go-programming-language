package chapter7_13

import (
	"database/sql"
	"fmt"
)

// 类型分支

// Exec方法使用SQL字面量替换在查询字符串中的每个'?'；SQL字面量表示相应参数的值，它有可能是一个布尔值，一个数字，一个字符串，或者nil空值。
// 用这种方式构造查询可以帮助避免SQL注入攻击；这种攻击就是对手可以通过利用输入内容中不正确的引号来控制查询语句。
func listTracks(db sql.DB, artist string, minYear, maxYear int) {
	result, err := db.Exec(
		"SELECT * FROM tracks WHERE artist = ? AND ? <= year AND year <= ?",
		artist, minYear, maxYear)
	// ...

	fmt.Printf("%T\n", result)
	fmt.Printf("%T\n", err)
}

// 在Exec函数内部，我们可能会找到像下面这样的一个函数，它会将每一个参数值转换成它的SQL字面量符号。
func sqlQuote(x interface{}) string {
	if x == nil {
		return "NULL"
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d", x)
	} else if b, ok := x.(bool); ok {
		if b {
			return "TRUE"
		}
		return "FALSE"
	} else if s, ok := x.(string); ok {
		// return sqlQuoteString(s) // (not shown)
		return s
	} else {
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

// switch语句可以简化if-else链，如果这个if-else链对一连串值做相等测试。
func sqlQuote2(x interface{}) string {
	switch x.(type) {
	case nil: // ...
	case int, uint: // ...
	case bool: // ...
	case string: // ...
	default: // ...
	}
	return ""
}

// 使用类型分支的扩展形式来重写sqlQuote函数会让这个函数更加的清晰：
// 当多个case需要相同的操作时，比如int和uint的情况，类型分支可以很容易的合并这些情况。
func sqlQuote3(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x) // x has type interface{} here.
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		// return sqlQuoteString(x) // (not shown)
		return x
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}
