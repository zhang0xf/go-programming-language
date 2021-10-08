package chapter3_3

import (
	"log"
	"net/http"
)

func MandelbrotWebServer() {
	HandlerMandelbrotRequest := func(w http.ResponseWriter, r *http.Request) {
		DrawMandelbrotForWeb(w)
	}
	http.HandleFunc("/", HandlerMandelbrotRequest)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
