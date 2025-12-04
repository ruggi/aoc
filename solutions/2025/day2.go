package aoc2025

// https://adventofcode.com/2025/day/2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ruggi/aoc/solutions"
)

func init() {
	solutions.Register(2025, 2, []solutions.SolutionFunc{
		day2Part1,
		day2Part2,
	})
}

func isValidIdPart1(id string) bool {
	if id[0] == '0' {
		return true
	}
	if len(id) == 1 || len(id)%2 != 0 {
		return true
	}

	half := len(id) / 2
	left := id[:half]
	right := id[half:]
	if left != right {
		return true
	}

	count := strings.Count(id, left)
	if count == 2 { // only twice
		return false
	}

	return true
}

func day2Part1(input string) (string, error) {
	ranges := strings.Split(input, ",")
	invalidIDs := []int64{}
	for _, r := range ranges {
		startStr, endStr, _ := strings.Cut(r, "-")
		start, err := strconv.ParseInt(startStr, 10, 64)
		if err != nil {
			return "", fmt.Errorf("parse start: %w", err)
		}
		end, err := strconv.ParseInt(endStr, 10, 64)
		if err != nil {
			return "", fmt.Errorf("parse end: %w", err)
		}

		for i := start; i <= end; i++ {
			if !isValidIdPart1(strconv.FormatInt(i, 10)) {
				invalidIDs = append(invalidIDs, i)
			}
		}
	}

	sum := int64(0)
	for _, id := range invalidIDs {
		sum += id
	}

	return fmt.Sprintf("%d", sum), nil
}

func isValidIdPart2(id string) bool {
	if len(id) == 1 {
		return true
	}

	for pos := 1; pos <= len(id)/2; pos++ {
		if id[0] == '0' {
			continue
		}
		if len(id)%pos != 0 {
			continue
		}

		repeats := len(id) / pos
		if repeats < 2 {
			continue
		}

		matches := true
		for i := 1; i < repeats; i++ {
			start := i * pos
			if id[start:start+pos] != id[:pos] {
				matches = false
				break
			}
		}

		if matches {
			count := strings.Count(id, id[:pos])
			if count >= 2 { // at least twice
				return false
			}
		}
	}

	return true
}

func day2Part2(input string) (string, error) {
	ranges := strings.Split(input, ",")
	invalidIDs := []int64{}
	for _, r := range ranges {
		startStr, endStr, _ := strings.Cut(r, "-")
		start, err := strconv.ParseInt(startStr, 10, 64)
		if err != nil {
			return "", fmt.Errorf("parse start: %w", err)
		}
		end, err := strconv.ParseInt(endStr, 10, 64)
		if err != nil {
			return "", fmt.Errorf("parse end: %w", err)
		}

		for i := start; i <= end; i++ {
			if !isValidIdPart2(strconv.FormatInt(i, 10)) {
				invalidIDs = append(invalidIDs, i)
			}
		}
	}

	sum := int64(0)
	for _, id := range invalidIDs {
		sum += id
	}
	return fmt.Sprintf("%d", sum), nil
}
