package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	strInRunes := []rune(str)
	var res strings.Builder

	for i := 0; i < len(strInRunes); i++ {
		nextRuneIdx := i + 1
		repeatCount := 1
		currentRune := strInRunes[i]

		if unicode.IsDigit(currentRune) {
			return "", ErrInvalidString
		}

		if currentRune == '\\' {
			if nextRuneIdx >= len(strInRunes) ||
				(!unicode.IsDigit(strInRunes[nextRuneIdx]) &&
					strInRunes[nextRuneIdx] != '\\') {
				return "", ErrInvalidString
			}
			i++
			nextRuneIdx++
			currentRune = strInRunes[i]
		}

		if nextRuneIdx < len(strInRunes) && unicode.IsDigit(strInRunes[nextRuneIdx]) {
			var err error
			repeatCount, err = strconv.Atoi(string(strInRunes[nextRuneIdx]))
			if err != nil {
				return "", err
			}
			i++
		}
		res.WriteString(strings.Repeat(string(currentRune), repeatCount))
	}
	return res.String(), nil
}
