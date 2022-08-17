package chapter11_3

import (
	"exercise/chapter07/chapter7_9"
	"fmt"
	"math"
	"testing"
)

// usage : go test -run=Coverage -coverprofile=./data/c.out ./chapter11_3/
//         go tool cover -html=./data/c.out

// 下面的代码是一个表格驱动的测试，用于测试第七章的表达式求值程序：

func TestCoverage(t *testing.T) {
	var tests = []struct {
		input string
		env   chapter7_9.Env
		want  string // expected error from Parse/Check or result from Eval
	}{
		{"x % 2", nil, "unexpected '%'"},
		{"!true", nil, "unexpected '!'"},
		{"log(10)", nil, `unknown function "log"`},
		{"sqrt(1, 2)", nil, "call to sqrt has 2 args, want 1"},
		{"sqrt(A / pi)", chapter7_9.Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", chapter7_9.Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", chapter7_9.Env{"F": -40}, "-40"},
	}

	for _, test := range tests {
		expr, err := chapter7_9.Parse(test.input)
		if err == nil {
			err = expr.Check(map[chapter7_9.Var]bool{})
		}
		if err != nil {
			if err.Error() != test.want {
				t.Errorf("%s: got %q, want %q", test.input, err, test.want)
			}
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		if got != test.want {
			t.Errorf("%s: %v => %s, want %s",
				test.input, test.env, got, test.want)
		}
	}
}

// 下面这个命令可以显示测试覆盖率工具的使用用法：
// $ go tool cover
// Usage of 'go tool cover':
// Given a coverage profile produced by 'go test':
//     go test -coverprofile=c.out

// Open a web browser displaying annotated source code:
//     go tool cover -html=c.out
// ...

// go tool命令运行Go工具链的底层可执行程序。
// 这些底层可执行程序放在$GOROOT/pkg/tool/${GOOS}_${GOARCH}目录。
// 因为有go build命令的原因，我们很少直接调用这些底层工具。

// 现在我们可以用-coverprofile标志参数重新运行测试：
// $ go test -run=Coverage -coverprofile=c.out gopl.io/ch7/eval
// ok      gopl.io/ch7/eval         0.032s      coverage: 68.5% of statements

// 这个标志参数通过在测试代码中插入生成钩子来统计覆盖率数据。
// 也就是说，在运行每个测试前，它将待测代码拷贝一份并做修改，在每个词法块都会设置一个布尔标志变量。
// 当被修改后的被测试代码运行退出时，将统计日志数据写入c.out文件，并打印一部分执行的语句的一个总结。

// 如果使用了-covermode=count标志参数，那么将在每个代码块插入一个计数器而不是布尔标志量。
// 在统计结果中记录了每个块的执行次数，这可以用于衡量哪些是被频繁执行的热点代码。

// 为了收集数据，我们运行了测试覆盖率工具，打印了测试日志，生成一个HTML报告，然后在浏览器中打开（图11.3）。
// $ go tool cover -html=c.out

// 实现100%的测试覆盖率听起来很美，但是在具体实践中通常是不可行的，也不是值得推荐的做法。
// 因为那只能说明代码被执行过而已，并不意味着代码就是没有BUG的；因为对于逻辑复杂的语句需要针对不同的输入执行多次。
// 测试从本质上来说是一个比较务实的工作，编写测试代码和编写应用代码的成本对比是需要考虑的。
// 测试覆盖率工具可以帮助我们快速识别测试薄弱的地方，但是设计好的测试用例和编写应用代码一样需要严密的思考。
