package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func solve(file string) (int64, int64) {
	txt, _ := os.ReadFile(file)
	parts := strings.Split(strings.TrimSpace(string(txt)), "\n\n")
	seed := strings.TrimSpace(parts[0])
	rulesLines := strings.Split(strings.TrimSpace(parts[1]), "\n")
	rules := map[string][]string{}
	for _, line := range rulesLines {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = []string{string(parts[0][0]) + parts[1], parts[1] + string(parts[0][1])}
	}

	// fmt.Println("Seed:", seed)
	// fmt.Println("Rules:", rules)

	var evolve = func(seed string, rules map[string][]string, n int) int64 {

		m := map[string]int64{}
		for i := 0; i < len(seed)-1; i++ {
			m[seed[i:i+2]]++
		}

		for range n {
			mm := map[string]int64{}
			for k, v := range m {
				if newPair, ok := rules[k]; ok {
					mm[newPair[0]] += v
					mm[newPair[1]] += v
				} else {
					mm[k] = v
				}
			}
			m = mm
		}
		// fmt.Println("Pairs Counts:", m)

		freq := map[byte]int64{}
		for k, v := range m {
			freq[k[1]] += v
		}

		minFreq, maxFreq := int64(math.MaxInt64), int64(0)
		for _, v := range freq {
			minFreq = min(minFreq, v)
			maxFreq = max(maxFreq, v)
		}

		return maxFreq - minFreq
	}

	return evolve(seed, rules, 10), evolve(seed, rules, 40)
}

func main() {
	start := time.Now()
	fmt.Println("Day 14: Extended Polymerization")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 3555
	fmt.Println("\tPart Two:", p2) // 4439442043739
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
