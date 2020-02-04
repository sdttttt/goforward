package main

import (
	"fmt"
	"net"
)

func main() {

	if ln, err := net.Listen("tcp", ":10086"); err != nil {
		println(err.Error())
	} else {
		for {
			if conn, err := ln.Accept(); err == nil {
				connhandler(conn)
			} else {
				break
			}
		}
	}
}

func connhandler(conn net.Conn) {
	var buf = make([]byte, 2048)

	_, err := conn.Read(buf)
	if err != nil {
		println(err.Error())
	}

	fmt.Println(string(buf))
	println(conn.LocalAddr().String())
	println(conn.RemoteAddr().String())
}
