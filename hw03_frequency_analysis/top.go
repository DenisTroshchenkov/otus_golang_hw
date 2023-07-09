package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var cutSymbols = regexp.MustCompile(`\.|\?|!|,|;|:|\.\.\.|\(|\)`)

type Word struct {
	Value string
	Count uint64
}

func Top10(str string) []string {
	return top10Impl(str, false)
}

func Top10WithAsterisk(str string) []string {
	return top10Impl(str, true)
}

func top10Impl(str string, withAsterisk bool) []string {
	if str == "" {
		return []string{}
	}
	wordsMap := make(map[string]uint64)
	for _, word := range strings.Fields(str) {
		if withAsterisk {
			word = NormalizeWord(word)
			if word == "" || word == "-" {
				continue
			}
			wordsMap[word]++
		} else {
			wordsMap[word]++
		}
	}

	if len(wordsMap) == 0 {
		return []string{}
	}

	wordsSlice := make([]Word, 0, len(wordsMap))
	for word, count := range wordsMap {
		wordsSlice = append(wordsSlice, Word{word, count})
	}

	sort.Slice(
		wordsSlice,
		func(i, j int) bool {
			if wordsSlice[i].Count == wordsSlice[j].Count {
				return strings.Compare(wordsSlice[i].Value, wordsSlice[j].Value) <= 0
			}
			return wordsSlice[i].Count > wordsSlice[j].Count
		},
	)

	maxLen := 10
	if len(wordsSlice) < maxLen {
		maxLen = len(wordsSlice)
	}

	res := make([]string, 0, maxLen)
	for _, word := range wordsSlice[:maxLen] {
		res = append(res, word.Value)
	}

	return res
}

func NormalizeWord(word string) string {
	word = strings.ToLower(word)
	return string(cutSymbols.ReplaceAll([]byte(word), []byte{}))
}
