package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func main() {
	if res, err := ntp.Time("0.beevik-ntp.pool.ntp.org"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println(res)
	}
}
