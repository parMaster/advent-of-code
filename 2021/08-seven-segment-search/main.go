package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func sortLetters(words []string) []string {
	sorted := make([]string, 0, len(words))
	for _, word := range words {
		letters := strings.Split(word, "")
		slices.Sort(letters)
		sorted = append(sorted, strings.Join(letters, ""))
	}
	return sorted
}

// removes s2 segments from s1
func remove(s1 string, s2 string) string {
	for _, c := range s2 {
		s1 = strings.ReplaceAll(s1, string(c), "")
	}
	return s1
}

func deduce(patterns []string) map[int]string {
	m := map[int]string{}
	for _, pattern := range patterns {
		switch len(pattern) {
		case 2:
			m[1] = pattern
		case 3:
			m[7] = pattern
		case 4:
			m[4] = pattern
		case 7:
			m[8] = pattern
		}
	}

	for _, pattern := range patterns {
		switch len(pattern) {
		case 5:
			// 3: only 5-segment that includes "7" segments is "3"
			if remove(m[7], pattern) == "" {
				m[3] = pattern
			} else {
				// 2 and 5: remove every "5" segment from "4" and if one is left, it's "5"
				if len(remove(m[4], pattern)) == 1 {
					m[5] = pattern
				} else {
					// otherwise it's "2"
					m[2] = pattern
				}
			}
		case 6:
			// 9: "4" removed from "9" leaves only 2 segments
			if len(remove(pattern, m[4])) == 2 {
				m[9] = pattern
			} else {
				// 0
				if len(remove(pattern, m[7])) == 3 {
					m[0] = pattern
				}
				// 6
				if len(remove(pattern, m[7])) == 4 {
					m[6] = pattern
				}
			}
		}
	}

	return m
}

func solve(file string) (int, int) {
	data, _ := os.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	p1 := 0
	p2 := 0
	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), " | ")

		patterns := strings.Split(parts[0], " ")
		patterns = sortLetters(patterns)

		digits := strings.Split(parts[1], " ")
		digits = sortLetters(digits)
		for _, digit := range digits {
			if slices.Index([]int{2, 4, 3, 7}, len(digit)) != -1 {
				p1++
			}
		}

		m := deduce(patterns)
		result := ""
		for _, digit := range digits {
			for k, v := range m {
				if digit == v {
					result += fmt.Sprint(k)
				}
			}
		}
		intResult, _ := strconv.Atoi(result)
		p2 += intResult
	}

	return p1, p2
}

func main() {
	start := time.Now()
	fmt.Println("Day 8: Seven Segment Search")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 369
	fmt.Println("\tPart Two:", p2) // 1031553
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
