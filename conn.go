package goforward

import "net/http"

type ProxyHTTPServer struct{}

func (proxy *ProxyHTTPServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == HttpsMethod {
		println("HTTPS")
		proxy.httpsHandler(res, req)
	} else {
		proxy.httpHandler(res, req)
	}
}
