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
	aim                int
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

func (s *sumbarine) rudimentarilyExecute(c command) {
	switch c.instruction {
	case "forward":
		s.horizontalPosition += c.value
	case "up":
		s.depth -= c.value
	case "down":
		s.depth += c.value
	}
}

func (s *sumbarine) sophisticatedlyExecute(c command) {
	switch c.instruction {
	case "forward":
		s.horizontalPosition += c.value
		s.depth += s.aim * c.value
	case "up":
		s.aim -= c.value
	case "down":
		s.aim += c.value
	}
}

func (s *sumbarine) pilot(commands []command, f func(c command)) {
	for _, c := range commands {
		f(c)
	}
}
