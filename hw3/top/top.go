package top

import (
	"sort"
	"strings"
)

type wordStat struct {
	word  string
	count int
}

func Top10(text string) []string {
	return Top(text, 10)
}

func Top(text string, topCount int) (res []string) {
	res = make([]string, 0, topCount)

	words := prepareText(text)
	wordCounter := wordsCount(words)

	var statistics []wordStat
	for word, count := range wordCounter {
		statistics = append(statistics, wordStat{word, count})
	}

	sort.Slice(statistics, func(i, j int) bool {
		return statistics[i].count > statistics[j].count
	})

	for index, val := range statistics {
		if index >= topCount {
			break
		}

		res = append(res, val.word)
	}

	return
}

func wordsCount(words []string) map[string]int {
	res := make(map[string]int)

	for _, val := range words {
		res[val]++
	}

	return res
}

func prepareText(text string) (res []string) {
	words := strings.Split(strings.ToLower(text), " ")
	for _, val := range words {
		word := removePunctuation(val)

		if len(word) > 0 {
			res = append(res, word)
		}
	}

	return
}

func removePunctuation(s string) (res string) {
	res = s
	for strings.HasSuffix(res, ".") || strings.HasSuffix(res, ",") || strings.HasSuffix(res, "-") {
		res = res[:len(res)-1]
	}
	return
}
