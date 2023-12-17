package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

var moves = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

type Field map[image.Point]int

type QItem struct {
	score    int
	lastMove image.Point
	stepSize int
}

type Queue map[image.Point]QItem

func read(in string) (m Field, w int, h int) {
	m = make(Field)
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		w = len(l) - 1
		h = y
		for x, r := range strings.TrimSpace(l) {
			m[image.Point{x, y}] = int(r - 0x30)
		}
	}
	return
}

func minTravel(f Field, w, h int) int {

	show(f, w, h)
	start := image.Point{0, 0}
	// starting point and initial step direction and size
	queue := Queue{start: {0, image.Point{1, 0}, 1}}
	visited := Field{start: 0}

	for len(queue) > 0 {
		nextQueue := Queue{}
		for point, q := range queue {
			for _, m := range moves {
				pos := point.Add(m)

				// check out of bounds
				if pos.X > w || pos.X < 0 || pos.Y > h || pos.Y < 0 {
					continue
				}

				stepSize := q.stepSize
				// switched direction?
				if q.lastMove != m {
					stepSize = 1
				}

				// stepped more that 3 times in one direction ?
				if stepSize > 3 {
					continue
				}

				stepSize++

				// moving back ?
				if q.lastMove.Add(m) == image.Pt(0, 0) {
					continue
				}

				// if not visited or found shorter way
				newScore := q.score + f[pos]
				if lastScore, ok := visited[pos]; !ok || newScore < lastScore {
					nextQueue[pos] = QItem{newScore, m, stepSize}
					visited[pos] = newScore
				}
			}
		}
		queue = nextQueue
	}

	show(visited, w, h)

	return visited[image.Pt(w, h)]
}

func p1(f string) int {
	in, _ := os.ReadFile(f)
	field, w, h := read(string(in))
	return minTravel(field, w, h)
}

func main() {
	fmt.Println("Day 17: Clumsy Crucible")
	fmt.Println("\tPart One:", p1("../aoc-inputs/2023/17/input.txt")) // 772 low
	// fmt.Println("\tPart Two:", p2("../aoc-inputs/2023/16/input.txt")) // 7154
}

func show(m map[image.Point]int, w, h int) {
	fmt.Println(w+1, "x", h+1, ":")
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			fmt.Print((m[image.Pt(x, y)]), " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
