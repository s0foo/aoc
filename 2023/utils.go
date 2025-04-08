package aoc

import (
	//"fmt"
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

func extractLineNumbers(s string) ([]int, []int, []int) {
	var lineSet []int
	var lineCoord []int
	var symbolPos []int
	var num int	
	var numStr string

	for i, char := range s {
		if unicode.IsDigit(char) {
			if len(numStr) == 0 {
				lineCoord = append(lineCoord, i)
			}
			numStr += string(char)
		} else {
			if char != '.' {
				symbolPos = append(symbolPos, i)
			}
			if len(numStr) > 0 {
				lineCoord = append(lineCoord, i-1)
				num, _ = strconv.Atoi(numStr)
				lineSet = append(lineSet, num)
			}
			numStr = ""
		}
	}

	return lineSet, lineCoord, symbolPos
}
