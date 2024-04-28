package utils

import "strings"

func CleanInput(str string) string {
	lowerCaseStrg := strings.ToLower(str)
	words := strings.Fields(lowerCaseStrg)
	firstWord := ""
	if len(words) != 0 {
		firstWord = words[0]
	}
	return firstWord
}
