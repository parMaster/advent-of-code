package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func solve(file string) (p1, p2 int) {
	in, _ := os.ReadFile(file)
	lines := strings.Split(string(in), "\n")

	oper := []string{}
	json.Unmarshal([]byte("[\""+strings.Join(strings.Fields(lines[len(lines)-1]), "\",\"")+"\"]"), &oper)

	m := [][]int{}
	for _, line := range lines[:len(lines)-1] {
		var l []int
		json.Unmarshal([]byte("["+strings.Join(strings.Fields(line), ",")+"]"), &l)
		m = append(m, l)
	}

	res := m[0]
	for _, line := range m[1:] {
		for i, v := range line {
			switch oper[i] {
			case "*":
				res[i] *= v
			case "+":
				res[i] += v
			}
		}
	}

	for _, v := range res {
		p1 += v
	}

	// p2
	nums := []string{}
	for i := len(lines[0]) - 1; i >= 0; i-- {
		num := ""
		for j := range len(lines) {

			if j == len(lines)-1 {
				nums = append(nums, num)
				num = ""
			}

			switch lines[j][i] {
			case '*':
				res := 1
				for _, v := range nums {
					if intV, _ := strconv.Atoi(strings.Trim(v, " ")); intV > 0 {
						res *= intV
					}
				}
				p2 += res
				nums = []string{}
			case '+':
				res := 0
				for _, v := range nums {
					intV, _ := strconv.Atoi(strings.Trim(v, " "))
					res += intV
				}
				p2 += res
				nums = []string{}

			default:
				num = num + string(lines[j][i])
			}
		}

	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 06: Trash Compactor")
	p1, p2 := solve("../aoc-inputs/2025/06/input.txt")
	// p1, p2 := solve("input-pub.txt")
	fmt.Println("\tPart One:", p1) // 6757749566978
	fmt.Println("\tPart Two:", p2) // 10603075273949
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
