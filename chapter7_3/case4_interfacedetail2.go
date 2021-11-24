package chapter7_3

import (
	"bytes"
	"fmt"
)

// 更多方法的接口类型会告诉我们更多关于它的值持有的信息，并且对实现它的类型要求更加严格。
// 那么关于interface{}类型，它没有任何方法，请讲出哪些具体的类型实现了它？
// 这看上去好像没有用，但实际上interface{}被称为空接口类型是不可或缺的。因为空接口类型对实现它的类型没有要求，所以我们可以将任意一个值赋给空接口类型。

func InterfaceDetail2() {
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
	fmt.Printf("%d\n", any.(int))
}

// 对于创建的一个interface{}值持有一个boolean，float，string，map，pointer，或者任意其它的类型；我们当然不能直接对它持有的值做操作，因为interface{}没有任何方法。
// 我们会在7.10章中学到一种用类型断言来获取interface{}中值的方法。
