package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func solve(file string) (int64, int) {
	txt, _ := os.ReadFile(file)
	parts := strings.Split(strings.TrimSpace(string(txt)), "\n\n")
	seed := strings.TrimSpace(parts[0])
	rulesLines := strings.Split(strings.TrimSpace(parts[1]), "\n")
	rules := map[string]string{}
	for _, line := range rulesLines {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}

	// fmt.Println("Seed:", seed)
	// fmt.Println("Rules:", rules)

	for range 10 {
		for i := 0; i < len(seed)-1; i++ {
			if val, ok := rules[seed[i:i+2]]; ok {
				seed = seed[:i+1] + val + seed[i+1:]
				i++
			}
		}
	}

	freq := map[byte]int64{}
	for i := 0; i < len(seed); i++ {
		freq[seed[i]]++
	}
	minFreq, maxFreq := int64(math.MaxInt64), int64(0)
	for _, v := range freq {
		minFreq = min(minFreq, v)
		maxFreq = max(maxFreq, v)
	}

	return maxFreq - minFreq, 0
}

func main() {
	start := time.Now()
	fmt.Println("Day 14: Extended Polymerization")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 3555
	fmt.Println("\tPart Two:", p2) //
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
