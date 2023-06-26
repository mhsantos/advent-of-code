// day-1-1.go solves the first part of day 1 of Advent of Code 2021
// https://adventofcode.com/2021/day/1
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mhsantos/advent-2021/internal/utils"
)

func main() {
	lines := utils.ReadInput("day-2-1.txt")
	fmt.Printf("Factor: %d\n", calculatePositionFactor(lines))
	fmt.Printf("Factor aim: %d\n", calculateFactorUsingAim(lines))
}

// Part 1
func calculatePositionFactor(lines []string) int {
	horizontal, depth := 0, 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		command := parts[0]
		units, err := strconv.Atoi(parts[1])
		if err != nil {
			return -1
		}
		switch command {
		case "forward":
			horizontal += units
		case "down":
			depth += units
		case "up":
			depth -= units
		}
	}
	return horizontal * depth
}

// Part 2
func calculateFactorUsingAim(lines []string) int {
	horizontal, depth, aim := 0, 0, 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		command := parts[0]
		units, err := strconv.Atoi(parts[1])
		if err != nil {
			return -1
		}
		switch command {
		case "forward":
			horizontal += units
			depth += aim * units
		case "down":
			aim += units
		case "up":
			aim -= units
		}
	}
	return horizontal * depth
}

func byteArrStringAndInt(bytes []byte) (string, int, error) {
	line := string(bytes[:])
	parts := strings.Split(line, " ")
	command := parts[0]
	units, err := strconv.Atoi(parts[1])
	return command, units, err
}
