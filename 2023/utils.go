package aoc

import (
	"math"
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
			if (i == len(s)-1) && (len(numStr) > 0) {
				lineCoord = append(lineCoord, i)
				num, _ = strconv.Atoi(numStr)
				lineSet = append(lineSet, num)
			}
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

func extractLineNumbersGear(s string) ([]int, []int, []int) {
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
			if (i == len(s)-1) && (len(numStr) > 0) {
				lineCoord = append(lineCoord, i)
				num, _ = strconv.Atoi(numStr)
				lineSet = append(lineSet, num)
			}
		} else {
			if char == '*' {
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

func intersection(u, v []int) []int {
	m := make(map[int]bool)
	var c []int

	for _, e := range u {
		m[e] = true
	}

	for _, e := range v {
		if m[e] {
			c = append(c, e)
		}
	}

	return c
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
