package chapter1_7

import (
	"exercise/chapter01/chapter1_4"
	"log"
	"net/http"
)

func LissajousWebServer() {
	HandlerLissajousRequest := func(w http.ResponseWriter, r *http.Request) {
		chapter1_4.Lissajous(w)
	}
	http.HandleFunc("/", HandlerLissajousRequest)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
