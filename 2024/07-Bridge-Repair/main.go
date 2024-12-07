package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// https://adventofcode.com/2024/day/7

func check(test, acc uint64, nums []uint64, concat bool) bool {
	if len(nums) == 0 {
		if test == acc {
			return true
		}
		return false
	}
	if acc > test { // early stop
		return false
	}

	res := false
	var con uint64
	if concat {
		fmt.Sscanf(fmt.Sprintf("%d%d", acc, nums[0]), "%d", &con)
		res = res || check(test, con, nums[1:], concat)
	}

	return res ||
		check(test, acc+nums[0], nums[1:], concat) ||
		check(test, acc*nums[0], nums[1:], concat)
}

func solve(f string) (p1, p2 uint64) {
	in, _ := os.ReadFile(f)
	for _, l := range strings.Split(string(in), "\n") {
		var nums []uint64
		json.Unmarshal([]byte("["+strings.ReplaceAll(strings.ReplaceAll(l, ":", ""), " ", ",")+"]"), &nums)
		test := nums[0]
		nums = nums[1:]

		if check(test, nums[0], nums[1:], false) {
			p1 += test
		} else if check(test, nums[0], nums[1:], true) {
			p2 += test
		}
	}
	p2 += p1

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 07: Bridge Repair")
	// p1, p2 := solve("input_pub.txt")
	p1, p2 := solve("../aoc-inputs/2024/07/input.txt")
	fmt.Println("\tPart One:", p1) // 3119088655389
	fmt.Println("\tPart Two:", p2) // 264184041398847
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
