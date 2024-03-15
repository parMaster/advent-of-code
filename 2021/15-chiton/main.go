package main

import (
	"container/heap"
	"fmt"
	"image"
	"os"
	"slices"
	"strings"
	"time"
)

type Field map[image.Point]int

func read(file string) (m Field, w int, h int) {
	in, _ := os.ReadFile(file)
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

var moves map[rune]image.Point = map[rune]image.Point{
	'R': {1, 0},
	'L': {-1, 0},
	'U': {0, -1},
	'D': {0, 1},
}
var backwards map[rune]rune = map[rune]rune{
	'R': 'L', 'L': 'R', 'U': 'D', 'D': 'U',
}

type Key struct {
	point image.Point
	dir   rune
}

// Item of a Priority Queue
type Item struct {
	key   Key
	score int
	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func dijkstra(f Field, w, h int) int {
	results := []int{}

	itemD := &Item{
		score: 0,
		key: Key{
			point: image.Point{0, 0},
			dir:   'D',
		},
		index: 0,
	}
	itemR := &Item{
		score: 0,
		key: Key{
			point: image.Point{0, 0},
			dir:   'R',
		},
		index: 0,
	}
	dist := map[Key]int{itemR.key: 0, itemD.key: 0}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, itemR)
	heap.Push(&pq, itemD)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		if item.key.point == image.Pt(w, h) {
			// fmt.Println("Reached ", w, "X", h, " with score:", item.score, item.key.steps)
			results = append(results, item.score)
		}

		for _, dir := range []rune{'U', 'D', 'R', 'L'} {
			pos := item.key.point.Add(moves[dir])

			// (?) in bounds
			if pos.X > w || pos.X < 0 || pos.Y > h || pos.Y < 0 {
				continue
			}

			// (?) backwards
			if item.key.dir == backwards[dir] {
				continue
			}

			pointToCheck := &Item{
				score: item.score + f[pos],
				key:   Key{point: pos, dir: dir},
				// key:   Key{point: pos, dir: dir, steps: steps},
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

// tile the field 5x5
func tile(m Field, w, h int) (Field, int, int) {
	// fmt.Println("Expanding", w, "x", h, "to", w*5, "x", h*5)

	newW := w*5 + 4
	newH := h*5 + 4
	newM := make(Field)
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			for i := 0; i < 5; i++ {
				nv := m[image.Pt(x, y)] + i
				if nv > 9 {
					nv %= 9
				}
				newM[image.Pt(x+(w+1)*i, y)] = nv
			}
		}
	}

	for y := 0; y <= h; y++ {
		for x := 0; x <= newW; x++ {
			for i := 0; i < 5; i++ {
				nv := newM[image.Pt(x, y)] + i
				if nv > 9 {
					nv %= 9
				}
				newM[image.Pt(x, y+(h+1)*i)] = nv
			}
		}
	}

	return newM, newW, newH
}

func solve(file string) (int, int) {
	field, w, h := read(file)

	p1 := dijkstra(field, w, h)

	p2 := 0
	field, w, h = tile(field, w, h)
	p2 = dijkstra(field, w, h)

	return p1, p2
}

func main() {
	start := time.Now()
	fmt.Println("Day 15: Chiton")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 739
	fmt.Println("\tPart Two:", p2) // 3040
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
