package aoc2025

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ruggi/aoc/solutions"
)

// https://adventofcode.com/2025/day/5

func init() {
	solutions.Register(2025, 5, []solutions.SolutionFunc{
		day5Part1,
		day5Part2,
	})
}

type ingRange struct {
	start int64
	end   int64
}

type ranges []ingRange

func (r ranges) Len() int {
	return len(r)
}

func (r ranges) Less(i, j int) bool {
	return r[i].start < r[j].start
}

func (r ranges) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func isFresh(ranges []ingRange, id int64) bool {
	for _, r := range ranges {
		if id >= r.start && id <= r.end {
			return true
		}
	}
	return false
}

func parseRanges(strRanges string) (ranges, error) {
	ranges := ranges{}
	lines := strings.Split(strRanges, "\n")
	for _, line := range lines {
		start, end, _ := strings.Cut(line, "-")
		startInt, err := strconv.ParseInt(start, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse start: %w", err)
		}
		endInt, err := strconv.ParseInt(end, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse end: %w", err)
		}
		ranges = append(ranges, ingRange{start: startInt, end: endInt})
	}
	return ranges, nil
}

func day5Part1(input string) (string, error) {
	parts := strings.Split(input, "\n\n")

	ranges, err := parseRanges(parts[0])
	if err != nil {
		return "", fmt.Errorf("parse ranges: %w", err)
	}
	ingredients := strings.Split(parts[1], "\n")

	count := 0
	for _, ing := range ingredients {
		ingInt, err := strconv.ParseInt(ing, 10, 64)
		if err != nil {
			return "", fmt.Errorf("parse ingredient: %w", err)
		}
		if isFresh(ranges, ingInt) {
			count++
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func day5Part2(input string) (string, error) {
	parts := strings.Split(input, "\n\n")
	r, err := parseRanges(parts[0])
	if err != nil {
		return "", fmt.Errorf("parse ranges: %w", err)
	}

	// sort the ranges
	sort.Sort(r)

	// merge overlapping ranges
	merged := ranges{}
	if len(r) > 0 {
		merged = append(merged, r[0])
		for i := 1; i < len(r); i++ {
			last := &merged[len(merged)-1]
			if r[i].start <= last.end {
				// ranges overlap, merge them
				if r[i].end > last.end {
					last.end = r[i].end
				}
			} else {
				// no overlap, add as new range
				merged = append(merged, r[i])
			}
		}
	}

	count := int64(0)
	for _, r := range merged {
		count += r.end - r.start + 1
	}

	return fmt.Sprintf("%d", count), nil
}
