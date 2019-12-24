package main

import (
	"fmt"
	"github.com/alexdemen/go-otus/hw3/top"
)

func main() {
	text := "- - - - - - - - 1 1 2 2 3 3 4 4 5 5 6 6 6 6 6 6 6 7 7 8 8 8,"
	topWords := top.Top(text, 5)
	fmt.Println(topWords)
}
