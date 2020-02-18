package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/spf13/pflag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	timeout := pflag.Int("timeout", 10, "connection timeout")
	pflag.Parse()

	//TODO проверка наличия
	addr := os.Args[len(os.Args)-2]
	port := os.Args[len(os.Args)-1]

	ctxTimeout, _ := context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Second)

	dialer := &net.Dialer{}
	conn, err := dialer.DialContext(ctxTimeout, "tcp", addr+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt)

	ctxCancel, _ := context.WithCancel(context.Background())
	handleConn(ctxCancel, conn)
}

func handleConn(ctx context.Context, conn net.Conn) {
	go write(ctx, conn, streamReader(os.Stdin))
	go read(ctx, conn)
}

func streamReader(reader io.Reader) chan string{
	dataChan := make(chan string)

	go func() {
		scanner := bufio.NewScanner(reader)

		for {
			if !scanner.Scan() {
				break
			}

			data := scanner.Text()
			dataChan <- data
		}
	}()

	return dataChan
}

func read(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !scanner.Scan() {
				return
			}

			text := scanner.Text()
			fmt.Println(text)
		}
	}
}

func write(ctx context.Context, conn net.Conn, input <-chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case data := <- input:
			n, err := conn.Write([]byte(data))
			fmt.Println("writted - ", n)

			if err != nil {
				return
			}
		}
	}
}
