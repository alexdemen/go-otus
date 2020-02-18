package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "localhost:3302")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			defer c.Close()
			io.Copy(os.Stdout, c)
		}(conn)
	}
}
