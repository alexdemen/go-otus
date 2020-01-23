package main

import (
	"flag"
	"github.com/alexdemen/go-otus/hw6/fcopy/copiyng"
	"log"
)

func main() {
	if err := copiyng.Copy(*srcFile, *dstFile, *offset, *limit); err != nil {
		log.Fatal(err)
	}
}

var srcFile = flag.String("src", "", "Файл для копирования.")
var dstFile = flag.String("dst", "", "Файл-результат копирования.")
var offset = flag.Int("offset", 0, "Позиция начала копирования.")
var limit = flag.Int("limit", 0, "Количество копируемых байт.")

func init() {
	flag.Parse()
}
