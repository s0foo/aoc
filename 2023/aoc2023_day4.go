package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

		for _, s := range winningNumbers {
			num, _ := strconv.Atoi(s)
			winningNumbersInt = append(winningNumbersInt, num)
		}
		for _, s := range havingNumbers {
			num, _ := strconv.Atoi(s)
			havingNumbersInt = append(havingNumbersInt, num)
		}

		l := len(intersection(winningNumbersInt, havingNumbersInt))
		if l > 0 {
			sum += powInt(2, l-1)
		}
	}

	return sum
}
