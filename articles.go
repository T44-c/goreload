package main

import "strings"

func fixArticles(line string) string {
	words := strings.Split(line, " ")
	for i := 0; i+1 < len(words); i++ {
		isArticle := words[i] == "a" || words[i] == "A"
		if isArticle && startsWithVowelOrH(words[i+1]) {
			words[i] += "n"
		}
	}
	return strings.Join(words, " ")
}

func startsWithVowelOrH(word string) bool {
	for _, r := range word {
		if isAlnum(r) {
			return strings.ContainsRune("aeiouhAEIOUH", r)
		}
	}
	return false
}
