package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
	"time"
)

func solve(filename string) (int, int) {
	file, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	g := map[string][]string{}
	smallCaves := map[string]bool{}
	for _, line := range lines {
		caves := strings.Split(line, "-")
		g[caves[0]] = append(g[caves[0]], caves[1])
		g[caves[1]] = append(g[caves[1]], caves[0])
		if byte(caves[0][0]) >= 97 && byte(caves[0][0]) <= 122 {
			smallCaves[caves[0]] = true
		}
		if byte(caves[1][0]) >= 97 && byte(caves[1][0]) <= 122 {
			smallCaves[caves[1]] = true
		}
	}

	var p1 func(path []string, visited map[string]int) (paths [][]string)
	p1 = func(path []string, visited map[string]int) (paths [][]string) {
		for _, next := range g[path[len(path)-1]] {
			if next == "start" {
				continue
			}
			if next == "end" {
				paths = append(paths, append(path, "end"))
				continue
			}

			if slices.Contains(path, next) && smallCaves[next] && visited[next] >= 1 {
				continue
			}

			vNew := maps.Clone(visited)
			vNew[next]++

			paths = append(paths, p1(append(path, next), vNew)...)
		}
		return paths
	}

	var p2 func(path []string, visited map[string]int) (paths [][]string)
	p2 = func(path []string, visited map[string]int) (paths [][]string) {
		for _, next := range g[path[len(path)-1]] {
			if next == "start" {
				continue
			}
			if next == "end" {
				paths = append(paths, append(path, "end"))
				continue
			}

			doubleVisited := false
			for cave, v := range visited {
				if v > 1 && smallCaves[cave] {
					doubleVisited = true
					break
				}
			}

			if slices.Contains(path, next) && smallCaves[next] && (visited[next] >= 2 || doubleVisited) {
				continue
			}

			vNew := maps.Clone(visited)
			vNew[next]++

			paths = append(paths, p2(append(path, next), vNew)...)
		}
		return paths
	}

	return len(p1([]string{"start"}, map[string]int{})), len(p2([]string{"start"}, map[string]int{}))
}

func main() {
	start := time.Now()
	fmt.Println("Day 12: Passage Pathing")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 3421
	fmt.Println("\tPart Two:", p2) // 84870
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
