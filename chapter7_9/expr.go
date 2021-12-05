package chapter7_9

// An Expr is an arithmetic expression.
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64
}

// 为了计算一个包含变量的表达式，我们需要一个environment变量将变量的名字映射成对应的值：
type Env map[Var]float64
