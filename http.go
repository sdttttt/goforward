package goforward

import (
	"bytes"
	"net/textproto"
	"strings"
)

func ParseHttp(firstLine string, reader *textproto.Reader) (string, *bytes.Buffer) {

	buf := bytes.NewBuffer(make([]byte, 4096))
	targetHost := getTargetHost(firstLine)
	buf.WriteString(firstLine + CR)

	b := make([]byte, 4096)
	reader.R.Read(b)

	println(string(b))

	buf.Write(b)
	return targetHost, buf
}

func getTargetHost(firstLine string) string {
	films := strings.Split(firstLine, SPACE)
	return films[1]
}
