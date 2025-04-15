package aoc

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func mapOutput(in int, maps [][]int) int {
	var out int
	id := false
	for _, m := range maps {
		dst := m[0]
		src := m[1]
		length := m[2]
		if (in >= src) && (in < src+length) {
			id = true
			d := in - src
			out = dst + d
		}
	}
	if !id {
		out = in
	}
	return out
}

func day5a(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	lowestLocation := 1000000000

	var seeds []int
	var seedsRaw []string

	var seed2Soil bool
	var seed2SoilMaps [][]int
	var soil2fertilizer bool
	var soil2fertilizerMaps [][]int
	var fertilizer2water bool
	var fertilizer2waterMaps [][]int
	var water2light bool
	var water2lightMaps [][]int
	var light2temperature bool
	var light2temperatureMaps [][]int
	var temperature2humidity bool
	var temperature2humidityMaps [][]int
	var humidity2location bool
	var humidity2locationMaps [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "seeds") {
			seedsRaw = strings.Split(line, ":")
			seeds = str2IntSlice(strings.Fields(seedsRaw[1]))
		}
		if strings.HasPrefix(line, "seed-to-soil") {
			seed2Soil = true
		}
		if strings.HasPrefix(line, "soil-to-fertilizer") {
			soil2fertilizer = true
			seed2Soil = false
		}
		if strings.HasPrefix(line, "fertilizer-to-water") {
			fertilizer2water = true
			soil2fertilizer = false
		}
		if strings.HasPrefix(line, "water-to-light") {
			water2light = true
			fertilizer2water = false
		}
		if strings.HasPrefix(line, "light-to-temperature") {
			light2temperature = true
			water2light = false
		}
		if strings.HasPrefix(line, "temperature-to-humidity") {
			temperature2humidity = true
			light2temperature = false
		}
		if strings.HasPrefix(line, "humidity-to-location") {
			humidity2location = true
			temperature2humidity = false
		}

		if seed2Soil && !(strings.HasPrefix(line, "seed-to-soil")) {
			if len(line) > 1 {
				m := strings.Fields(line)
				n := str2IntSlice(m)
				seed2SoilMaps = append(seed2SoilMaps, n)
			}
		}
		if soil2fertilizer && !(strings.HasPrefix(line, "soil-to-fertilizer")) {
			if len(line) > 1 {
				m := strings.Fields(line)
				n := str2IntSlice(m)
				soil2fertilizerMaps = append(soil2fertilizerMaps, n)
			}
		}
		if fertilizer2water && !(strings.HasPrefix(line, "fertilizer-to-water")) {
			if len(line) > 1 {
				m := strings.Fields(line)
				n := str2IntSlice(m)
				fertilizer2waterMaps = append(fertilizer2waterMaps, n)
			}
		}
		if water2light && !(strings.HasPrefix(line, "water-to-light")) {
			if len(line) > 1 {
				m := strings.Fields(line)
				n := str2IntSlice(m)
				water2lightMaps = append(water2lightMaps, n)
			}
		}
		if light2temperature && !(strings.HasPrefix(line, "light-to-temperature")) {
			if len(line) > 1 {
				m := strings.Fields(line)
				n := str2IntSlice(m)
				light2temperatureMaps = append(light2temperatureMaps, n)
			}
		}
		if temperature2humidity && !(strings.HasPrefix(line, "temperature-to-humidity")) {
			if len(line) > 1 {
				m := strings.Fields(line)
				n := str2IntSlice(m)
				temperature2humidityMaps = append(temperature2humidityMaps, n)
			}
		}
		if humidity2location && !(strings.HasPrefix(line, "humidity-to-location")) {
			if len(line) > 1 {
				m := strings.Fields(line)
				n := str2IntSlice(m)
				humidity2locationMaps = append(humidity2locationMaps, n)
			}
		}
	}

	for _, s := range seeds {
		var out int
		out = mapOutput(s, seed2SoilMaps)
		out = mapOutput(out, soil2fertilizerMaps)
		out = mapOutput(out, fertilizer2waterMaps)
		out = mapOutput(out, water2lightMaps)
		out = mapOutput(out, light2temperatureMaps)
		out = mapOutput(out, temperature2humidityMaps)
		out = mapOutput(out, humidity2locationMaps)

		if out < lowestLocation {
			lowestLocation = out
		}
	}

	return lowestLocation
}
