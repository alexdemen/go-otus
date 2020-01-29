package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Не задана переменная для вывода.")
	}

	for _, key := range os.Args[1:] {
		fmt.Println(key + "=" + os.Getenv(key))
	}
}
