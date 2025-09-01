package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func AnalyzeText(text string) {
	wordsMap := make(map[string]int)
	re := regexp.MustCompile(`[.,!?]`)
	words := strings.Fields(strings.ToLower(re.ReplaceAllString(text, "")))

	for _, word := range words {
		wordsMap[word]++
	}

	countCircleWords := len(words)
	countUniqueWords := len(wordsMap)

	var keyMostPopularWord string
	var valueMostPopularWord int

	for key, value := range wordsMap {
		if value > valueMostPopularWord {
			keyMostPopularWord = key
			valueMostPopularWord = value
		}

	}

	fmt.Printf(`Количество слов: %d
Количество уникальных слов: %d
Самое часто встречающееся слово: "%s" (встречается %d раз)
Топ-5 самых часто встречающихся слов:
`, countCircleWords, countUniqueWords, keyMostPopularWord, valueMostPopularWord)
	topWords := getTopWords(wordsMap, 5)

	for _, word := range topWords {
		fmt.Printf(`"%s": %d раз
`, word, wordsMap[word])

	}
}

func getTopWords(wordMap map[string]int, n int) []string {

	type pair struct {
		word  string
		count int
	}

	var pairs []pair

	for w, c := range wordMap {
		pairs = append(pairs, pair{w, c})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	var result []string

	for i := 0; i < n && i < len(pairs); i++ {
		result = append(result, pairs[i].word)
	}
	return result

}
