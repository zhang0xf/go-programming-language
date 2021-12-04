package chapter7_7

import (
	"fmt"
	"log"
	"net/http"
)

// 为了避免方法ServeHTTP重复，新建了database2用以区分database1
// usage : ./exercise
// chrome input : http://localhost:8000/list
//                http://localhost:8000/price?item=socks
//                http://localhost:8000/price?item=shoes
//                http://localhost:8000/price?item=hat
//                http://localhost:8000/help
// 或使用1.5节fetch

// 现在handler基于URL的路径部分（req.URL.Path）来决定执行什么逻辑。

type database2 map[string]dollars

func Http2() {
	db := database2{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

func (db database2) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			// http.ResponseWriter是另一个接口。它在io.Writer上增加了发送HTTP相应头的方法。
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			// 等效地，我们可以使用实用的http.Error函数：
			// msg := fmt.Sprintf("no such page: %s\n", req.URL)
			// http.Error(w, msg, http.StatusNotFound) // 404
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}
