package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func day1a() int {
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

	return sum
}

func day1b() int {
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

	return sum
}

func day2a() int {
	file, err := os.Open("data/day2.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	sum := 0
	var split []string
	var sets []string
	var cubes []string
	var game int
	redValue := 0
	greenValue := 0
	blueValue := 0
	redMAX := 12
	greenMAX := 13
	blueMAX := 14

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		split = strings.Split(line, ":")
		game = extractNumbers(split[0])

		sets = strings.Split(split[1], ";")
		redValue = 0
		greenValue = 0
		blueValue = 0

		for _, set := range sets {
			cubes = strings.Split(set, ",")
			for _, atom := range cubes {
				if strings.Contains(atom, "red") {
					if extractNumbers(atom) > redMAX {
						redValue = extractNumbers(atom)
					}
				}
				if strings.Contains(atom, "green") {
					if extractNumbers(atom) > greenMAX {
						greenValue = extractNumbers(atom)
					}
				}
				if strings.Contains(atom, "blue") {
					if extractNumbers(atom) > blueMAX {
						blueValue = extractNumbers(atom)
					}
				}
			}
		}

		if (redValue < redMAX) && (greenValue < greenMAX) && (blueValue < blueMAX) {
			sum += game
		}
	}

	return sum
}

func day2b() int {
	file, err := os.Open("data/day2.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	sum := 0
	var split []string
	var sets []string
	var cubes []string
	//	var game int
	redValue := 0
	greenValue := 0
	blueValue := 0
	redMAX := 0
	greenMAX := 0
	blueMAX := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		split = strings.Split(line, ":")
		//game = extractNumbers(split[0])

		sets = strings.Split(split[1], ";")
		redValue = 0
		greenValue = 0
		blueValue = 0
		redMAX = 0
		greenMAX = 0
		blueMAX = 0

		for _, set := range sets {
			cubes = strings.Split(set, ",")
			for _, atom := range cubes {
				if strings.Contains(atom, "red") {
					redValue = extractNumbers(atom)
					if redValue > redMAX {
						redMAX = redValue
					}
				}
				if strings.Contains(atom, "green") {
					greenValue = extractNumbers(atom)
					if greenValue > greenMAX {
						greenMAX = greenValue
					}
				}
				if strings.Contains(atom, "blue") {
					blueValue = extractNumbers(atom)
					if blueValue > blueMAX {
						blueMAX = blueValue
					}
				}
			}
		}

		sum += redMAX * greenMAX * blueMAX
	}

	return sum
}

func day3a(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	sum := 0

	var lineNumbers [][]int
	var lineCoord [][]int
	var linePos [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s, t, u := extractLineNumbers(line)
		lineNumbers = append(lineNumbers, s)
		lineCoord = append(lineCoord, t)
		linePos = append(linePos, u)
	}

	above := false
	middle := false
	below := false

	start := 0
	end := 0

	// first line
	for n, s := range lineNumbers[0] {
		start = lineCoord[0][2*n]
		end = lineCoord[0][2*n+1]
		for _, p := range linePos[0] {
			if (p == start-1) || ((p >= start) && (p <= end)) || (p == end+1) {
				middle = true
			}
		}
		for _, p := range linePos[1] {
			if (p == start-1) || ((p >= start) && (p <= end)) || (p == end+1) {
				below = true
			}
		}
		if middle || below {
			sum += s
		}
		middle = false
		below = false
	}

	// middle lines
	lastLine := len(lineCoord) - 1
	for c := 1; c < lastLine; c++ {
		for n, s := range lineNumbers[c] {
			start = lineCoord[c][2*n]
			end = lineCoord[c][2*n+1]
			for _, p := range linePos[c-1] {
				if (p == start-1) || ((p >= start) && (p <= end)) || (p == end+1) {
					above = true
				}
			}
			for _, p := range linePos[c] {
				if (p == start-1) || ((p >= start) && (p <= end)) || (p == end+1) {
					middle = true
				}
			}
			for _, p := range linePos[c+1] {
				if (p == start-1) || ((p >= start) && (p <= end)) || (p == end+1) {
					below = true
				}
			}
			if above || middle || below {
				sum += s
			}
			above = false
			middle = false
			below = false
		}
	}

	// last line
	for n, s := range lineNumbers[lastLine] {
		start = lineCoord[lastLine][2*n]
		end = lineCoord[lastLine][2*n+1]
		for _, p := range linePos[lastLine-1] {
			if (p == start-1) || ((p >= start) && (p <= end)) || (p == end+1) {
				above = true
			}
		}
		for _, p := range linePos[lastLine] {
			if (p == start-1) || ((p >= start) && (p <= end)) || (p == end+1) {
				middle = true
			}
		}
		if above || middle {
			sum += s
		}
		above = false
		below = false
	}

	return sum
}

func day3b(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	gearRatio := 0

	var lineNumbers [][]int
	var lineCoord [][]int
	var gearPos [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s, t, u := extractLineNumbersGear(line)
		lineNumbers = append(lineNumbers, s)
		lineCoord = append(lineCoord, t)
		gearPos = append(gearPos, u)
	}

	start := 0
	end := 0

	m := make(map[int][]int)

	// first line
	for n, s := range lineNumbers[0] {
		start = lineCoord[0][2*n]
		end = lineCoord[0][2*n+1]
		for i, p := range gearPos[0] {
			if (p == start-1) || (p == end+1) {
				m[i] = append(m[i], s)
			}
		}
		for i, p := range gearPos[1] {
			if (p == start-1) || ((p >= start) && (p <= end)) || (p == end+1) {
				m[1000+i] = append(m[1000+i], s)
			}
		}
	}

	// middle lines
	lastLine := len(lineCoord) - 1
	for c := 1; c < lastLine; c++ {
		for n, s := range lineNumbers[c] {
			start = lineCoord[c][2*n]
			end = lineCoord[c][2*n+1]
			for i, p := range gearPos[c-1] {
				if (p == start-1) || ((p >= start) && (p <= end)) || (p == end+1) {
					m[(c-1)*1000+i] = append(m[(c-1)*1000+i], s)
				}
			}
			for i, p := range gearPos[c] {
				if (p == start-1) || (p == end+1) {
					m[c*1000+i] = append(m[c*1000+i], s)
				}
			}
			for i, p := range gearPos[c+1] {
				if (p == start-1) || ((p >= start) && (p <= end)) || (p == end+1) {
					m[(c+1)*1000+i] = append(m[(c+1)*1000+i], s)
				}
			}
		}
	}

	// last line
	for n, s := range lineNumbers[lastLine] {
		start = lineCoord[lastLine][2*n]
		end = lineCoord[lastLine][2*n+1]
		for i, p := range gearPos[lastLine-1] {
			if (p == start-1) || ((p >= start) && (p <= end)) || (p == end+1) {
				m[((lastLine)-1)*1000+i] = append(m[((lastLine)-1)*1000+i], s)
			}
		}
		for i, p := range gearPos[lastLine] {
			if (p == start-1) || (p == end+1) {
				m[(lastLine)*1000+i] = append(m[(lastLine)*1000+i], s)
			}
		}
	}

	for i, _ := range m {
		if len(m[i]) == 2 {
			gearRatio += m[i][0] * m[i][1]
		}
	}

	return gearRatio
}
