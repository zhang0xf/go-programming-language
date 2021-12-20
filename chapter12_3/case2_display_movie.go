package chapter12_3

import (
	"fmt"
	"os"
	"reflect"
)

// 现在我们的Display函数总算完工了，让我们看看它的表现吧。下面的Movie类型是在4.5节的电影类型上演变来的：

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

// 让我们声明一个该类型的变量，然后看看Display函数如何显示它：

func DisplayMovie() {

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	Display("strangelove", strangelove)
}

// 我们也可以使用Display函数来显示标准库中类型的内部结构，例如*os.File类型：

func DisplayStdLibStruct() {
	Display("os.Stderr", os.Stderr)
}

// 可以看出，反射能够访问到结构体中未导出的成员。
// 需要当心的是这个例子的输出在不同操作系统上可能是不同的，并且随着标准库的发展也可能导致结果不同。（这也是将这些成员定义为私有成员的原因之一！）
// 我们甚至可以用Display函数来显示reflect.Value 的内部构造（在这里设置为*os.File的类型描述体）。当然不同环境得到的结果可能有差异：

func DisplayReflectValue() {
	Display("rV", reflect.ValueOf(os.Stderr))
}

func DiffTwoDisplay() {
	// 观察下面两个例子的区别：

	var i interface{} = 3

	Display("i", i)
	// Output:
	// Display i (int):
	// i = 3

	fmt.Printf("\n")

	Display("&i", &i)
	// Output:
	// Display &i (*interface {}):
	// (*&i).type = int
	// (*&i).value = 3
}

// 在第一个例子中，Display函数调用reflect.ValueOf(i)，它返回一个Int类型的值。
// 正如我们在12.2节中提到的，reflect.ValueOf总是返回一个具体类型的 Value，因为它是从一个接口值提取的内容。

// 在第二个例子中，Display函数调用的是reflect.ValueOf(&i)，它返回一个指向i的指针，对应Ptr类型。
// 在switch的Ptr分支中，对这个值调用 Elem 方法，返回一个Value来表示变量 i 本身，对应Interface类型。
// 像这样一个间接获得的Value，可能代表任意类型的值，包括接口类型。display函数递归调用自身，这次它分别打印了这个接口的动态类型和值。
