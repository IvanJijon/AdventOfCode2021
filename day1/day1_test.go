package day1

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type Day1Suite struct{}

var _ = Suite(&Day1Suite{})

func (s *Day1Suite) TestImportData(c *C) {
	depths, _ := importData("day1_input_test")
	c.Assert(depths, DeepEquals, []int{141, 140, 160, 161})
}

func (s *Day1Suite) TestIfNextMeasurmentIncreasesInDepth(c *C) {
	current := 2
	next := 3
	c.Assert(isDeeper(current, next), Equals, true)

	current = 2
	next = 1
	c.Assert(isDeeper(current, next), Equals, false)
}

func (s *Day1Suite) TestCountHowManyMeasurementsAreLargerThanThePreviousOne(c *C) {
	depths, _ := importData("day1_input_test")
	c.Assert(reducePartOne(depths), Equals, 2)

	// Answer for part one:
	depths, _ = importData("day1_input")
	c.Assert(reducePartOne(depths), Equals, 1692)
}

func (s *Day1Suite) TestSumMeasurementsByChunksOfThree(c *C) {
	depths, _ := importData("day1_input_test")
	c.Assert(sumMeasurementsByChunksOfThree(0, depths), Equals, 441)
	c.Assert(sumMeasurementsByChunksOfThree(1, depths), Equals, 461)
}

func (s *Day1Suite) TestCountHowManyMeasurementsAreLargerThanThePreviousOneByChunksOfThree(c *C) {
	depths, _ := importData("day1_input_test")
	c.Assert(reducePartTwo(depths), Equals, 1)

	// Answer for part two
	depths, _ = importData("day1_input")
	c.Assert(reducePartTwo(depths), Equals, 1724)
}
