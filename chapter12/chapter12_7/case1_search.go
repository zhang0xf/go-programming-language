package chapter12_7

import (
	"fmt"
	"net/http"
)

// 获取结构体字段标签

// 在4.5节我们使用构体成员标签用于设置对应JSON对应的名字。其中json成员标签让我们可以选择成员的名字和抑制零值成员的输出。
// 在本节，我们将看到如何通过反射机制类获取成员标签。

// 对于一个web服务，大部分HTTP处理函数要做的第一件事情就是展开请求中的参数到本地变量中。
// 我们定义了一个工具函数，叫params.Unpack，通过使用结构体成员标签机制来让HTTP处理函数解析请求参数更方便。

// 首先，我们看看如何使用它。下面的search函数是一个HTTP请求处理函数。
// 它定义了一个匿名结构体类型的变量，用结构体的每个成员表示HTTP请求的参数。
// 其中结构体成员标签指明了对于请求参数的名字，为了减少URL的长度这些参数名通常都是神秘的缩略词。
// Unpack将请求参数填充到合适的结构体成员中，这样我们可以方便地通过合适的类型类来访问这些参数。

// search implements the /search URL endpoint.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10 // set default
	if err := Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

// $ go build gopl.io/ch12/search
// $ ./search &
// $ ./fetch 'http://localhost:12345/search'
// Search: {Labels:[] MaxResults:10 Exact:false}
// $ ./fetch 'http://localhost:12345/search?l=golang&l=programming'
// Search: {Labels:[golang programming] MaxResults:10 Exact:false}
// $ ./fetch 'http://localhost:12345/search?l=golang&l=programming&max=100'
// Search: {Labels:[golang programming] MaxResults:100 Exact:false}
// $ ./fetch 'http://localhost:12345/search?x=true&l=golang&l=programming'
// Search: {Labels:[golang programming] MaxResults:10 Exact:true}
// $ ./fetch 'http://localhost:12345/search?q=hello&x=123'
// x: strconv.ParseBool: parsing "123": invalid syntax
// $ ./fetch 'http://localhost:12345/search?q=hello&max=lots'
// max: strconv.ParseInt: parsing "lots": invalid syntax
