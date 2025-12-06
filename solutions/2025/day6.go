package aoc2025

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ruggi/aoc/solutions"
)

// https://adventofcode.com/2025/day/6

func init() {
	solutions.Register(2025, 6, []solutions.SolutionFunc{
		day6Part1,
		day6Part2,
	})
}

type problem struct {
	numbers    []int64
	numbersStr []string
	op         string
	result     int64
	width      int
}

func day6Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")

	columns := len(strings.Fields(lines[0]))
	problems := make([]problem, columns)

	operatorsLine := lines[len(lines)-1]
	numbersLines := lines[:len(lines)-1]

	// place the operators in the problems
	for i, op := range strings.Fields(operatorsLine) {
		problems[i].op = op
	}
	// place the numbers in the problems
	for _, numbersLine := range numbersLines {
		numbers := strings.Fields(numbersLine)
		for i, num := range numbers {
			numInt, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				return "", fmt.Errorf("parse number: %w", err)
			}
			problems[i].numbers = append(problems[i].numbers, numInt)
		}
	}
	// calculate the results
	for i, problem := range problems {
		result := problem.numbers[0]
		for i := 1; i < len(problem.numbers); i++ {
			switch problem.op {
			case "+":
				result += problem.numbers[i]
			case "*":
				result *= problem.numbers[i]
			}
		}
		problems[i].result = result
	}

	sum := int64(0)
	for _, problem := range problems {
		sum += problem.result
	}

	return fmt.Sprintf("%d", sum), nil
}

func day6Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")

	columns := len(strings.Fields(lines[0]))
	problems := make([]problem, columns)

	operatorsLine := lines[len(lines)-1]
	numbersLines := lines[:len(lines)-1]

	// place the operators in the problems
	re := regexp.MustCompile(`([\*\+])(\s+)`)
	matches := re.FindAllStringSubmatch(operatorsLine, -1)
	for i, match := range matches {
		problems[i].op = match[1]
		problems[i].width = len(match[2])
		if i == len(matches)-1 {
			problems[i].width += 1
		}
	}

	// place the numbers (as strings) in the problems
	for _, line := range numbersLines {
		offset := 0
		for i, p := range problems {
			problems[i].numbersStr = append(problems[i].numbersStr, line[offset:offset+p.width])
			offset += p.width + 1
		}
		offset = 0
	}

	// rearrange digits so cephalopod are happy lol
	for problemIndex, problem := range problems {
		// create a matrix with the numbers (rows x cols)
		matrix := make([][]string, len(problem.numbersStr))
		for i := range matrix {
			matrix[i] = make([]string, problem.width)
		}
		for i := range matrix {
			for j := range matrix[i] {
				matrix[i][j] = string(problem.numbersStr[i][j])
			}
		}

		// rotate: read columns as rows
		rotated := []string{}
		for j := range problem.width {
			n := ""
			for i := range matrix {
				n += matrix[i][j]
			}
			rotated = append(rotated, n)
		}

		numbers := []int64{}
		for _, n := range rotated {
			nn, err := strconv.ParseInt(strings.TrimSpace(n), 10, 64)
			if err != nil {
				return "", fmt.Errorf("parse number: %w", err)
			}
			numbers = append(numbers, nn)
		}
		problems[problemIndex].numbers = numbers
	}

	// calculate the results
	for i, problem := range problems {
		result := problem.numbers[0]
		for i := 1; i < len(problem.numbers); i++ {
			switch problem.op {
			case "+":
				result += problem.numbers[i]
			case "*":
				result *= problem.numbers[i]
			}
		}
		problems[i].result = result
	}

	sum := int64(0)
	for _, problem := range problems {
		sum += problem.result
	}

	return fmt.Sprintf("%d", sum), nil
}
