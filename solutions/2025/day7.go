package aoc2025

// https://adventofcode.com/2025/day/7

import (
	"fmt"
	"strings"

	"github.com/ruggi/aoc/solutions"
	"github.com/samber/lo"
)

func init() {
	solutions.Register(2025, 7, []solutions.SolutionFunc{
		day7Part1,
		day7Part2,
	})
}

func day7Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")

	splits := 0
	currentColumns := []int{strings.Index(lines[0], "S")}
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		// go through clean passes first
		for _, col := range currentColumns {
			if line[col] == '.' {
				line = line[:col] + "|" + line[col+1:]
			}
		}
		// then go through the splitters
		for _, col := range currentColumns {
			if line[col] == '^' {
				splits++
				// replace left and right of the column with |
				line = line[:col-1] + "|" + line[col:]
				line = line[:col+1] + "|" + line[col+2:]
				// add the new columns to the list
				currentColumns = append(currentColumns, col-1, col+1)
				currentColumns = lo.Filter(currentColumns, func(c int, _ int) bool {
					return c != col
				})
			}
			lines[i] = line
		}
		currentColumns = lo.Uniq(currentColumns)
	}

	return fmt.Sprintf("%d", splits), nil
}

func day7Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")

	col := strings.Index(lines[0], "S")
	memo := make(map[string]int64)
	splits := traverse(append([]string{}, lines[1:]...), col, 1, memo)

	return fmt.Sprintf("%d", splits), nil
}

func traverse(lines []string, col int, splits int64, memo map[string]int64) int64 {
	if len(lines) == 0 {
		return splits
	}
	if col < 0 || col >= len(lines[0]) {
		return 0
	}

	memoKey := fmt.Sprintf("%d,%d", len(lines), col)
	if val, ok := memo[memoKey]; ok {
		return val
	}

	result := int64(0)
	next := lines[1:]
	if lines[0][col] == '^' {
		result = traverse(next, col-1, splits, memo) + traverse(next, col+1, splits, memo)
	} else {
		result = traverse(next, col, splits, memo)
	}

	memo[memoKey] = result
	return result
}
