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

func isInputDataCoherent(report []string, index int) bool {
	return !(len(report) == 0 || index >= len(report[0]))
}

func gammaRateAsBinaryString(report []string, index int) string {
	if !isInputDataCoherent(report, index) {
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
	if !isInputDataCoherent(report, index) {
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

func filterReportByBitPositionAndWantedValue(report []string, index int, value string) []string {
	if !isInputDataCoherent(report, index) {
		return []string{}
	}

	filteredReport := []string{}
	for _, e := range report {
		if strings.Split(e, "")[index] == value {
			filteredReport = append(filteredReport, e)
		}
	}
	return filteredReport
}

func filterReportWhenMCVandLCVAreTheSame(report []string, index int, co2ScrubberRating rating) []string {
	filteredReport := []string{}
	for _, entry := range report {
		if strings.Split(entry, "")[index] == co2ScrubberRating.String() {
			filteredReport = append(filteredReport, entry)
		}
	}
	return filteredReport
}

type rating string

func (r rating) String() string {
	return string(r)
}

const (
	oxygenGeneratorRating rating = "1"
	co2ScrubberRating     rating = "0"
)

func findWantedRating(report []string, index int, r rating) (int, error) {
	if len(report) == 0 {
		return -1, errors.New("cannot find the wanted rating on an empty report")
	}
	if len(report) == 1 {
		return stringToInt(report[0]), nil
	}

	mcv, lcv := mostAndLeastCommonValuesInGivenPosition(report, index)

	if mcv == lcv {
		report = filterReportWhenMCVandLCVAreTheSame(report, index, r)
	} else {
		if r == co2ScrubberRating {
			report = filterReportByBitPositionAndWantedValue(report, index, lcv)
		} else {
			report = filterReportByBitPositionAndWantedValue(report, index, mcv)
		}
	}

	return findWantedRating(report, index+1, r)
}
