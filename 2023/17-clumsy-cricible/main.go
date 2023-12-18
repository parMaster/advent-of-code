// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
	"image"
	"os"
	"slices"
	"strings"
)

var moves map[rune]image.Point = map[rune]image.Point{
	'R': {1, 0},
	'L': {-1, 0},
	'U': {0, -1},
	'D': {0, 1},
}
var backwards map[rune]rune = map[rune]rune{
	'R': 'L', 'L': 'R', 'U': 'D', 'D': 'U',
}

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

func main() {
	fmt.Println("Day 17: Clumsy Crucible")
	fmt.Println("\tPart One:", p1("../aoc-inputs/2023/17/input.txt")) // 785
	// fmt.Println("\tPart Two:", p2("../aoc-inputs/2023/17/input.txt")) //
}

type Field map[image.Point]int

var f Field
var w, h int

type Key struct {
	point image.Point
	steps int
	dir   rune
}

// Item of a Priority Queue
type Item struct {
	key   Key
	path  string
	score int
	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func p1(file string) int {
	results := []int{}

	in, _ := os.ReadFile(file)
	f, w, h = read(string(in))
	// show(f, w, h)

	itemR := &Item{
		score: 0,
		key: Key{
			point: image.Point{0, 0},
			dir:   'R', // it seems to be enough to have only one starting key, 'D' is redundant?
		},
		index: 0,
	}
	dist := map[Key]int{itemR.key: 0}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, itemR)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		if item.key.point == image.Pt(w, h) {
			// fmt.Println("Reached ", w, "X", h, " with score:", item.score)
			results = append(results, item.score)
		}

		for _, dir := range []rune{'U', 'D', 'R', 'L'} {
			pos := item.key.point.Add(moves[dir])

			// (?) in bounds
			if pos.X > w || pos.X < 0 || pos.Y > h || pos.Y < 0 {
				continue
			}

			steps := 1
			if item.key.dir == dir {
				steps = item.key.steps + 1
			}

			// fourth step in the same direction
			if steps == 4 {
				continue
			}

			// (?) backwards
			if item.key.dir == backwards[dir] {
				continue
			}

			pointToCheck := &Item{
				score: item.score + f[pos],
				key:   Key{point: pos, dir: dir, steps: steps},
			}

			if _, ok := dist[pointToCheck.key]; ok && dist[pointToCheck.key] <= item.score+f[pos] {
				// already was here with the same key and better result
				continue
			}

			dist[pointToCheck.key] = item.score + f[pos]

			heap.Push(&pq, pointToCheck)
		}
	}
	return slices.Min(results)
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.path = value
	item.score = priority
	heap.Fix(pq, item.index)
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
