package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PartOne(file string) int64 {
	input, _ := os.ReadFile(file)
	l := strings.Split(strings.TrimSpace(string(input)), "\n")

	times := []int64{}
	json.Unmarshal([]byte("["+strings.Join(strings.Fields(strings.Split(l[0], ": ")[1]), ",")+"]"), &times)

	distances := []int64{}
	json.Unmarshal([]byte("["+strings.Join(strings.Fields(strings.Split(l[1], ": ")[1]), ",")+"]"), &distances)

	total := int64(1)
	for r, t := range times {
		total = total * race(t, distances[r])
	}

	return total
}

func race(t, d int64) (wins int64) {
	for hold := int64(1); hold < t; hold++ {
		if hold*(t-hold) > d {
			wins++
		}
	}
	return wins
}

func PartTwo(file string) int64 {
	input, _ := os.ReadFile(file)
	l := strings.Split(strings.TrimSpace(string(input)), "\n")

	var re = regexp.MustCompile(`(?m)(\d+)`)
	t, _ := strconv.ParseInt(strings.Join(re.FindAllString(l[0], -1), ""), 10, 64)
	d, _ := strconv.ParseInt(strings.Join(re.FindAllString(l[1], -1), ""), 10, 64)

	return race(t, d)
}

func main() {
	fmt.Println("Day 6: Wait for it \n\tPart One:", PartOne("../aoc-inputs/2023/06/input1.txt")) // 500346
	fmt.Println("\tPart Two:", PartTwo("../aoc-inputs/2023/06/input1.txt"))
	fmt.Println("\tPart Two Solving quadratic equation:", PartTwoQuadratic("../aoc-inputs/2023/06/input1.txt"))
}

func PartTwoQuadratic(file string) int64 {
	input, _ := os.ReadFile(file)
	l := strings.Split(strings.TrimSpace(string(input)), "\n")

	var re = regexp.MustCompile(`(?m)(\d+)`)
	t, _ := strconv.ParseInt(strings.Join(re.FindAllString(l[0], -1), ""), 10, 64)
	d, _ := strconv.ParseInt(strings.Join(re.FindAllString(l[1], -1), ""), 10, 64)

	return int64((t*t+int64(math.Sqrt(float64(t*t-4*d))))/2) - int64((t*t-int64(math.Sqrt(float64(t*t-4*d))))/2)
}
