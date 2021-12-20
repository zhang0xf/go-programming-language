package chapter12_3

// 对于目前的实现，如果遇到对象图中含有回环，Display将会陷入死循环，例如下面这个首尾相连的链表：

// a struct that points to itself
type Cycle struct {
	Value int
	Tail  *Cycle
}

func DisplayDeadLoop() {

	var c Cycle
	c = Cycle{42, &c}

	Display("c", c)
}

// Display会永远不停地进行深度递归打印：

// Display c (display.Cycle):
// c.Value = 42
// (*c.Tail).Value = 42
// (*(*c.Tail).Tail).Value = 42
// (*(*(*c.Tail).Tail).Tail).Value = 42
// ...ad infinitum...

// 许多Go语言程序都包含了一些循环的数据。让Display支持这类带环的数据结构需要些技巧，需要额外记录迄今访问的路径；相应会带来成本。
// 通用的解决方案是采用 unsafe 的语言特性，我们将在13.3节看到具体的解决方案。

// 带环的数据结构很少会对fmt.Sprint函数造成问题，因为它很少尝试打印完整的数据结构。
// 例如，当它遇到一个指针的时候，它只是简单地打印指针的数字值。在打印包含自身的slice或map时可能卡住，但是这种情况很罕见，不值得付出为了处理回环所需的开销。
