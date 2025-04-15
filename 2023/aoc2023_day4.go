package aoc

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func day4a(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var winningNumbersInt []int
		var havingNumbersInt []int

		line := scanner.Text()

		card := strings.Split(line, ":")
		numbers := strings.Split(card[1], "|")
		winningNumbers := strings.Fields(numbers[0])
		havingNumbers := strings.Fields(numbers[1])

		winningNumbersInt = str2IntSlice(winningNumbers)
		havingNumbersInt = str2IntSlice(havingNumbers)

		l := len(intersection(winningNumbersInt, havingNumbersInt))
		if l > 0 {
			sum += powInt(2, l-1)
		}
	}

	return sum
}

func day4b(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	m := make(map[int]int)
	sum := 0
	cl := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var winningNumbersInt []int
		var havingNumbersInt []int

		line := scanner.Text()

		card := strings.Split(line, ":")
		numbers := strings.Split(card[1], "|")
		winningNumbers := strings.Fields(numbers[0])
		havingNumbers := strings.Fields(numbers[1])

		winningNumbersInt = str2IntSlice(winningNumbers)
		havingNumbersInt = str2IntSlice(havingNumbers)

		l := len(intersection(winningNumbersInt, havingNumbersInt))
		m[cl] += 1
		if l > 0 {
			for i := 1; i <= l; i++ {
				m[cl+i] += m[cl]
			}
		}
		cl++
	}

	for _, v := range m {
		sum += v
	}

	return sum
}
