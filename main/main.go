package main

import (
	"bufio"
	"github.com/sdttttt/goforward"
	"net"
)

func main() {

	if ln, err := net.Listen("tcp", ":10086"); err != nil {
		println(err.Error())
	} else {
		for {
			if conn, err := ln.Accept(); err == nil {
				go connhandler(conn)
			} else {
				break
			}
		}
	}
}

func connhandler(conn net.Conn) {

	client := &goforward.ClientConn{
		Reader: bufio.NewReader(conn),
		Conn:   conn,
	}

	connType, targetHost, port, protocol := client.ParseConnectionInfo()

	println("ctype => ", connType)
	println("target => ", targetHost)
	println("port => ", port)
	println("proto => ", protocol)
}
