package chapter8_7

import "fmt"

// 下面这个例子更微妙。
// ch这个channel的buffer大小是1，所以会交替的为空或为满，所以只有一个case可以进行下去，无论i是奇数或者偶数，它都会打印0 2 4 6 8。

func SelectChannel() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		fmt.Printf("--------------------i = %d\n", i)
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
			fmt.Printf("向ch写一个数:%d\n", i)
		}
	}
}

// 如果多个case同时就绪时，select会随机地选择一个执行，这样来保证每一个channel都有平等的被select的机会。
// 增加上面例子的buffer大小会使其输出变得不确定，因为当buffer既不为满也不为空时，select语句的执行情况就像是抛硬币的行为一样是随机的。
