package top

import "strings"

func top10(text string) []string{
	return top(text, 10)
}

func top(text string, topCount int) []string{
	return prepareText(text)
}

func prepareText(text string) (res []string){
	res = strings.Split(text, " ")
	return
}
