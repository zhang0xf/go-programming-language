// Usage : ./exercise(启动Web服务,使用chrome访问localhost:8000/,返回surface图像)

package chapter3_2

import (
	"log"
	"net/http"
)

func WebServerWithSVGSurface() {
	HandlerRequestAndSVG := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		SVGSurfaceWeb(w)
	}
	http.HandleFunc("/", HandlerRequestAndSVG)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
