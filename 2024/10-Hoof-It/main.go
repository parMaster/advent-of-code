package main

import (
	"fmt"
	"image"
	"os"
	"strings"
	"time"
)

type Field map[image.Point]byte

// ReadField reads a string and returns a Field, width and height
func ReadField(in string) (f Field, th []image.Point, w int, h int) {
	f = make(map[image.Point]byte)
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		w = len(l) - 1
		h = y
		for x, r := range strings.TrimSpace(l) {
			val := byte(r - '0')
			f[image.Point{x, y}] = val
			if val == 0 {
				th = append(th, image.Point{x, y})
			}
		}
	}
	return
}

var XYDir = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func (f Field) search(start image.Point, bounds image.Rectangle, height byte) (finish []image.Point) {
	if height == 9 {
		return []image.Point{start}
	}

	for _, step := range XYDir {
		next := start.Add(step)
		if next.In(bounds) && f[next] == height+1 {
			finish = append(finish, f.search(next, bounds, f[next])...)
		}
	}

	return finish
}

func solve(file string) (p1, p2 int) {
	in, _ := os.ReadFile(file)
	field, trailheads, w, h := ReadField(string(in))
	bounds := image.Rect(0, 0, w+1, h+1)

	for _, t := range trailheads {
		finishes := field.search(t, bounds, 0)

		p2 += len(finishes)

		uniq := map[image.Point]struct{}{}
		for _, f := range finishes {
			uniq[f] = struct{}{}
		}
		p1 += len(uniq)
	}

	return
}

func main() {
	start := time.Now()
	fmt.Println("Day 10: Hoof It")
	p1, p2 := solve("../aoc-inputs/2024/10/input.txt")
	fmt.Println("\tPart One:", p1) // 719
	fmt.Println("\tPart Two:", p2) // 1530
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
