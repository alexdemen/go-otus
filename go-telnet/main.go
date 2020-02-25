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

	argCount := len(os.Args)
	if argCount < 3 {
		log.Fatal("incorrect count of args")
	}
	addr := os.Args[argCount-2]
	port := os.Args[argCount-1]

	conn, err := net.DialTimeout("tcp", addr+":"+port, time.Duration(*timeout)*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ctxCancel, cancel := context.WithCancel(context.Background())

	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt)

	handleConn(ctxCancel, conn, cancel)

	select {
	case <-ctxCancel.Done():
	case <-done:
		ctxCancel.Done()
	}

}

func handleConn(ctx context.Context, conn net.Conn, cancel context.CancelFunc) {
	go write(ctx, conn, streamReader(os.Stdin, cancel))

	go func() {
		err := read(ctx, conn, streamReader(conn, cancel))
		if err != nil {
			cancel()
		}
	}()
}

func streamReader(reader io.Reader, cancel context.CancelFunc) chan string {
	dataChan := make(chan string)

	go func() {
		scanner := bufio.NewScanner(reader)

		for {
			if !scanner.Scan() {
				cancel()
				break
			}

			data := scanner.Text()
			dataChan <- data
		}
	}()

	return dataChan
}

func read(ctx context.Context, conn net.Conn, input <-chan string) error {

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			select {
			case data := <-input:
				fmt.Println(data)
			default:
			}
		}
	}
}

func write(ctx context.Context, conn net.Conn, input <-chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			select {
			case data := <-input:
				_, err := conn.Write([]byte(data))

				if err != nil {
					return
				}
			default:
			}
		}
	}
}
