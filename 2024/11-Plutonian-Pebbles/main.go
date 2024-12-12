package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func blink(stone int) (stones []int) {
	if stone == 0 {
		stones = []int{1}
	} else if ss := strconv.Itoa(stone); len(ss)%2 == 0 {
		s1, s2 := ss[:len(ss)/2], ss[len(ss)/2:]
		d1, _ := strconv.Atoi(s1)
		d2, _ := strconv.Atoi(s2)
		stones = []int{d1, d2}
	} else {
		stones = []int{stone * 2024}
	}

	return stones
}

func freqSum(freqs map[int]int) (sum int) {
	for _, f := range freqs {
		sum += f
	}
	return
}

func solve(f string) (p1, p2 int) {
	in, _ := os.ReadFile(f)
	freqs := map[int]int{}

	for _, sn := range strings.Split(string(in), " ") {
		n, _ := strconv.Atoi(sn)
		freqs[n] = 1
	}

	// "lanternfish" solution - counting frequencies
	for i := range 75 {
		if i == 25 {
			p1 = freqSum(freqs)
		}
		stepFreqs := map[int]int{}
		for stone, freq := range freqs {
			stones := blink(stone)
			for _, s := range stones {
				stepFreqs[s] += freq
			}
		}
		freqs = stepFreqs
	}

	return p1, freqSum(freqs)
}

func main() {
	start := time.Now()
	fmt.Println("Day 11: Plutonian Pebbles")
	p1, p2 := solve("../aoc-inputs/2024/11/input.txt")
	fmt.Println("\tPart One:", p1) // 233875
	fmt.Println("\tPart Two:", p2) // 277444936413293
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
