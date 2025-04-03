package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func day1a() (int, error) {
	file, err := os.Open("data/day1.txt")
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

func day1b() (int, error) {
	file, err := os.Open("data/day1.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	sum := 0

	var numbers = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var calibration string
		var firstDigit, lastDigit string
		var foundFirst = false
		var indF = 0
		var indL = 0
		var newPosF = 0
		var newPosL = 0
		var posF = 100
		var posL = 0

		for ind, num := range numbers {
			newPosF = strings.Index(line, num)
			if newPosF >= 0 && newPosF < posF {
				posF = newPosF
				indF = ind
			}
			newPosL = strings.LastIndex(line, num)
			if newPosL > posL {
				posL = newPosL
				indL = ind
			}
		}

		for i, char := range line {
			if unicode.IsDigit(char) {
				if !foundFirst {
					if i < posF {
						firstDigit = string(char)
					} else {
						firstDigit = strconv.Itoa(indF + 1)
					}
					foundFirst = true
				}
				if i >= posL {
					lastDigit = string(char)
				} else {
					lastDigit = strconv.Itoa(indL + 1)
				}
			}
		}

		calibration = firstDigit + lastDigit
		calibrationValue, err := strconv.Atoi(calibration)
		if err != nil {
			log.Fatalf("cannot convert string %s to int: %v", calibration, err)
		}

		sum += calibrationValue
	}

	return sum, nil
}
