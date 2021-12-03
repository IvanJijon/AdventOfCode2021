package day1

import (
	"bufio"
	"os"
	"strconv"
)

func importData(inputFile string) ([]int, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var depths []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		d, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, d)
	}

	return depths, scanner.Err()
}

func isDeeper(current, next int) bool {
	return next > current
}

func reducePartOne(depths []int) int {
	counter := 0
	for current := 0; current < len(depths)-1; current++ {
		if isDeeper(depths[current], depths[current+1]) {
			counter++
		}
	}

	return counter
}

func reducePartTwo(depths []int) int {
	counter := 0
	for current := 0; current < len(depths)-3; current++ {
		if isDeeper(sumMeasurementsByChunksOfThree(current, depths), sumMeasurementsByChunksOfThree(current+1, depths)) {
			counter++
		}
	}

	return counter
}

func sumMeasurementsByChunksOfThree(index int, depths []int) int {
	return depths[index] + depths[index+1] + depths[index+2]
}
