package aoc2025

// https://adventofcode.com/2025/day/10

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ruggi/aoc/solutions"
	"github.com/samber/lo"
)

func init() {
	solutions.Register(2025, 10, []solutions.SolutionFunc{
		day10Part1,
		day10Part2,
	})
}

type machine struct {
	lights  []bool
	buttons [][]int
	joltage []int
}

func parseLights(s string) []bool {
	lights := []bool{}
	s = strings.Trim(s, "[]")
	fields := strings.Split(s, "")
	for _, field := range fields {
		if field == "." {
			lights = append(lights, false)
		} else {
			lights = append(lights, true)
		}
	}
	return lights
}

func parseJoltage(s string) []int {
	joltage := []int{}
	s = strings.Trim(s, "{}")
	fields := strings.Split(s, ",")
	for _, field := range fields {
		n, _ := strconv.Atoi(field)
		joltage = append(joltage, n)
	}
	return joltage
}

func parseButtons(s string) [][]int {
	buttons := [][]int{}
	fields := strings.Fields(s)
	for _, field := range fields {
		field = strings.Trim(field, "()")
		tokens := strings.Split(field, ",")
		sequence := []int{}
		for _, token := range tokens {
			n, _ := strconv.Atoi(token)
			sequence = append(sequence, n)
		}
		buttons = append(buttons, sequence)
	}
	return buttons
}

func day10Part1(input string) (string, error) {
	machines := []machine{}
	lines := strings.Split(input, "\n")
	reMachine := regexp.MustCompile(`\[.*\]`)
	reJoltage := regexp.MustCompile(`\{.*\}`)

	for _, line := range lines {
		strLights := reMachine.FindStringSubmatch(line)[0]
		strJoltage := reJoltage.FindStringSubmatch(line)[0]
		strButtons := reJoltage.ReplaceAllString(reMachine.ReplaceAllString(line, ""), "")

		machines = append(machines, machine{
			lights:  parseLights(strLights),
			buttons: parseButtons(strButtons),
			joltage: parseJoltage(strJoltage),
		})

	}

	solutions := []int{}
	for _, m := range machines {
		for i := 1; i <= 100; i++ { // try up to 100 buttons (IT SUCKS BUT IT WORKS! :D)
			permutations := permWithRep(m.buttons, i)
			tryPermutation := func() bool {
				for _, p := range permutations {
					ok := tryButtons(m.lights, p)
					if ok {
						solutions = append(solutions, i)
						return true
					}
				}
				return false
			}
			found := tryPermutation()
			if found {
				break
			}
		}
	}

	return fmt.Sprintf("%d", lo.Sum(solutions)), nil
}

func tryButtons(lightsTarget []bool, sequence [][]int) bool {
	lights := make([]bool, len(lightsTarget))
	for _, seq := range sequence {
		for _, button := range seq {
			lights[button] = !lights[button]
		}
	}
	for i := range lights {
		if lights[i] != lightsTarget[i] {
			return false
		}
	}
	return true
}

func permWithRep[T any](arr []T, length int) [][]T {
	if length == 0 {
		return [][]T{{}}
	}

	if len(arr) == 0 {
		return [][]T{}
	}

	var result [][]T

	subPerms := permWithRep(arr, length-1)

	for _, subPerm := range subPerms {
		for _, elem := range arr {
			newPerm := make([]T, 0, length)
			newPerm = append(newPerm, elem)
			newPerm = append(newPerm, subPerm...)
			result = append(result, newPerm)
		}
	}

	return result
}

func day10Part2(input string) (string, error) {
	return "", nil
}
