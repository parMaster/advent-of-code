package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func solve(f string) (p1, p2 int) {
	nums := read(f)

	for i, n1 := range nums {
		for j, n2 := range nums[i:] {
			if n1+n2 == 2020 {
				p1 = n1 * n2
			}
			for _, n3 := range nums[j:] {
				if n1+n2+n3 == 2020 {
					p2 = n1 * n2 * n3
				}
			}
		}
	}

	return p1, p2
}

func read(filename string) []int {
	fc, _ := os.ReadFile(filename)

	lines := strings.Split(string(fc), "\n")
	nums := []int{}
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}
	return nums
}

func main() {
	start := time.Now()
	fmt.Println("Day 01: Report Repair")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 1006875
	fmt.Println("\tPart Two:", p2) // 165026160
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
