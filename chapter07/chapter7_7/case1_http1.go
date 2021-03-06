package chapter7_7

import (
	"fmt"
	"log"
	"net/http"
)

// usage : ./exercise
// chrome input : localhost:8000
// 或使用1.5节fetch

// 目前为止，这个服务器不考虑URL，只能为每个请求列出它全部的库存清单。

func Http1() {
	db := database1{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database1 map[string]dollars

func (db database1) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
