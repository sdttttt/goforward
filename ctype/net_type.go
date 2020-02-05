package ctype

type ProtocolType int

const (
	HTTP  ProtocolType = 1
	HTTPS ProtocolType = 2
	TCP   ProtocolType = 3

	HTTP_PORT   = "80"
	HTTP_PREFIX = "http://"
)
