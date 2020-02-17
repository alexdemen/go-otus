package main

import (
	"bufio"
	"context"
	"github.com/spf13/pflag"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	timeout := pflag.Int("timeout", 10, "connection timeout")
	pflag.Parse()

	addr := os.Args[len(os.Args)-2]
	port := os.Args[len(os.Args)-1]
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Second)

	dialer := &net.Dialer{}
	conn, err := dialer.DialContext(ctx, "tcp", addr+":"+port)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	handleConn(ctx, conn)

	time.Sleep(5 * time.Minute)

}

func handleConn(ctx context.Context, conn net.Conn) {
	go write(ctx, conn)
}

func write(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		select {
		default:
			if !scanner.Scan() {
				return
			}

			data := scanner.Text()
			_, err := conn.Write([]byte(data))
			
			if err != nil {
				return
			}
		}
	}
}
