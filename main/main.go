package main

import (
	"bufio"
	"log"
	"net"
	"net/textproto"

	"github.com/sdttttt/goforward"
)

func main() {
	if ln, err := net.Listen("tcp", ":10086"); err != nil {
		log.Println(err.Error())
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
	// implement Text-based protocol
	reader := textproto.NewReader(bufio.NewReader(conn))
	goforward.NewClient(conn, reader)
}
