package main

import (
	"net/http"

	"github.com/sdttttt/goforward"
)

func main() {
	http.ListenAndServe("0.0.0.0:10086", &goforward.ProxyHTTPServer{})
}
