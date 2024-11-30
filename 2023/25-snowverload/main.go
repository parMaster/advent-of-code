package main

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"
	"time"

	"container/heap"
)

// traversing graph once, keeps history of visited nodes and minds the cycles
func traverseOnce(g map[string][]string, source string, f func(currentNode string)) {

	s := NewStack(string(""))
	s.Push(source)

	history := make(map[string]bool)

	for !s.IsEmpty() {
		currentNode := *s.PopFirst()
		if !history[currentNode] {
			f(currentNode)
			history[currentNode] = true
		}
		for _, v := range g[currentNode] {
			_, ok := history[v]
			if !ok {
				s.Push(v)
			}
		}
	}
}

// Item of a Priority Queue
type Item struct {
	score int
	index int
	node  string
	path  []string
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func dijkstra(g map[string][]string, start, finish string, avoidConnections []string) []string {
	result := []string{}

	startItem := &Item{
		score: 0,
		index: 0,
		node:  start,
		path:  []string{start},
	}

	dist := map[string]int{start: 0}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, startItem)

	minScore := 999999
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		if item.node == finish {
			if item.score < minScore {
				minScore = item.score
			}
			result = item.path
		}

		for _, v := range g[item.node] {

			// avoid connections
			if slices.Contains(avoidConnections, item.node+"-"+v) {
				continue
			}

			if _, ok := dist[v]; ok && dist[v] <= item.score+1 {
				// already was here with the same key and better result
				continue
			}

			pointToCheck := &Item{
				score: item.score + 1,
				node:  v,
				path:  append(item.path, v),
			}

			dist[v] = item.score + 1

			heap.Push(&pq, pointToCheck)
		}
	}
	return result
}

func path_to_connections(strPath []string) []string {
	path := strPath
	connections := []string{}
	for i := 0; i < len(path)-1; i++ {
		connections = append(connections, path[i]+"-"+path[i+1])
		connections = append(connections, path[i+1]+"-"+path[i])
	}
	return connections
}

func countIslands(graph map[string][]string) (bool, int) {

	nodes := map[string]bool{}
	for k := range graph {
		nodes[k] = false
	}

	island := 0
	islandsSize := map[int]int{}
	for node, visited := range nodes {
		if !visited {
			island++
			traverseOnce(graph, node, func(currentNode string) {
				nodes[currentNode] = true
				islandsSize[island]++
			})
		}
	}
	// fmt.Println(islandsSize)

	if len(islandsSize) == 2 {
		// fmt.Println(islanSize)
		return true, islandsSize[1] * islandsSize[2]
	}
	return false, 0
}

func solve(filename string) (int, int) {
	txt, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(txt)), "\n")

	g := map[string][]string{}
	for _, line := range lines {
		from := strings.Split(line, ": ")[0]
		if _, ok := g[from]; !ok {
			g[from] = []string{}
		}
		to := strings.Fields(strings.Split(line, ": ")[1])
		for _, t := range to {
			g[from] = append(g[from], t)
			if _, ok := g[t]; !ok {
				g[t] = []string{}
			}
			g[t] = append(g[t], from)
		}
	}

	var keys []string
	for k := range g {
		keys = append(keys, k)
	}

	for {
		k1 := keys[rand.Intn(len(keys))]
		k2 := keys[rand.Intn(len(keys))]
		if k1 == k2 {
			continue
		}

		ok, res := checkPair(g, k1, k2)
		if ok {
			return res, 0
		}
	}
}

func checkPair(g map[string][]string, k1, k2 string) (bool, int) {
	avoidConnections := []string{}
	for range 3 {
		path := dijkstra(g, k1, k2, avoidConnections)
		// fmt.Println("Found path:", path)
		avoidConnections = append(avoidConnections, path_to_connections(path)...)
	}

	cutGraph := map[string][]string{}
	if len(avoidConnections) > 0 {
		for k, v := range g {
			for _, vv := range v {
				if !slices.Contains(avoidConnections, k+"-"+vv) && !slices.Contains(avoidConnections, vv+"-"+k) {
					cutGraph[k] = append(cutGraph[k], vv)
				}
			}
		}
	}

	return countIslands(cutGraph)
}

func main() {
	start := time.Now()
	fmt.Println("Day 25: Snowverload (wait a minute... calculating)")
	p1, _ := solve("../aoc-inputs/2023/25/input.txt")
	fmt.Println("\tPart One:", p1) // 569904
	// fmt.Println("\tPart Two:", p2) //
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
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
