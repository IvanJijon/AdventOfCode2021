package day2

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type Day2Suite struct{}

var _ = Suite(&Day2Suite{})

func (s *Day2Suite) TestImportData(c *C) {
	commands, _ := importData("day2_input_test")
	c.Assert(commands, DeepEquals, []command{
		{"forward", 5},
		{"down", 5},
		{"forward", 8},
		{"up", 3},
		{"down", 8},
		{"forward", 2},
	})
}

func (s *Day2Suite) TestCommandRudimentaryExecution(c *C) {
	submarine := &sumbarine{0, 0, 0}
	com := command{"forward", 5}
	submarine.rudimentarilyExecute(com)
	c.Assert(submarine.horizontalPosition, Equals, 5)
	com = command{"down", 5}
	submarine.rudimentarilyExecute(com)
	c.Assert(submarine.depth, Equals, 5)
	com = command{"up", 3}
	submarine.rudimentarilyExecute(com)
	c.Assert(submarine.depth, Equals, 2)
}

func (s *Day2Suite) TestCommandSophisticatedExecution(c *C) {
	submarine := &sumbarine{0, 0, 0}
	com := command{"forward", 5}
	submarine.sophisticatedlyExecute(com)
	c.Assert(submarine.horizontalPosition, Equals, 5)
	com = command{"down", 5}
	submarine.sophisticatedlyExecute(com)
	c.Assert(submarine.aim, Equals, 5)
	com = command{"forward", 8}
	submarine.sophisticatedlyExecute(com)
	c.Assert(submarine.horizontalPosition, Equals, 13)
	c.Assert(submarine.depth, Equals, 8*5)
}

func (s *Day2Suite) TestPilotFollowingASeriesOfCommandsRudimentary(c *C) {
	submarine := &sumbarine{0, 0, 0}
	commands, _ := importData("day2_input_test")

	submarine.pilot(commands, submarine.rudimentarilyExecute)

	c.Assert(submarine.horizontalPosition, Equals, 15)
	c.Assert(submarine.depth, Equals, 10)

	// Answer for part one
	submarine.horizontalPosition = 0
	submarine.depth = 0
	commands, _ = importData("day2_input")

	submarine.pilot(commands, submarine.rudimentarilyExecute)

	c.Assert(submarine.horizontalPosition, Equals, 1970)
	c.Assert(submarine.depth, Equals, 916)

	//1970 x 916 = 1804520
}

func (s *Day2Suite) TestPilotFollowingASeriesOfCommandsSophisticatedly(c *C) {
	submarine := &sumbarine{0, 0, 0}
	commands, _ := importData("day2_input_test")

	submarine.pilot(commands, submarine.sophisticatedlyExecute)

	c.Assert(submarine.horizontalPosition, Equals, 15)
	c.Assert(submarine.depth, Equals, 60)

	// Answer for part two
	submarine.horizontalPosition = 0
	submarine.depth = 0
	submarine.aim = 0
	commands, _ = importData("day2_input")

	submarine.pilot(commands, submarine.sophisticatedlyExecute)

	c.Assert(submarine.horizontalPosition, Equals, 1970)
	c.Assert(submarine.depth, Equals, 1000556)

	//1970 x 1000556 = 1971095320
}
