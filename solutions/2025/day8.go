package aoc2025

// https://adventofcode.com/2025/day/8

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/ruggi/aoc/solutions"
	"github.com/samber/lo"
)

func init() {
	solutions.Register(2025, 8, []solutions.SolutionFunc{
		day8Part1,
		day8Part2,
	})
}

type junctionBox string

func (j junctionBox) position() position {
	parts := strings.Split(string(j), ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])
	return position{x: x, y: y, z: z}
}

type junctionBoxes []junctionBox

func (b junctionBoxes) Len() int {
	return len(b)
}

func (b junctionBoxes) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b junctionBoxes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

type position struct {
	x int
	y int
	z int
}

func euclideanDistance3d(a position, b position) float64 {
	return math.Sqrt(float64(a.x-b.x)*float64(a.x-b.x) + float64(a.y-b.y)*float64(a.y-b.y) + float64(a.z-b.z)*float64(a.z-b.z))
}

type circuit struct {
	boxes []junctionBox
}

func (c circuit) contains(b junctionBox) bool {
	return slices.Contains(c.boxes, b)
}

type circuits []circuit

func (c circuits) Len() int {
	return len(c)
}

func (c circuits) Less(i, j int) bool {
	return len(c[i].boxes) > len(c[j].boxes) // bigger circuits first
}

func (c circuits) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type distance struct {
	key    string
	length float64
}

func (d distance) boxes() (junctionBox, junctionBox) {
	parts := strings.Split(d.key, " ")
	return junctionBox(parts[0]), junctionBox(parts[1])
}

type distances []distance

func (d distances) Len() int {
	return len(d)
}

func (d distances) Less(i, j int) bool {
	return d[i].length < d[j].length
}

func (d distances) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func day8Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")

	boxes := []junctionBox{}
	for _, line := range lines {
		boxes = append(boxes, junctionBox(line))
	}

	// create a list of distances, sorted by length
	distMap := map[string]float64{}
	for i := range boxes {
		for j := range boxes {
			if i == j {
				continue
			}
			keys := junctionBoxes{boxes[i], boxes[j]}
			sort.Sort(keys)
			key := string(keys[0]) + " " + string(keys[1])
			distMap[key] = euclideanDistance3d(boxes[i].position(), boxes[j].position())
		}
	}
	distances := distances{}
	for k, v := range distMap {
		distances = append(distances, distance{key: k, length: v})
	}
	sort.Sort(distances)

	circuits := circuits{}
	for _, box := range boxes {
		circuits = append(circuits, circuit{boxes: []junctionBox{box}})
	}

	limit := 1000 // or 10 for the test input
	for _, d := range distances[:limit] {
		a, b := d.boxes()

		// find which circuits contain a and b
		circuitA := -1
		circuitB := -1
		for i, c := range circuits {
			if c.contains(a) {
				circuitA = i
			}
			if c.contains(b) {
				circuitB = i
			}
		}

		// both are in the same circuit, nothing to do
		if circuitA == circuitB && circuitA != -1 {
			continue
		}

		if circuitA != -1 && circuitB != -1 {
			// both are in different circuits, merge them
			circuits[circuitA].boxes = append(circuits[circuitA].boxes, circuits[circuitB].boxes...)
			circuits = append(circuits[:circuitB], circuits[circuitB+1:]...)
		} else if circuitA != -1 {
			// only a is in a circuit, add b to it
			circuits[circuitA].boxes = append(circuits[circuitA].boxes, b)
		} else if circuitB != -1 {
			// only b is in a circuit, add a to it
			circuits[circuitB].boxes = append(circuits[circuitB].boxes, a)
		}
	}

	sort.Sort(circuits)

	topThree := circuits[:3]
	result := lo.Reduce(topThree, func(acc int, c circuit, _ int) int {
		return acc * len(c.boxes)
	}, 1)
	return fmt.Sprintf("%d", result), nil
}

func day8Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")

	boxes := []junctionBox{}
	for _, line := range lines {
		boxes = append(boxes, junctionBox(line))
	}

	// create a list of distances, sorted by length
	distMap := map[string]float64{}
	for i := range boxes {
		for j := range boxes {
			if i == j {
				continue
			}
			keys := junctionBoxes{boxes[i], boxes[j]}
			sort.Sort(keys)
			key := string(keys[0]) + " " + string(keys[1])
			distMap[key] = euclideanDistance3d(boxes[i].position(), boxes[j].position())
		}
	}
	distances := distances{}
	for k, v := range distMap {
		distances = append(distances, distance{key: k, length: v})
	}
	sort.Sort(distances)

	circuits := circuits{}
	for _, box := range boxes {
		circuits = append(circuits, circuit{boxes: []junctionBox{box}})
	}

	// the junction boxes that cause a single circuit
	tiebreakers := []junctionBox{}

	for _, d := range distances {
		a, b := d.boxes()

		// find which circuits contain a and b
		circuitA := -1
		circuitB := -1
		for i, c := range circuits {
			if c.contains(a) {
				circuitA = i
			}
			if c.contains(b) {
				circuitB = i
			}
		}

		// both are in the same circuit, nothing to do
		if circuitA == circuitB && circuitA != -1 {
			continue
		}

		if circuitA != -1 && circuitB != -1 {
			// both are in different circuits, merge them
			circuits[circuitA].boxes = append(circuits[circuitA].boxes, circuits[circuitB].boxes...)
			circuits = append(circuits[:circuitB], circuits[circuitB+1:]...)
		} else if circuitA != -1 {
			// only a is in a circuit, add b to it
			circuits[circuitA].boxes = append(circuits[circuitA].boxes, b)
		} else if circuitB != -1 {
			// only b is in a circuit, add a to it
			circuits[circuitB].boxes = append(circuits[circuitB].boxes, a)
		}

		if len(circuits) == 1 {
			tiebreakers = append(tiebreakers, a, b)
			break
		}
	}

	wallDist := lo.Reduce(tiebreakers, func(acc int, b junctionBox, _ int) int {
		return b.position().x * acc
	}, 1)

	return fmt.Sprintf("%d", wallDist), nil
}
