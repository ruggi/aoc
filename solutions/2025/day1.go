package aoc2025

// https://adventofcode.com/2025/day/1

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ruggi/aoc/solutions"
)

func init() {
	solutions.Register(2025, 1, []solutions.SolutionFunc{
		day1Part1,
		day1Part2,
	})
}

func parseInstruction(instruction string) (int, error) {
	direction := instruction[0]
	distance, err := strconv.Atoi(instruction[1:])
	if err != nil {
		return 0, fmt.Errorf("parse distance: %w", err)
	}

	sign := 1
	if direction == 'L' { // left goes negative
		sign = -1
	}

	return distance * sign, nil
}

func day1Part1(input string) (string, error) {
	lockLength := 100
	dial := 50
	zeros := 0

	instructions := strings.Split(input, "\n")

	for ins := range instructions {
		step, err := parseInstruction(instructions[ins])
		if err != nil {
			return "", fmt.Errorf("parse instruction: %w", err)
		}

		unit := 1
		if step < 0 {
			unit = -1
		}
		for i := 0; i < int(math.Abs(float64(step))); i++ {
			dial += unit
			if dial < 0 {
				dial = lockLength - 1
			} else if dial >= lockLength {
				dial = 0
			}

		}
		if dial == 0 {
			zeros++
		}
	}

	return fmt.Sprintf("%d", zeros), nil
}

func day1Part2(input string) (string, error) {
	lockLength := 100
	dial := 50

	zeros := 0

	instructions := strings.Split(input, "\n")

	for ins := range instructions {
		step, err := parseInstruction(instructions[ins])
		if err != nil {
			return "", fmt.Errorf("parse instruction: %w", err)
		}

		unit := 1
		if step < 0 {
			unit = -1
		}
		for i := 0; i < int(math.Abs(float64(step))); i++ {
			dial += unit
			if dial < 0 {
				dial = lockLength - 1
			} else if dial >= lockLength {
				dial = 0
			}
			if dial == 0 {
				zeros++
			}
		}
	}

	return fmt.Sprintf("%d", zeros), nil
}
