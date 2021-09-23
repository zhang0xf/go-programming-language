package chapter1_7

import (
	"exercise/chapter1_4"
	"log"
	"net/http"
)

func WebServerWithLissajous() {
	HandlerRequestAndLissajous := func(w http.ResponseWriter, r *http.Request) {
		chapter1_4.Lissajous(w)
	}
	http.HandleFunc("/", HandlerRequestAndLissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
