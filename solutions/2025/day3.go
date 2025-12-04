package aoc2025

// https://adventofcode.com/2025/day/3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ruggi/aoc/solutions"
)

func init() {
	solutions.Register(2025, 3, []solutions.SolutionFunc{
		day3Part1,
		day3Part2,
	})
}

func day3Part1(input string) (string, error) {
	banks := []string{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		banks = append(banks, line)
	}

	sum := 0
	for _, b := range banks {
		hl := byte('0')
		hlIndex := 0
		for i := 0; i < len(b)-1; i++ {
			s := b[i]
			if s > hl {
				hl = b[i]
				hlIndex = i
			}
		}

		hr := byte('0')
		for i := hlIndex + 1; i < len(b); i++ {
			if b[i] > hr {
				hr = b[i]
			}
		}

		num, err := strconv.Atoi(fmt.Sprintf("%s%s", string(hl), string(hr)))
		if err != nil {
			return "", fmt.Errorf("parse number: %w", err)
		}
		sum += num
	}

	return fmt.Sprintf("%d", sum), nil
}

func day3Part2(input string) (string, error) {
	batCount := 12

	banks := []string{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		banks = append(banks, line)
	}

	sum := int64(0)
	for _, b := range banks {
		values := []string{}

		hIndex := -1
		for j := range batCount {
			h := byte('0')
			for i := hIndex + 1; i < len(b)-(batCount-j-1); i++ {
				s := b[i]
				if s > h {
					h = b[i]
					hIndex = i
				}
			}
			values = append(values, string(h))
		}

		numString := strings.Join(values, "")
		num, err := strconv.ParseInt(numString, 10, 64)
		if err != nil {
			return "", fmt.Errorf("parse number: %w", err)
		}
		sum += num
	}

	return fmt.Sprintf("%d", sum), nil
}
