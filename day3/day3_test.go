package day3

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type Day3Suite struct{}

var _ = Suite(&Day3Suite{})

func (s *Day3Suite) TestImportData(c *C) {
	report, _ := importData("day3_input_test")
	c.Assert(report, DeepEquals, []string{
		"011001101000",
		"010101111100",
		"000000111101",
		"001001001010",
		"010011000001",
	})
}

func (s *Day3Suite) TestGamaRateAsBinaryString(c *C) {
	report := []string{
		"",
	}
	c.Assert(gammaRateAsBinaryString(report, 0), Equals, "")

	report = []string{
		"000",
		"100",
		"101",
		"010",
		"001",
	}
	c.Assert(gammaRateAsBinaryString(report, 0), Equals, "000")

	report = []string{
		"001",
		"100",
		"101",
		"010",
		"101",
	}
	c.Assert(gammaRateAsBinaryString(report, 0), Equals, "101")

	report, _ = importData("day3_input_test")
	c.Assert(gammaRateAsBinaryString(report, 0), Equals, "010001101000")
}

func (s *Day3Suite) TestStringToInt(c *C) {
	r := ""
	c.Assert(stringToInt(r), Equals, 0)
	r = "000"
	c.Assert(stringToInt(r), Equals, 0)
	r = "001"
	c.Assert(stringToInt(r), Equals, 1)
	r = "010001101000"
	c.Assert(stringToInt(r), Equals, 1128)
}

func (s *Day3Suite) TestEpsilonRateFromGammaRate(c *C) {
	gammaRate := ""
	c.Assert(epsilonRateFromGammaRate(gammaRate), Equals, "")
	gammaRate = "000"
	c.Assert(epsilonRateFromGammaRate(gammaRate), Equals, "111")
	gammaRate = "111"
	c.Assert(epsilonRateFromGammaRate(gammaRate), Equals, "000")
	gammaRate = "1010"
	c.Assert(epsilonRateFromGammaRate(gammaRate), Equals, "0101")
}

func (s *Day3Suite) TestAnswerDay3PartOne(c *C) {

	report, _ := importData("day3_input")
	gammaRate := gammaRateAsBinaryString(report, 0)
	epsilonRate := epsilonRateFromGammaRate(gammaRate)

	gr := stringToInt(gammaRate)
	er := stringToInt(epsilonRate)

	c.Assert(gr*er, Equals, 3374136)
}

func (s *Day3Suite) TestMostAndLeastCommonValuesInGivenPosition(c *C) {
	report := []string{
		"",
	}
	index := 0

	mcv, lcv := mostAndLeastCommonValuesInGivenPosition(report, index)
	c.Assert(mcv, Equals, "")
	c.Assert(lcv, Equals, "")

	report = []string{
		"001",
		"100",
		"101",
		"010",
		"101",
		"100",
	}
	mcv, lcv = mostAndLeastCommonValuesInGivenPosition(report, 0)
	c.Assert(mcv, Equals, "1")
	c.Assert(lcv, Equals, "0")
	mcv, lcv = mostAndLeastCommonValuesInGivenPosition(report, 1)
	c.Assert(mcv, Equals, "0")
	c.Assert(lcv, Equals, "1")
	mcv, lcv = mostAndLeastCommonValuesInGivenPosition(report, 2)
	c.Assert(mcv, Equals, "=")
	c.Assert(lcv, Equals, "=")
}
