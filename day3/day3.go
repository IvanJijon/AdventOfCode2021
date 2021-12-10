package day3

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
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
	mcv, _ := mostAndLeastCommonValuesInGivenPosition(report, index)

	return mcv + gammaRateAsBinaryString(report, index+1)
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

func mostAndLeastCommonValuesInGivenPosition(report []string, index int) (string, string) {
	if len(report) == 0 || index >= len(report[0]) {
		return "", ""
	}
	mostCommon := "="
	leastCommon := "="
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
		leastCommon = "0"
	} else if ones < zeros {
		mostCommon = "0"
		leastCommon = "1"
	}

	return mostCommon, leastCommon
}

func filterReportByBitPositionAndWantedValue(report []string, position int, value string) []string {
	if len(report) == 0 || position >= len(report[0]) {
		return []string{}
	}

	filteredReport := []string{}
	for _, e := range report {
		if strings.Split(e, "")[position] == value {
			filteredReport = append(filteredReport, e)
		}
	}
	return filteredReport
}

func filterReportWhenMCVandLCVAreTheSame(report []string, position int, dominant string) []string {
	filteredReport := []string{}
	for _, entry := range report {
		if strings.Split(entry, "")[position] == dominant { // pasarle a tipe
			filteredReport = append(filteredReport, entry)
		}
	}
	return filteredReport
}

func findTheOxygenGeneratorRating(report []string, position int) (int, error) {
	if len(report) == 0 {
		return -1, errors.New("cannot find the Oxygen Generator Rating on an empty report")
	}
	if len(report) == 1 {
		return stringToInt(report[0]), nil
	}

	mcv, lcv := mostAndLeastCommonValuesInGivenPosition(report, position)

	if mcv == lcv {
		report = filterReportWhenMCVandLCVAreTheSame(report, position, "1")
	} else {
		report = filterReportByBitPositionAndWantedValue(report, position, mcv)
	}

	return findTheOxygenGeneratorRating(report, position+1)
}

func findTheCO2ScrubberRating(report []string, position int) (int, error) {
	if len(report) == 0 {
		return -1, errors.New("cannot find the CO2 Scrubber Rating on an empty report")
	}
	if len(report) == 1 {
		return stringToInt(report[0]), nil
	}

	mcv, lcv := mostAndLeastCommonValuesInGivenPosition(report, position)

	if mcv == lcv {
		report = filterReportWhenMCVandLCVAreTheSame(report, position, "0")
	} else {
		report = filterReportByBitPositionAndWantedValue(report, position, lcv)
	}

	return findTheCO2ScrubberRating(report, position+1)
}
