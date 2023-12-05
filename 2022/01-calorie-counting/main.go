package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func maxcalories(filename string) (int, error) {

	var currentMax, currentSum, current int

	f, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}

	b := bufio.NewReader(f)
	for {
		l, err := b.ReadString('\n')
		if err != nil {
			break
		}

		l = strings.TrimSpace(l)

		if l != "" {
			current, err = strconv.Atoi(l)
			if err != nil {
				return 0, fmt.Errorf("atoi error: %w", err)
			}

			currentSum += current
		}

		if l == "" {
			currentMax = max(currentMax, currentSum)
			currentSum = 0
		}
	}

	return currentMax, nil
}

func top3calories(filename string) (int, error) {

	var currentMax, currentSum, current int
	var sums sort.IntSlice

	f, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}

	b := bufio.NewReader(f)
	for {
		l, err := b.ReadString('\n')
		if err != nil {
			break
		}

		l = strings.TrimSpace(l)

		if l != "" {
			current, err = strconv.Atoi(l)
			if err != nil {
				return 0, fmt.Errorf("atoi error: %w", err)
			}

			currentSum += current
		}

		if l == "" {
			sums = append(sums, currentSum)
			currentSum = 0
		}
	}

	sort.Sort(sort.Reverse(sums))

	currentMax = sums[0] + sums[1] + sums[2]

	return currentMax, nil
}

func main() {

	maxCalories, err := maxcalories("../aoc-inputs/2022/01/input.txt")
	if err != nil {
		fmt.Printf("maxcalories returned error: %e", err)
	}

	fmt.Println("Day 1: Calorie Counting\n\tPart One:", maxCalories)

	top3Calories, err := top3calories("../aoc-inputs/2022/01/input.txt")
	if err != nil {
		fmt.Printf("top3calories returned error: %e", err)
	}

	fmt.Println("\tPart Two:", top3Calories)
}
