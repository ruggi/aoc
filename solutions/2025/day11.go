package aoc2025

// https://adventofcode.com/2025/day/11

import (
	"fmt"
	"strings"

	"github.com/ruggi/aoc/solutions"
)

func init() {
	solutions.Register(2025, 11, []solutions.SolutionFunc{
		day11Part1,
		day11Part2,
	})
}

type device struct {
	name        string
	connections []string
}

func day11Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")

	deviceMap := map[string]device{}
	start := ""
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		deviceName := parts[0]
		connections := strings.Split(parts[1], " ")
		if deviceName == "you" {
			start = deviceName
		}
		deviceMap[deviceName] = device{
			name:        deviceName,
			connections: connections,
		}
	}

	paths := dfs(deviceMap, start, map[string]bool{})
	return fmt.Sprintf("%d", paths), nil
}

func dfs(deviceMap map[string]device, cur string, visited map[string]bool) int {
	if cur == "out" {
		return 1
	}
	if visited[cur] {
		return 0
	}
	visited[cur] = true
	defer delete(visited, cur)

	paths := 0
	device := deviceMap[cur]
	for _, connection := range device.connections {
		paths += dfs(deviceMap, connection, visited)
	}
	return paths
}

func memoKey(cur string, seenDac bool, seenFft bool) string {
	return fmt.Sprintf("%s,%t,%t", cur, seenDac, seenFft)
}

func dfs2(deviceMap map[string]device, cur string, visited map[string]bool, seenDac bool, seenFft bool, memo map[string]int) int {
	if cur == "out" {
		if seenDac && seenFft {
			return 1
		}
		return 0
	}
	if visited[cur] {
		return 0
	}

	seenDac = seenDac || cur == "dac"
	seenFft = seenFft || cur == "fft"

	key := memoKey(cur, seenDac, seenFft)
	if val, ok := memo[key]; ok {
		return val
	}

	visited[cur] = true
	defer delete(visited, cur)

	paths := 0
	device := deviceMap[cur]
	for _, connection := range device.connections {
		paths += dfs2(deviceMap, connection, visited, seenDac, seenFft, memo)
	}

	memo[key] = paths
	return paths
}

func day11Part2(input string) (string, error) {
	lines := strings.Split(input, "\n")

	deviceMap := map[string]device{}
	start := ""
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		deviceName := parts[0]
		connections := strings.Split(parts[1], " ")
		if deviceName == "svr" {
			start = deviceName
		}
		deviceMap[deviceName] = device{
			name:        deviceName,
			connections: connections,
		}
	}

	paths := dfs2(deviceMap, start, map[string]bool{}, false, false, make(map[string]int))
	return fmt.Sprintf("%d", paths), nil

}
