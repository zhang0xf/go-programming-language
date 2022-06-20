package chapter7_9

// An Expr is an arithmetic expression.
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64

	// Check reports errors in this Expr and adds its Vars to the set.
	Check(vars map[Var]bool) error
}

// 为了计算一个包含变量的表达式，我们需要一个environment变量将变量的名字映射成对应的值：
type Env map[Var]float64

// 延申：
// 让我们往Expr接口中增加另一个方法。Check方法对一个表达式语义树检查出静态错误。
// 甚至在解释型语言中，为了静态错误检查语法是非常常见的；静态错误就是不用运行程序就可以检测出来的错误。
// 通过将静态检查和动态的部分分开，我们可以快速的检查错误并且对于多次检查只执行一次而不是每次表达式计算的时候都进行检查。
