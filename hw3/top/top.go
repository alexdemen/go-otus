package top

import "strings"

func top10(text string) []string {
	return top(text, 10)
}

func top(text string, topCount int) []string {
	words := prepareText(text)

	return words
}

func wordsCount(words []string) map[string]int {
	res := make(map[string]int)

	return res
}

func prepareText(text string) (res []string) {
	words := strings.Split(strings.ToLower(text), " ")

	for _, val := range words {
		res = append(res, val)
	}

	return
}
