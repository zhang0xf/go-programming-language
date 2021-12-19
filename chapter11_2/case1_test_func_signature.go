package chapter11_2

import "testing"

// 每个测试函数必须导入testing包。测试函数有如下的签名：
func TestName(t *testing.T) {
	// ...
}

// 测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头：
func TestSin(t *testing.T) { /* ... */ }
func TestCos(t *testing.T) { /* ... */ }
func TestLog(t *testing.T) { /* ... */ }

// 其中t参数用于报告测试失败和附加的日志信息。
