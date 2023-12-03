package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(file string) []string {
	input, _ := os.ReadFile(file)
	lines := strings.Fields(strings.TrimSpace(string(input)))

	ll := []string{}
	for _, l := range lines {
		ll = append(ll, l)
	}
	return ll
}

func isNum(c byte) bool {
	if (c-0x30) >= 0 && (c-0x30) <= 9 {
		return true
	}
	return false
}

func isSym(c byte) bool {
	if !isNum(c) && c != 46 {
		return true
	}
	return false
}

// square around the [x,y]
var ax = []int{-1, 0, 1, -1, 1, -1, 0, 1}
var ay = []int{-1, -1, -1, 0, 0, 1, 1, 1}

// run checkfunc for every character around [x, y] y - line, x - # of char in line.
func checkAround(l []string, x, y int, checkFunc func(byte) bool) bool {
	for _, dy := range ay {
		for _, dx := range ax {
			if y+dy > 0 && y+dy < len(l) && x+dx > 0 && x+dx < len(l[y+dy]) {
				if checkFunc(l[y+dy][x+dx]) {
					return true
				}
			}
		}
	}

	return false
}

func PartOne(file string) (sum int) {
	ll := readLines(file)
	for y, l := range ll {
		pn, valid := "", false
		l = strings.TrimSpace(l) + "."

		for x := 0; x < len(l); x++ {
			if isNum(l[x]) {
				pn += string(l[x])
				if checkAround(ll, x, y, isSym) {
					valid = true
				}
			} else {
				if valid && pn != "" {
					n, _ := strconv.Atoi(pn)
					sum += n
				}
				pn, valid = "", false
			}
		}
	}
	return
}

func catchNumber(s string, x int) int {
	if x < 0 || x >= len(s) || !isNum(s[x]) {
		return -1
	}

	pn := string(s[x])
	for i := 1; x+i < len(s) && isNum(s[x+i]); i++ {
		pn += string(s[x+i])
	}
	for i := 1; x-i >= 0 && isNum(s[x-i]); i++ {
		pn = string(s[x-i]) + pn
	}

	n, _ := strconv.Atoi(pn)
	return n
}

// check around, collect the numbers, multiply
func grindTheGearsAround(l []string, x, y int) int {
	n, n1 := 0, 0
	for _, dy := range ay {
		for _, dx := range ax {
			if y+dy >= 0 && y+dy < len(l) && x+dx >= 0 && x+dx < len(l[y+dy]) {
				n = catchNumber(l[y+dy], x+dx)
				if n > 0 {
					if n1 == 0 {
						n1 = n
						continue
					}
					if n1 != n {
						return n * n1
					}
				}
			}
		}
	}
	return 0
}

func PartTwo(file string) (sum int) {
	ll := readLines(file)

	for y, l := range ll {
		l = strings.TrimSpace(l) + "."

		for x := 0; x < len(l); x++ {
			if byte('*') == l[x] {
				n := grindTheGearsAround(ll, x, y)
				sum += n
			}
		}
	}

	return
}

func main() {
	fmt.Println("Day 3: Gear Ratios\n\tPart One:", PartOne("input1.txt"))
	fmt.Println("\tPart Two:", PartTwo("input1.txt"))
}
