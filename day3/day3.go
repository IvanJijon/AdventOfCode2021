package day3

import (
	"bufio"
	"os"
	"strconv"
)

func importData(inputFile string) ([]string, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var report []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		report = append(report, scanner.Text())
	}

	return report, scanner.Err()
}

func gammaRateAsBinaryString(report []string, index int) string {
	if len(report) == 0 || index >= len(report[0]) {
		return ""
	}

	mostCommon := "0"
	ones := 0
	zeros := 0
	for _, r := range report {
		if r[index] == '1' {
			ones += 1
			continue
		}
		zeros += 1
	}

	if ones > zeros {
		mostCommon = "1"
	}

	return mostCommon + gammaRateAsBinaryString(report, index+1)
}

func stringToInt(input string) int {
	res, _ := strconv.ParseInt(input, 2, 64)
	return int(res)
}

func epsilonRateFromGammaRate(input string) string {
	complement := ""
	for _, b := range input {
		if b == '1' {
			complement += "0"
			continue
		}
		complement += "1"
	}
	return complement
}
