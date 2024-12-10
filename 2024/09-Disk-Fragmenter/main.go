package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func p1(m []int) (p1 uint64) {
	var left, i uint64
	right := uint64(len(m) - 1)
	for left <= right {

		// file
		for range m[left] {
			p1 += i * (left / 2)
			i++
		}
		left += 1

		// filling space
		for j := 0; j < m[left] && left < right; j++ {
			m[right] -= 1
			p1 += i * (right / 2)
			i++
			for m[right] == 0 {
				right -= 2
			}
		}
		left += 1
	}
	return
}

type num struct {
	space bool
	size  int
	val   int
}

func calc(nums []num) (sum uint64) {
	i := 0
	for _, n := range nums {
		for range n.size {
			if !n.space {
				sum += uint64(i * n.val)
			}
			i++
		}
	}
	return
}

func p2(in []int) (res uint64) {
	m := []num{}
	for i := 0; i < len(in); i++ {
		n := num{space: i%2 != 0, size: in[i]}
		if !n.space {
			n.val = i / 2
		}
		m = append(m, n)
	}

	retry := true
	for retry {
		retry = false
		for right := len(m) - 1; right > 0; right -= 1 {
			if !m[right].space {
				for left, t := range m[:right] {
					if t.space && t.size >= m[right].size {
						m[left] = m[right]
						m[right].space = true
						if t.size > m[right].size {
							m = slices.Insert(m, left+1, num{true, t.size - m[right].size, 0})
						}
						retry = true
						break
					}
				}
				if retry {
					break
				}
			}
		}
	}

	show(m)
	return calc(m)
}

func solve(file string) (uint64, uint64) {
	in, _ := os.ReadFile(file)
	// in = []byte("2333133121414131402") // 1928, 2858

	m := []int{}
	for _, c := range in {
		d := int(c - '0')
		m = append(m, d)
	}

	return p1(slices.Clone(m)), p2(slices.Clone(m))
}

func main() {
	start := time.Now()
	fmt.Println("Day 09: Disk Fragmenter")
	p1, p2 := solve("../aoc-inputs/2024/09/input.txt")
	fmt.Println("\tPart One:", p1) // 6432869891895
	fmt.Println("\tPart Two:", p2) // 6467290479134
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}

func show(nums []num) (sum uint64) {
	fmt.Println()
	for _, n := range nums {
		if !n.space {
			fmt.Print(strings.Repeat(fmt.Sprintf("%d", n.val), n.size))
		} else {
			fmt.Print(strings.Repeat(string("."), n.size))
		}
		fmt.Print("|")
	}
	fmt.Println()
	return
}
