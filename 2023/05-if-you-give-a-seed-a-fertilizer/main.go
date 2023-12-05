package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Map struct {
	d int64 // destination range start
	s int64 // source range start
	l int64 // range length
}

type Stage int64

type Maps map[Stage][]Map

type Seeds []int64

func mustint64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func readInput(file string) (seeds Seeds, maps Maps) {

	seeds, maps = Seeds{}, Maps{}
	input, _ := os.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	var re = regexp.MustCompile(`(?m)(\d+)`)

	seedLines := re.FindAllString(lines[0], -1)
	for _, s := range seedLines {
		seeds = append(seeds, mustint64(s))
	}

	var currentStage Stage
	for _, ms := range lines[2:] {
		mv := re.FindAllString(ms, -1)
		if len(mv) == 3 {
			m := Map{d: mustint64(mv[0]), s: mustint64(mv[1]), l: mustint64(mv[2])}
			maps[currentStage] = append(maps[currentStage], m)
		}
		if ms == "" {
			currentStage++
		}
	}

	return
}

func (m Map) locate(seed int64) int64 {

	if seed >= m.s && seed < m.s+m.l {
		seed += (m.d - m.s)
	}

	return seed
}

func Solve(file string) int64 {

	seeds, maps := readInput(file)

	// fmt.Println(seeds)

	for i := Stage(0); int(i) < len(maps); i++ {
		for is, s := range seeds {
			for _, m := range maps[i] {
				if seeds[is] != m.locate(s) {
					seeds[is] = m.locate(s)
					break
				}
				// fmt.Println("stage", i, "seeds", seeds)
			}
		}
	}

	// fmt.Println(seeds)
	return slices.Min(seeds)
}

func main() {
	location := Solve("input1.txt")
	fmt.Println(location) // 484023871

}
