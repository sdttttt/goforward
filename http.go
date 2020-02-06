package goforward

import (
	"io"
	"net"
	"net/http"
	"strings"
)

func (*ProxyHTTPServer) httpHandler(res http.ResponseWriter, req *http.Request) {
	proxy := http.DefaultTransport

	outReq := new(http.Request)
	*outReq = *req

	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}

	pres, err := proxy.RoundTrip(outReq)
	if err != nil {
		res.WriteHeader(http.StatusBadGateway)
		return
	}

	for k, v := range pres.Header {
		for _, vv := range v {
			res.Header().Add(k, vv)
		}
	}

	res.WriteHeader(pres.StatusCode)
	io.Copy(res, pres.Body)
	pres.Body.Close()
}
