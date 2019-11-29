package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {

	if res, err := requestTime("0.beevik-ntp.pool.ntp.org"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println(res)
	}

}

func requestTime(host string) (time.Time, error) {
	resTime, err := ntp.Time(host)
	if err != nil {
		return time.Time{}, err
	}

	return resTime, nil
}
