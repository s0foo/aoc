package aoc

import (
    //"strings"
    "strconv"
    "unicode"
)

func extractNumbers(s string) int {
	var numStr string
	for _, char := range s {
		if unicode.IsDigit(char) {
			numStr += string(char)
		}
	}

	if len(numStr) == 0 {
		return 0
	}

	num, _ := strconv.Atoi(numStr)

	return num
}
