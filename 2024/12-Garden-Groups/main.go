package main

import (
	"fmt"
	"image"
	"os"
	"strings"
	"time"
)

type Field map[image.Point]rune

// ReadField reads a string and returns a Field, width and height
func ReadField(in string) (f Field, w int, h int) {
	f = make(map[image.Point]rune)
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		w = len(l)
		h = y + 1
		for x, r := range strings.TrimSpace(l) {
			f[image.Point{x, y}] = r
		}
	}
	return
}

// directions	↑ → ↓ ←
// directions	0 1 2 3
var XYDir = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

// diagonal directions
var DiagDir = []image.Point{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}}

// counting corners because sides == corners
func (f Field) corners(pt image.Point) (c int) {
	neighbors := 0
	for i := range XYDir {
		if f[pt.Add(XYDir[i])] == f[pt] {
			neighbors++
		}
	}

	switch neighbors {
	case 0:
		// point
		return 4
	case 1:
		// ends
		return 2
	case 4:
		// middle or +
		for _, dd := range DiagDir {
			if f[pt.Add(dd)] != f[pt] {
				c++
			}
		}
		return c
	case 2:
		// line: | or  -
		if (f[pt.Add(XYDir[0])] == f[pt] && f[pt.Add(XYDir[2])] == f[pt] &&
			f[pt.Add(XYDir[1])] != f[pt] && f[pt.Add(XYDir[3])] != f[pt]) ||
			(f[pt.Add(XYDir[1])] == f[pt] && f[pt.Add(XYDir[3])] == f[pt] &&
				f[pt.Add(XYDir[0])] != f[pt] && f[pt.Add(XYDir[2])] != f[pt]) {
			return 0
		}
		// L
		for i := range XYDir {
			c1 := pt.Add(XYDir[i])
			c2 := pt.Add(XYDir[(i+1)%4])
			if f[c1] == f[pt] && f[c2] == f[pt] {
				// inner corner?
				if f[pt.Add(XYDir[i]).Add(XYDir[(i+1)%4])] != f[pt] {
					return 2
				}
				return 1
			}
		}
	default:
		// T in every direction
		for i := range XYDir {
			c1 := pt.Add(XYDir[i])
			c2 := pt.Add(XYDir[(i+1)%4])
			c3 := pt.Add(XYDir[(i+2)%4])
			if f[c1] == f[pt] && f[c2] == f[pt] && f[c3] == f[pt] {
				// inner corner 1?
				if f[pt.Add(XYDir[i]).Add(XYDir[(i+1)%4])] != f[pt] {
					c++
				}
				// inner corner 2?
				if f[pt.Add(XYDir[(i+1)%4]).Add(XYDir[(i+2)%4])] != f[pt] {
					c++
				}
			}
		}
	}

	return
}

func (f Field) perimeter(pt image.Point) (p int) {
	for _, dir := range XYDir {
		if rn, ok := f[pt.Add(dir)]; !ok || f[pt] != rn {
			p++
		}
	}

	return
}

func solve(file string) (p1, p2 int) {
	in, _ := os.ReadFile(file)
	field, w, h := ReadField(string(in))
	field.Show(w, h)

	seen := map[image.Point]bool{}
	for p := range field {
		if seen[p] {
			continue
		}

		seen[p] = true
		region := map[image.Point]bool{p: true}

		q := []image.Point{p}
		for len(q) > 0 {
			newQ := []image.Point{}
			for _, qitem := range q {
				for _, step := range XYDir {
					check := qitem.Add(step)
					if !seen[check] && field[check] == field[p] {
						newQ = append(newQ, check)
						seen[check] = true
						region[check] = true
					}
				}
			}
			q = newQ
		}

		sides := 0
		perimeter := 0
		for pt := range region {
			sides += field.corners(pt)
			perimeter += field.perimeter(pt)
		}
		p1 += len(region) * perimeter
		p2 += len(region) * sides
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 12: Garden Groups")
	// p1, p2 := solve("input2.txt")
	p1, p2 := solve("../aoc-inputs/2024/12/input.txt")
	fmt.Println("\tPart One:", p1) // 1467094
	fmt.Println("\tPart Two:", p2) // 881182
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}

// Show prints the Field
func (f Field) Show(w, h int) {
	fmt.Println(w, "x", h, ":")
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			fmt.Print(string(f[image.Pt(x, y)]))
		}
		fmt.Println()
	}
	fmt.Println()
}
