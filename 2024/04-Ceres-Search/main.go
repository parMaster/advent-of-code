package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var dir = [][2]int{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}

type Field [][]rune

func read(file string) Field {
	in, _ := os.ReadFile(file)
	f := Field{}
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		f = append(f, []rune{})
		for _, r := range strings.TrimSpace(l) {
			f[y] = append(f[y], r)
		}
	}
	return f
}

func (f Field) searchDir(term string, y, x int, dir [2]int) (res int) {
	r := rune(term[0])
	if y < 0 || y >= len(f) || x < 0 || x >= len(f[y]) {
		return
	}

	if f[y][x] == r {
		if len(term) == 1 {
			return 1
		}
		ny := y + dir[0]
		nx := x + dir[1]
		res += f.searchDir(term[1:], ny, nx, dir)
	}

	return res
}

// M.S
// .A.
// M.S
func (f Field) checkMAS(y, x int) (res int) {
	if f[y][x] != 'A' {
		return
	}

	if ((f[y-1][x-1] == 'M' && f[y+1][x+1] == 'S') || (f[y-1][x-1] == 'S' && f[y+1][x+1] == 'M')) &&
		((f[y+1][x-1] == 'M' && f[y-1][x+1] == 'S') || (f[y+1][x-1] == 'S' && f[y-1][x+1] == 'M')) {
		res = 1
	}

	return
}

func solve(f string) (p1, p2 int) {
	field := read(f)

	// Part 1 - search XMAS horiz, vert, diag, inverted
	for y := range field {
		for x := range field[y] {
			for _, d := range dir {
				p1 += field.searchDir("XMAS", y, x, d)
			}
		}
	}

	// Part 2 - searching two MAS in the shape of an X
	for y := 1; y < len(field)-1; y++ {
		for x := 1; x < len(field[y])-1; x++ {
			p2 += field.checkMAS(y, x)
		}
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 04: Ceres Search")
	// p1, p2 := solve("input_pub.txt")
	p1, p2 := solve("../aoc-inputs/2024/04/input.txt")
	fmt.Println("\tPart One:", p1) // 2507
	fmt.Println("\tPart Two:", p2) // 1969
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
