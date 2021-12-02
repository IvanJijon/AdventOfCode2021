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
	c.Assert(reduce(depths), Equals, 2)

	// Answer:
	depths, _ = importData("day1_input")
	c.Assert(reduce(depths), Equals, 1692)
}
