package top

import (
	"sort"
	"strings"
)

type wordStat struct{
	word string
	count int
}

func top10(text string) []string {
	return top(text, 10)
}

func top(text string, topCount int) (res []string) {
	words := prepareText(text)
	wordCounter := wordsCount(words)

	var statistics []wordStat
	for word, count := range wordCounter{
		statistics = append(statistics, wordStat{word, count})
	}

	sort.Slice(statistics, func(i,j int) bool{
		return statistics[i].count < statistics[j].count
	})

	for index, val := range statistics{
		if index >= topCount {
			break
		}

		res = append(res, val.word)
	}

	return
}

func wordsCount(words []string) map[string]int {
	res := make(map[string]int)

	for _, val := range words{
		res[val]++
	}

	return res
}

func prepareText(text string) (res []string) {
	words := strings.Split(strings.ToLower(text), " ")
	res = words
	//for _, val := range words {
	//	res = append(res, val)
	//}

	return
}
