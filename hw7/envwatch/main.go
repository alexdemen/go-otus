package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Не задана переменная для вывода.")
	}

	fmt.Println(os.Getenv(os.Args[1]))
	ioutil.WriteFile("test", []byte(os.Getenv(os.Args[1])), os.ModePerm)
}
