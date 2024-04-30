package main

import (
	"strings"
	"unicode"
)

func CountBytes(file []byte) int {
	return len(file)
}

func CountLines(file []byte) (lineCount int) {
	for i := 0; i < len(file); i++ {
		if file[i] == '\n' {
			lineCount++
		}
		if file[i] == '\r' {
			if i+1 < len(file) && file[i+1] == '\n' {
				continue
			}
			lineCount++
		}
	}
	return
}

func CountWords(file []byte) int {
	words := string(file)
	isWordSeparator := func(r rune) bool {
		return unicode.IsSpace(r) || r == '\n'
	}
	wordSlice := strings.FieldsFunc(words, isWordSeparator)
	return len(wordSlice)
}

func CountCharacters(file []byte) int {
	words := string(file)
	charSlice := strings.Split(words, "")
	return len(charSlice)
}
