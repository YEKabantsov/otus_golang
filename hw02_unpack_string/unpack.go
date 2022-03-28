package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	if err := validateString(str); err != nil {
		return "", err
	}

	var result string
	for _, ch := range str {
		if unicode.IsDigit(ch) {
			count, _ := strconv.Atoi(string(ch))

			if count == 0 {
				result = result[:len(result)-1]
				continue
			}

			lastChar := result[len(result)-1:]
			result += strings.Repeat(lastChar, count-1)
		} else {
			result += string(ch)
		}
	}

	return result, nil
}

func validateString(str string) error {
	if unicode.IsDigit(rune(str[0])) {
		return ErrInvalidString
	}

	if _, err := strconv.Atoi(str); err == nil {
		return ErrInvalidString
	}

	if isContainTwoDigitsNear(str) {
		return ErrInvalidString
	}

	onlyLetterAndDigits := regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
	if !onlyLetterAndDigits(str) {
		return ErrInvalidString
	}

	return nil
}

func isContainTwoDigitsNear(str string) bool {
	for i, ch := range str {
		if i == len(str)-1 {
			return false
		}

		if unicode.IsDigit(ch) && unicode.IsDigit(rune(str[i+1])) {
			return true
		}
	}
	return false
}
