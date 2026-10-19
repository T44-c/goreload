package main

import (
	"strings"
	"unicode"
)

const punctuation = ".,!?:;"

func isPunct(r rune) bool {
	return strings.ContainsRune(punctuation, r)
}

func isClosingBracket(r rune) bool {
	return strings.ContainsRune(")]}", r)
}

func isInsideNumber(runes []rune, i int) bool {
	return i > 0 && i+1 < len(runes) &&
		unicode.IsDigit(runes[i-1]) && unicode.IsDigit(runes[i+1])
}

func fixPunctuation(line string) string {
	runes := []rune(line)
	out := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); {
		r := runes[i]
		if !isPunct(r) || isInsideNumber(runes, i) {
			out = append(out, r)
			i++
			continue
		}

		if n := len(out); n > 0 && out[n-1] == ' ' {
			out = out[:n-1]
		}

		for i < len(runes) && isPunct(runes[i]) {
			out = append(out, runes[i])
			i++
		}

		if i < len(runes) && runes[i] == ' ' {
			i++
		}
		if i < len(runes) && !isClosingBracket(runes[i]) {
			out = append(out, ' ')
		}
	}
	return string(out)
}
