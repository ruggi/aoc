package aoc2025

// https://adventofcode.com/2025/day/9

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ruggi/aoc/solutions"
	"github.com/samber/lo"
)

func init() {
	solutions.Register(2025, 9, []solutions.SolutionFunc{
		day9Part1,
		day9Part2,
	})
}

func day9Part1(input string) (string, error) {
	positions := []position{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		positions = append(positions, position{x: x, y: y})
	}

	area := 0
	for _, a := range positions {
		for _, b := range positions {
			topLeft := position{x: a.x, y: a.y}
			bottomRight := position{x: b.x, y: b.y}

			if a.x > b.x {
				topLeft.x = b.x
				bottomRight.x = a.x
			}
			if a.y > b.y {
				topLeft.y = b.y
				bottomRight.y = a.y
			}

			width := bottomRight.x - topLeft.x + 1
			height := bottomRight.y - topLeft.y + 1
			curArea := width * height
			if curArea > area {
				area = curArea
			}
		}
	}

	return fmt.Sprintf("%d", area), nil

}

func day9Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")

	redPositions := []position{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		redPositions = append(redPositions, position{x: x, y: y})
	}

	// build a map of boundary tiles (red or green) without a full grid
	boundary := make(map[position]bool)

	// mark red tiles
	for _, pos := range redPositions {
		boundary[pos] = true
	}

	// mark green tiles connecting consecutive red tiles
	for i := range redPositions {
		from := redPositions[i]
		to := redPositions[(i+1)%len(redPositions)]

		if from.y == to.y {
			// horizontal line
			startX := from.x
			endX := to.x
			if startX > endX {
				startX, endX = endX, startX
			}
			for x := startX; x <= endX; x++ {
				boundary[position{x: x, y: from.y}] = true
			}
		} else if from.x == to.x {
			// vertical line
			startY := from.y
			endY := to.y
			if startY > endY {
				startY, endY = endY, startY
			}
			for y := startY; y <= endY; y++ {
				boundary[position{x: from.x, y: y}] = true
			}
		}
	}

	// compute the valid x ranges
	xRangesByY := make(map[int][2]int) // y -> [minX, maxX]
	yValues := []int{}
	for pos := range boundary {
		yValues = append(yValues, pos.y)
		r, ok := xRangesByY[pos.y]
		if ok {
			if pos.x < r[0] {
				r[0] = pos.x
			}
			if pos.x > r[1] {
				r[1] = pos.x
			}
			xRangesByY[pos.y] = r
		} else {
			xRangesByY[pos.y] = [2]int{pos.x, pos.x}
		}
	}

	yValues = lo.Uniq(yValues)
	sort.Ints(yValues)

	// cache y indexes
	yIndexes := make(map[int]int)
	for i, y := range yValues {
		yIndexes[y] = i
	}

	type rect struct {
		a, b    position
		maxArea int
	}

	rects := make([]rect, 0)
	for _, a := range redPositions {
		for _, b := range redPositions {
			// quick Y continuity check before doing anything else
			if !yRangeValid(yIndexes, a.y, b.y) {
				continue
			}

			x1, x2 := a.x, b.x
			y1, y2 := a.y, b.y
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			maxArea := (x2 - x1 + 1) * (y2 - y1 + 1)
			rects = append(rects, rect{a: a, b: b, maxArea: maxArea})
		}
	}

	sort.Slice(rects, func(i, j int) bool {
		return rects[i].maxArea > rects[j].maxArea
	})

	area := 0
	for _, r := range rects {
		// make it faster by skipping smaller areas
		if r.maxArea <= area {
			break
		}

		if !xRangeValid(yIndexes, yValues, xRangesByY, r.a, r.b) {
			continue
		}

		if r.maxArea > area {
			area = r.maxArea
		}
	}

	return fmt.Sprintf("%d", area), nil
}

func yRangeValid(yIndexes map[int]int, ay, by int) bool {
	if ay > by {
		ay, by = by, ay
	}
	ai, ok := yIndexes[ay]
	if !ok {
		return false
	}
	bi, ok := yIndexes[by]
	if !ok {
		return false
	}
	return bi-ai == by-ay
}

func xRangeValid(yIndexes map[int]int, yValues []int, xRangesByY map[int][2]int, a, b position) bool {
	if a.x > b.x {
		a.x, b.x = b.x, a.x
	}
	if a.y > b.y {
		a.y, b.y = b.y, a.y
	}
	ai := yIndexes[a.y]
	bi := yIndexes[b.y]

	for i := ai; i <= bi; i++ {
		y := yValues[i]
		r := xRangesByY[y]
		if a.x < r[0] || b.x > r[1] {
			return false
		}
	}
	return true
}
