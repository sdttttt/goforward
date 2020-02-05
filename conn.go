package goforward

import (
	"bytes"
	"log"
	"net"
	"net/textproto"
	"strings"

	"github.com/sdttttt/goforward/ctype"
)

func NewClient(conn net.Conn, reader *textproto.Reader) *ClientConn {

	connType, firstLine := parseConnectionInfo(reader)
	targetHost, buildRequest := handleConnectInfomation(firstLine, reader, connType)

	return &ClientConn{
		conn:         conn,
		reader:       reader,
		connType:     connType,
		targetHost:   targetHost,
		buildRequest: buildRequest,
	}
}

// Parse Connection Protocol
// return
//Protocol Type,
// Connection Header infomation,
// FirstLine of Connection infomation
func parseConnectionInfo(reader *textproto.Reader) (ctype.ProtocolType, string) {

	// Example: CONNECT www.google.com HTTP/1.1
	firstLine, err := reader.ReadLine()
	if err != nil {
		log.Println(err.Error())
	}
	protocolType := checkProtocol(firstLine)

	return protocolType, firstLine
}

// Check infomation of connect request FirstLine
// Support: HTTP, HTTPS
func checkProtocol(infomation string) ctype.ProtocolType {

	films := strings.Split(infomation, SPACE)

	switch films[0] {
	case "CONNECT":
		return ctype.HTTPS
	default:
		return ctype.HTTP
	}
}

func handleConnectInfomation(firstLine string, reader *textproto.Reader,
	connType ctype.ProtocolType) (string, *bytes.Buffer) {

	switch connType {
	case ctype.HTTPS:
		return ParseHttp(firstLine, reader)
	case ctype.HTTP:
		return ParseHttp(firstLine, reader)
	default:
		return ParseHttp(firstLine, reader)
	}

}

// Client Connection
type ClientConn struct {
	reader       *textproto.Reader
	conn         net.Conn
	connType     ctype.ProtocolType
	targetHost   string
	buildRequest *bytes.Buffer
}

func (cc *ClientConn) BuildTunnel(remoteConn net.Conn) {

}