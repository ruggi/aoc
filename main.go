package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/ruggi/aoc/solutions"
	"github.com/urfave/cli/v3"

	// 2025
	_ "github.com/ruggi/aoc/solutions/2025"
)

var config struct {
	year  int
	day   int
	input string
	part  int
}

func main() {
	app := &cli.Command{
		Name:  "aoc",
		Usage: "advent of code solutions",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "year",
				Aliases:     []string{"y"},
				Required:    true,
				Destination: &config.year,
			},
			&cli.IntFlag{
				Name:        "day",
				Aliases:     []string{"d"},
				Required:    true,
				Destination: &config.day,
			},
			&cli.StringFlag{
				Name:        "input",
				Aliases:     []string{"i"},
				Required:    true,
				Destination: &config.input,
			},
			&cli.IntFlag{
				Name:        "part",
				Aliases:     []string{"p"},
				Destination: &config.part,
				Value:       -1,
			},
		},
		Action: run,
	}
	err := app.Run(context.Background(), os.Args)
	if err != nil {
		panic(err)
	}
}

func run(ctx context.Context, c *cli.Command) error {
	inputData, err := os.ReadFile(config.input)
	if err != nil {
		return fmt.Errorf("read input: %w", err)
	}
	input := strings.TrimSpace(string(inputData))

	err = solutions.Run(config.year, config.day, config.part, input)
	if err != nil {
		return fmt.Errorf("run solution: %w", err)
	}

	return nil
}
