package main

func fixQuotes(line string) string {
	runes := []rune(line)

	var marks []int
	for i, r := range runes {
		if r == '\'' && !isApostrophe(runes, i) {
			marks = append(marks, i)
		}
	}

	drop := make([]bool, len(runes))
	for p := 0; p+1 < len(marks); p += 2 {
		opening, closing := marks[p], marks[p+1]
		for i := opening + 1; i < closing && runes[i] == ' '; i++ {
			drop[i] = true
		}
		for i := closing - 1; i > opening && runes[i] == ' '; i-- {
			drop[i] = true
		}
	}

	out := make([]rune, 0, len(runes))
	for i, r := range runes {
		if !drop[i] {
			out = append(out, r)
		}
	}
	return string(out)
}

func isApostrophe(runes []rune, i int) bool {
	return i > 0 && i+1 < len(runes) && isAlnum(runes[i-1]) && isAlnum(runes[i+1])
}
