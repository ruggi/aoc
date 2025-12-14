package aoc2025

// https://adventofcode.com/2025/day/10

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ruggi/aoc/common"
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

func parseMachines(input string) []machine {
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
	return machines
}

func day10Part1(input string) (string, error) {
	machines := parseMachines(input)

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

type equation struct {
	buttons []int
	result  int
}

func (e equation) String() string {
	s := lo.Map(e.buttons, func(b int, _ int) string {
		return fmt.Sprintf("x%d", b)
	})
	return strings.Join(s, " + ") + " = " + fmt.Sprintf("%d", e.result)
}

func day10Part2(input string) (string, error) {
	// Note! I did not think of this solution myself initially. I wasted way too much time on day 10 itself on part 2 and I came back later on to complete it.
	// This Reddit post pointed me in the right direction: https://www.reddit.com/r/adventofcode/comments/1pl8nsa/comment/ntqt12a/?context=3
	// It's not super efficient compared to, say, Z3, but it's a fun solution to implement and I learned a bunch from it.
	// Merry christmas! 🎄

	machines := parseMachines(input)

	totSolutions := []int{}
	for _, m := range machines {
		// build equations
		equations := []equation{}
		for i, value := range m.joltage {
			e := equation{result: value}
			for j := range m.buttons {
				if lo.SomeBy(m.buttons[j], func(b int) bool { return b == i }) {
					e.buttons = append(e.buttons, j)
				}
			}
			equations = append(equations, e)
		}

		// build augmented matrix
		numButtons := len(m.buttons)
		matrix := make([][]int, len(equations))
		for i, eq := range equations {
			matrix[i] = make([]int, numButtons+1)
			for _, b := range eq.buttons {
				matrix[i][b] = 1
			}
			matrix[i][numButtons] = eq.result
		}

		// apply gauss-jordan to the equations matrix
		reducedMatrix, freeVars := common.GaussJordanElimination(matrix, numButtons)

		// find minimum solution by trying different free variable values
		minSum := -1
		maxFreeVarValue := 200

		// build a mapping of which pivot variables depend on which free variables
		type dependentVar struct {
			variable    int
			pivotCoeff  int
			constant    int
			freeVarCoef map[int]int
		}

		dependentVars := []dependentVar{}
		for _, row := range reducedMatrix {
			// find the pivot column
			pivotCol := -1
			for j := range numButtons {
				if row[j] != 0 {
					pivotCol = j
					break
				}
			}
			if pivotCol == -1 {
				continue
			}

			dv := dependentVar{
				variable:    pivotCol,
				pivotCoeff:  row[pivotCol],
				constant:    row[numButtons],
				freeVarCoef: make(map[int]int),
			}

			for j := range numButtons {
				if j != pivotCol && row[j] != 0 {
					dv.freeVarCoef[j] = row[j]
				}
			}

			dependentVars = append(dependentVars, dv)
		}

		// free variable combinations
		var searchFreeVars func(idx int, freeVarValues []int)
		searchFreeVars = func(idx int, freeVarValues []int) {
			if idx == len(freeVars) {
				// compute full solution
				solution := make([]int, numButtons)
				for i, fv := range freeVars {
					solution[fv] = freeVarValues[i]
				}

				// compute dependent variables
				valid := true
				for _, dv := range dependentVars {
					value := dv.constant
					for j, coef := range dv.freeVarCoef {
						value -= coef * solution[j]
					}

					// check if divisible
					if value%dv.pivotCoeff != 0 {
						valid = false
						break
					}
					value /= dv.pivotCoeff

					if value < 0 {
						valid = false
						break
					}
					solution[dv.variable] = value
				}

				if !valid {
					return
				}

				// validate solution against original equations
				for _, eq := range equations {
					sum := 0
					for _, b := range eq.buttons {
						sum += solution[b]
					}
					if sum != eq.result {
						return
					}
				}

				// update min
				currentSum := lo.Sum(solution)
				if minSum == -1 || currentSum < minSum {
					minSum = currentSum
				}
				return
			}

			// try different values for this free variable
			for val := 0; val <= maxFreeVarValue; val++ {
				freeVarValues[idx] = val
				searchFreeVars(idx+1, freeVarValues)
				if minSum != -1 && val >= minSum {
					break // stop here, no point trying larger values
				}
			}
		}

		if len(freeVars) == 0 {
			// no free variables - unique solution
			solution := make([]int, numButtons)
			valid := true
			for _, dv := range dependentVars {
				if dv.constant%dv.pivotCoeff != 0 {
					valid = false
					break
				}
				value := dv.constant / dv.pivotCoeff
				if value < 0 {
					valid = false
					break
				}
				solution[dv.variable] = value
			}
			if valid {
				minSum = lo.Sum(solution)
			}
		} else {
			searchFreeVars(0, make([]int, len(freeVars)))
		}

		totSolutions = append(totSolutions, minSum)
	}

	return fmt.Sprintf("%d", lo.Sum(totSolutions)), nil
}
