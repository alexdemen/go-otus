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

	fmt.Println(os.Getenv(os.Args[1]))
}
