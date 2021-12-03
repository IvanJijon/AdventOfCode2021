package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type command struct {
	instruction string
	value       int
}

type sumbarine struct {
	horizontalPosition int
	depth              int
}

func importData(inputFile string) ([]command, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var commands []command
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		c := strings.Split(scanner.Text(), " ")
		i := c[0]
		v, _ := strconv.Atoi(c[1])
		commands = append(commands, command{i, v})
	}

	return commands, scanner.Err()
}

func (s *sumbarine) execute(c command) {
	switch c.instruction {
	case "forward":
		s.horizontalPosition += c.value
	case "up":
		s.depth -= c.value
	case "down":
		s.depth += c.value
	}
}

func (s *sumbarine) pilot(commands []command) {
	for _, c := range commands {
		s.execute(c)
	}
}
