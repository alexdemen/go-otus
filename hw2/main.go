package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	str, err := unpack(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}

func unpack(str string) (string, error) {
	runes := []rune(str)
	res := make([]rune, 0)
	escaping := false

	for i := 0; i < len(runes); i++ {
		curRune := runes[i]
		if unicode.IsLetter(curRune) || escaping {
			if next := i + 1; next < len(runes) && unicode.IsDigit(runes[next]) {
				repeat, _ := strconv.Atoi(string(runes[next]))
				res = repeatRune(res, curRune, repeat)
				i++
			} else {
				res = repeatRune(res, curRune, 1)
			}
			escaping = false
		} else if curRune == '\\' {
			escaping = true
		} else {
			return "", fmt.Errorf("%s is invalid string", str)
		}
	}

	return string(res), nil
}

func repeatRune(dst []rune, symbol rune, count int) []rune {
	res := dst
	for ; count != 0; count-- {
		res = append(res, symbol)
	}
	return res
}
