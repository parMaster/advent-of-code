package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
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

	var r1, r2, rc bool
	if concat {
		var con uint64
		fmt.Sscanf(fmt.Sprintf("%d%d", acc, nums[0]), "%d", &con)
		rc = check(test, con, nums[1:], concat)
	}
	r1 = check(test, acc+nums[0], nums[1:], concat)

	r2 = check(test, acc*nums[0], nums[1:], concat)

	return rc || r1 || r2
}

func solve(f string) (p1, p2 uint64) {
	in, _ := os.ReadFile(f)
	lines := strings.Split(string(in), "\n")

	throttle := make(chan struct{}, runtime.NumCPU())
	var rw sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(lines))

	for _, l := range lines {
		var nums []uint64
		json.Unmarshal([]byte("["+strings.ReplaceAll(strings.ReplaceAll(l, ":", ""), " ", ",")+"]"), &nums)
		test := nums[0]
		nums = nums[1:]

		throttle <- struct{}{}
		go func() {
			var p1r, p2r uint64
			if check(test, nums[0], nums[1:], false) {
				p1r = test
			} else if check(test, nums[0], nums[1:], true) {
				p2r = test
			}
			rw.Lock()
			p1 += p1r
			p2 += p2r
			rw.Unlock()
			<-throttle
			wg.Done()
		}()
	}
	wg.Wait()
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
