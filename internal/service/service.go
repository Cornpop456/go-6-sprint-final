package service

import (
	"strings"

	"github.com/Cornpop456/go-6-sprint-final/pkg/morse"
)

func isMorse(input string) bool {
	validMorseChars := ".- "

	for _, r := range input {
		if !strings.ContainsRune(validMorseChars, r) {
			return false
		}
	}
	return true
}

func Convert(s string) string {
	if isMorse(s) {
		return morse.ToText(s)
	}

	return morse.ToMorse(s)
}
