package solutions

import (
	"fmt"
	"time"
)

type SolutionFunc func(input string) (string, error)

type Solution struct {
	Year  int
	Day   int
	Parts []SolutionFunc
}

var solutions = []Solution{}

func Register(year int, day int, parts []SolutionFunc) {
	solutions = append(solutions, Solution{
		Year:  year,
		Day:   day,
		Parts: parts,
	})
}

func Run(year int, day int, part int, input string) error {
	for _, s := range solutions {
		if s.Year == year && s.Day == day {
			for i, p := range s.Parts {
				startTime := time.Now()
				if part != -1 && part != i+1 {
					continue
				}
				fmt.Println("----------------------------------------")
				fmt.Println("year", year, "day", day, "part", i+1)
				output, err := p(input)
				if err != nil {
					return fmt.Errorf("part %d: %w", i+1, err)
				}
				fmt.Println("=>", output)
				elapsed := time.Since(startTime)
				fmt.Printf("(duration: %dms)\n", elapsed.Milliseconds())
			}
			return nil
		}
	}
	return fmt.Errorf("solution not found")
}
