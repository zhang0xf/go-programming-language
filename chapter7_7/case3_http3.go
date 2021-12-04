package chapter7_7

import (
	"fmt"
	"log"
	"net/http"
)

// 一个ServeMux将一批http.Handler聚集到一个单一的http.Handler中。
// 再一次，我们可以看到满足同一接口的不同类型是可替换的：web服务器将请求指派给任意的http.Handler 而不需要考虑它后面的具体类型。
// 注释：HandlerFunc是func类型，且实现了ServeHTTP方法，只不过ServeHTTP的接收器是func（比如：list和price）<F12>
// 对于更复杂的应用，一些ServeMux可以通过组合来处理更加错综复杂的路由需求。??

func Http3() {
	db := database3{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	// db.list是一个方法值 (§6.4)
	// 所以db.list是一个实现了handler类似行为的函数，但是因为它没有方法，所以它不满足http.Handler接口并且不能直接传给mux.Handle。
	// 语句http.HandlerFunc(db.list)是一个转换而非一个函数调用，因为http.HandlerFunc是一个类型。
	// HandlerFunc显示了在Go语言接口机制中一些不同寻常的特点。这是一个实现了接口http.Handler的方法的函数类型。
	// ServeHTTP方法的行为是调用了它的函数本身。因此HandlerFunc是一个让函数值满足一个接口的适配器，这里函数和这个接口仅有的方法有相同的函数签名。
	// 实际上，这个技巧让一个单一的类型例如database以多种方式满足http.Handler接口(ServeHTTP方法)：一种通过它的list方法，一种通过它的price方法等等。
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))

	// 因为用法非常普遍，所以它实现了一个更简单的方法：
	// mux.HandleFunc("/list", db.list)
	// mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database3 map[string]dollars

func (db database3) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database3) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

// 为什么有锁？（mux）并行场景？
// 最后，一个重要的提示：就像我们在1.7节中提到的，web服务器在一个新的协程中调用每一个handler，
// 所以当handler获取其它协程或者这个handler本身的其它请求也可以访问到变量时，一定要使用预防措施，比如锁机制。
