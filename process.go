package main

import (
	"strings"
	"unicode"
)

func process(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = processLine(line)
	}
	return strings.Join(lines, "\n")
}

func processLine(line string) string {
	line = applyCommands(line)   
	line = normalizeSpaces(line) 
	line = fixPunctuation(line)  
	line = fixQuotes(line)       
	line = fixArticles(line)     
	return line
}

func normalizeSpaces(line string) string {
	return strings.Join(strings.Fields(line), " ")
}

func isAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
