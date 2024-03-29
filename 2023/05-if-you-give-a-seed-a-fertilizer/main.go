package main

/* Running on 4 cores:
date && go build -o torture && ./torture --bruteforce 4 && date
*/

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/workerpool"
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

func solve(start int64, length int64, maps Maps, results chan<- int64) {
	var loc int64
	var minlocation int64 = math.MaxInt64

	startTime := time.Now()
	fmt.Println("Planting", length, " \tseeds from", start)
	for s := start; s < start+length; s++ {
		seed := s
		// if seed%10000000 == 0 {
		// 	fmt.Print(".")
		// }

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
	rate := int64(float64(length) / time.Since(startTime).Seconds())
	fmt.Println("\u001b[1mPlanted ", length, "\tseeds from", start, "\tResult =", minlocation, "| Done in", time.Since(startTime).Truncate(time.Second), "-", rate, "seeds/sec\033[0m")
	results <- minlocation
}

func PartTwo(file string, cores int) int64 {
	seeds, maps := readInput(file)
	N := len(seeds) / 2
	results := make(chan int64, N)

	fmt.Print("Running ", N, " goroutines on ", cores, " cores. ")
	fmt.Println("Prepare to wait...")

	wp := workerpool.New(cores)

	for i := 0; i < N; i++ {
		start, length := seeds[i*2], seeds[i*2+1]
		wp.Submit(func() { solve(start, length, maps, results) })
	}

	var minlocation int64 = math.MaxInt64
	for a := 1; a <= N; a++ {
		minlocation = min(minlocation, <-results)
	}

	wp.StopWait()

	return minlocation
}

func main() {
	args := os.Args[1:]
	fmt.Println("Day 5: If You Give A Seed A Fertilizer")
	fmt.Println("\tPart One:", PartOne("../aoc-inputs/2023/05/input1.txt")) // 484023871
	if slices.Index(args, "--bruteforce") != -1 {
		cores, err := strconv.Atoi(args[len(args)-1])
		if err != nil {
			cores = runtime.NumCPU()
		}
		fmt.Println("\tPart Two:", PartTwo("../aoc-inputs/2023/05/input1.txt", cores)) // 46294175
	} else {
		fmt.Println("\tPart Two: (skipped by default, run with a '--bruteforce [cores]' option and prepare to wait up to 20 min)")
	}
}

// non-bruteforce solution is totally there, but since I'm havidng fun with multithreading,
// bruteforce improvements:
// + Worker pool of # of cores, instead of fan out
// - Start the _longest_ ranges first
