package goforward

const (
	// SPACE => " "
	SPACE = " "

	//CR => a NewLine character
	CR = "\r\n"

	//TCP => network connection Type
	TCP = "tcp"

	// HttpsCompleteRequest => http request has builded send to Client
	HttpsCompleteRequest = "HTTP/1.1 200 Connection established" + CR

	// HttpsMethod => request method of HTTPS
	HttpsMethod = "CONNECT"
)
