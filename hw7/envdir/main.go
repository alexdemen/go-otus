package main

import (
	"github.com/alexdemen/hw7/envdir/enviroment"
	"github.com/alexdemen/hw7/envdir/executor"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Неверное количество переданных элементов.")
	}

	envs, err := enviroment.ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(executor.RunCmd(os.Args[2:], envs))
}
