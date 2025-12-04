package aoc2025

// https://adventofcode.com/2025/day/4

import (
	"fmt"
	"strings"

	"github.com/ruggi/aoc/solutions"
)

func init() {
	solutions.Register(2025, 4, []solutions.SolutionFunc{
		day4Part1,
		day4Part2,
	})
}

func printGrid(grid [][]string) {
	for y := range grid {
		for x := range grid[y] {
			fmt.Print(grid[y][x])
		}
		fmt.Println()
	}
}

func day4Part1(input string) (string, error) {
	grid := parseGrid(input)

	// printGrid(grid)

	accessible := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "." {
				continue
			}
			neighbors := neighbors(grid, x, y)
			count := strings.Count(strings.Join(neighbors, ""), "@")
			// fmt.Println(x, y, neighbors, count)
			canAccess := count < 4
			if canAccess {
				accessible++
			}
		}
	}

	return fmt.Sprintf("%d", accessible), nil
}

func neighbors(grid [][]string, x int, y int) []string {
	getCell := func(x int, y int) string {
		if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
			return ""
		}
		return grid[y][x]
	}

	return []string{
		getCell(x-1, y-1), // top left
		getCell(x, y-1),   // top
		getCell(x+1, y-1), // top right
		getCell(x-1, y),   // left
		getCell(x+1, y),   // right
		getCell(x-1, y+1), // bottom left
		getCell(x, y+1),   // bottom
		getCell(x+1, y+1), // bottom right
	}
}

func parseGrid(input string) [][]string {
	grid := [][]string{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}

func day4Part2(input string) (string, error) {
	grid := parseGrid(input)

	removed := 0
	for {
		hadAccessiblePaper := false
		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] != "@" {
					continue
				}
				neighbors := neighbors(grid, x, y)
				count := strings.Count(strings.Join(neighbors, ""), "@")
				canAccess := count < 4
				if canAccess {
					removed++
					grid[y][x] = "."
					hadAccessiblePaper = true
				}
			}
		}
		if !hadAccessiblePaper {
			break
		}
		// printGrid(grid)
		// fmt.Println()
	}

	return fmt.Sprintf("%d", removed), nil
}
