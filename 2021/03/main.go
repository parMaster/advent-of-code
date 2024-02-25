package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func countBits(lines []string) map[int]int {
	m := map[int]int{}
	for _, l := range lines {
		for i, c := range l {
			if _, ok := m[i]; !ok {
				m[i] = 0
			}
			if c == '0' {
				m[i]++
			}
		}
	}
	return m
}

func binStringToInt(s string) int {
	n := 0
	for i, c := range s {
		n += int(c-'0') << (len(s) - 1 - i)
	}
	return n
}

func solve(filename string) (int, int) {
	f, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	m := countBits(lines)

	gamma, epsilon := 0, 0
	for k, v := range m {
		if v > len(lines)/2 {
			gamma += 1 << ((len(m) - 1) - k)
		} else {
			epsilon += 1 << ((len(m) - 1) - k)
		}
	}

	lifeSupport := func(linesInput []string, bitCriteria byte) int {
		lines := slices.Clone(linesInput)
		bit := 0
		var consider byte
		for len(lines) > 1 {
			o2lines := []string{}
			m := countBits(lines)
			if bitCriteria == '1' {
				consider = byte('0')
			} else {
				consider = byte('1')
			}
			if m[bit] > len(lines)/2 {
				consider = bitCriteria
			}

			for i := 0; i < len(lines); i++ {
				if lines[i][bit] == consider {
					o2lines = append(o2lines, lines[i])
				}
			}
			lines = slices.Clone(o2lines)
			bit++
		}
		return binStringToInt(lines[0])
	}

	return gamma * epsilon, lifeSupport(lines, '0') * lifeSupport(lines, '1')
}

func main() {
	start := time.Now()
	fmt.Println("Day 3: Binary Diagnostic")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 3009600
	fmt.Println("\tPart Two:", p2) // 6940518
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
