package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var tokenRe = regexp.MustCompile(
	`\((hex|bin|up|low|cap)(?:,\s*(\d+))?\)` +
		`|[^\s.,!?:;]*[\p{L}\p{N}][^\s.,!?:;]*`)


var commands = map[string]func(string) string{
	"hex": func(word string) string { return convertBase(word, 16) },
	"bin": func(word string) string { return convertBase(word, 2) },
	"up":  strings.ToUpper,
	"low": strings.ToLower,
	"cap": capitalize,
}

func applyCommands(line string) string {
	var pieces []string 
	var words []int     
	last := 0           

	for _, m := range tokenRe.FindAllStringSubmatchIndex(line, -1) {
		start, end := m[0], m[1]
		nameStart, nameEnd := m[2], m[3]   
		countStart, countEnd := m[4], m[5] 

		
		if start > last {
			pieces = append(pieces, line[last:start])
		}
		last = end

		
		if nameStart < 0 {
			words = append(words, len(pieces))
			pieces = append(pieces, line[start:end])
			continue
		}

		
		count := 1
		if countStart >= 0 {
			n, err := strconv.Atoi(line[countStart:countEnd])
			if err != nil { 
				n = len(words)
			}
			count = n
		}
		if count > len(words) { 
			count = len(words)
		}

		transform := commands[line[nameStart:nameEnd]]
		for _, idx := range words[len(words)-count:] {
			pieces[idx] = transform(pieces[idx])
		}
		
	}

	pieces = append(pieces, line[last:]) 
	return strings.Join(pieces, "")
}

func convertBase(word string, base int) string {
	n, err := strconv.ParseInt(word, base, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(n, 10)
}

func capitalize(word string) string {
	runes := []rune(strings.ToLower(word))
	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			break
		}
		if unicode.IsDigit(r) {
			break
		}
	}
	return string(runes)
}
