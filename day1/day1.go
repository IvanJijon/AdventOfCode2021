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

func reduce(depths []int) int {
	counter := 0
	for current := 0; current < len(depths)-1; current++ {
		if isDeeper(depths[current], depths[current+1]) {
			counter++
		}
	}

	return counter
}
