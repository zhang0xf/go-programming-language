package chapter11_4

import (
	"exercise/chapter11_2"
	"testing"
)

// usage : go test -bench=. ./chapter11_4/
//         go test -bench=. ./chapter11_4/ -benchmem

// 基准测试

// 基准测试是测量一个程序在固定工作负载下的性能。
// 在Go语言中，基准测试函数和普通测试函数写法类似，但是以Benchmark为前缀名，并且带有一个*testing.B类型的参数；*testing.B参数除了提供和*testing.T类似的方法，还有额外一些和性能测量相关的方法。

// 下面是IsPalindrome函数的基准测试，其中循环将执行N次。

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chapter11_2.IsPalindrome("A man, a plan, a canal: Panama")
	}
}

// 我们用下面的命令运行基准测试。和普通测试不同的是，默认情况下不运行任何基准测试。
// 我们需要通过-bench命令行标志参数手工指定要运行的基准测试函数。该参数是一个正则表达式，用于匹配要执行的基准测试函数的名字，默认值是空的。
// 其中“.”模式将可以匹配所有基准测试函数，但因为这里只有一个基准测试函数，因此和-bench=IsPalindrome参数是等价的效果。
// $ cd $GOPATH/src/gopl.io/ch11/word2
// $ go test -bench=.
// PASS
// BenchmarkIsPalindrome-8 1000000                1035 ns/op
// ok      gopl.io/ch11/word2      2.179s

// 结果中基准测试名的数字后缀部分，这里是8，表示运行时对应的GOMAXPROCS的值，这对于一些与并发相关的基准测试是重要的信息。
// 报告显示每次调用IsPalindrome函数花费1.035微秒，是执行1,000,000次的平均时间。
// 因为基准测试驱动器开始时并不知道每个基准测试函数运行所花的时间，它会尝试在真正运行基准测试前先尝试用较小的N运行测试来估算基准测试函数所需要的时间，然后推断一个较大的时间保证稳定的测量结果。
// 循环在基准测试函数内实现，而不是放在基准测试框架内实现，这样可以让每个基准测试函数有机会在循环启动前执行初始化代码，这样并不会显著影响每次迭代的平均运行时间。
// 如果还是担心初始化代码部分对测量时间带来干扰，那么可以通过testing.B参数提供的方法来临时关闭或重置计时器，不过这些一般很少会用到。

// 现在我们有了一个基准测试和普通测试，我们可以很容易测试改进程序运行速度的想法。
// 也许最明显的优化是在IsPalindrome函数中第二个循环的停止检查，这样可以避免每个比较都做两次：

// n := len(letters) / 2
// for i := 0; i < n; i++ {
// 	if letters[i] != letters[len(letters)-1-i] {
// 		return false
// 	}
// }
// return true

// 不过很多情况下，一个显而易见的优化未必能带来预期的效果。这个改进在基准测试中只带来了4%的性能提升。

// 另一个改进想法是在开始为每个字符预先分配一个足够大的数组，这样就可以避免在append调用时可能会导致内存的多次重新分配。
// 声明一个letters数组变量，并指定合适的大小，像下面这样，

// letters := make([]rune, 0, len(s))
// for _, r := range s {
//     if unicode.IsLetter(r) {
//         letters = append(letters, unicode.ToLower(r))
//     }
// }

// 这个改进提升性能约35%，报告结果是基于2,000,000次迭代的平均运行时间统计。
// 如这个例子所示，快的程序往往是伴随着较少的内存分配。
// -benchmem命令行标志参数将在报告中包含内存的分配数据统计。我们可以比较优化前后内存的分配情况：

// $ go test -bench=. -benchmem
// PASS
// BenchmarkIsPalindrome    1000000   1026 ns/op    304 B/op  4 allocs/op

// 这是优化之后的结果：

// $ go test -bench=. -benchmem
// PASS
// BenchmarkIsPalindrome    2000000    807 ns/op    128 B/op  1 allocs/op

// 用一次内存分配代替多次的内存分配节省了75%的分配调用次数和减少近一半的内存需求。

// 这个基准测试告诉了我们某个具体操作所需的绝对时间，但我们往往想知道的是两个不同的操作的时间对比。
// 比较型的基准测试就是普通程序代码。它们通常是单参数的函数，由几个不同数量级的基准测试函数调用，就像这样：

func benchmark(b *testing.B, size int) { /* ... */ }
func Benchmark10(b *testing.B)         { benchmark(b, 10) }
func Benchmark100(b *testing.B)        { benchmark(b, 100) }
func Benchmark1000(b *testing.B)       { benchmark(b, 1000) }

// 通过函数参数来指定输入的大小，但是参数变量对于每个具体的基准测试都是固定的。要避免直接修改b.N来控制输入的大小。

// 比较型的基准测试反映出的模式在程序设计阶段是很有帮助的，但是即使程序完工了也应当保留基准测试代码。
// 因为随着项目的发展，或者是输入的增加，或者是部署到新的操作系统或不同的处理器，我们可以再次用基准测试来帮助我们改进设计。
