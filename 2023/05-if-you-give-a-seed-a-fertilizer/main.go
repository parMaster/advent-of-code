package main

import (
	"fmt"
	"math"
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

type Stage int

type Maps map[int][]Map

type Seeds []int64

func mustInt64(s string) int64 {
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
		seeds = append(seeds, mustInt64(s))
	}

	var currentStage int
	for _, ms := range lines[2:] {
		mv := re.FindAllString(ms, -1)
		if len(mv) == 3 {
			m := Map{d: mustInt64(mv[0]), s: mustInt64(mv[1]), l: mustInt64(mv[2])}
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

func PartOne(file string) int64 {

	seeds, maps := readInput(file)

	for i := 0; i < len(maps); i++ {
		for is, s := range seeds {
			for _, m := range maps[i] {
				if seeds[is] != m.locate(s) {
					seeds[is] = m.locate(s)
					break
				}
			}
		}
	}

	return slices.Min(seeds)
}

func PartTwo(file string) int64 {
	seeds, maps := readInput(file)

	var minlocation int64 = math.MaxInt64
	var seed int64
	var loc int64
	for i := 0; i < len(seeds)/2; i++ {
		start, length := seeds[i*2], seeds[i*2+1]
		// fmt.Println("start", start, " length = ", length)
		for s := start; s < start+length; s++ {
			seed = s
			if seed%10000000 == 0 {
				fmt.Print(".")
			}

			for j := 0; j < len(maps); j++ {
				for mi := range maps[j] {
					loc = maps[j][mi].locate(seed)
					if seed != loc {
						seed = loc
						break
					}
				}
			}
			minlocation = min(minlocation, seed)
		}
		// fmt.Println("minlocation = ", minlocation)
	}

	return minlocation
}

func main() {
	fmt.Println("Day 5: If You Give A Seed A Fertilizer \n\tPart One:", PartOne("../aoc-inputs/2023/05/input1.txt")) // 484023871
	fmt.Println("\tPart Two:", PartTwo("../aoc-inputs/2023/05/input1.txt"))                                          // 46294175
}
