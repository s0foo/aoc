package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

func day1a() (int, error) {
	file, err := os.Open("data/day1a.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var firstDigit, lastDigit rune
		var calibration string
		foundFirst := false

		for _, char := range line {
			if unicode.IsDigit(char) {
				if !foundFirst {
					firstDigit = char
					foundFirst = true
				}
				lastDigit = char
			}
		}

		calibration = string(firstDigit) + string(lastDigit)
		calibrationValue, err := strconv.Atoi(calibration)
		if err != nil {
			log.Fatalf("cannot convert string %s to int: %v", calibration, err)
		}

		sum += calibrationValue
	}

	return sum, nil
}
