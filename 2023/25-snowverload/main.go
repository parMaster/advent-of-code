package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/exp/maps"
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

func solve(filename string) (int, int) {
	txt, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(txt)), "\n")

	g := map[string][]string{}
	connections := map[string][2]string{}
	nodes := map[string]bool{}
	for _, line := range lines {
		from := strings.Split(line, ": ")[0]
		if _, ok := g[from]; !ok {
			g[from] = []string{}
		}
		nodes[from] = false
		to := strings.Fields(strings.Split(line, ": ")[1])
		for _, t := range to {
			g[from] = append(g[from], t)
			if _, ok := g[t]; !ok {
				g[t] = []string{}
			}
			g[t] = append(g[t], from)
			connections[from+"-"+t] = [2]string{from, t}
			nodes[t] = false
		}
	}

	// connList := maps.Keys(connections)
	// slices.Sort(connList)
	fmt.Println(len(connections))

	iter := 0
	for c1, c1v := range connections {
		for c2, c2v := range connections {
			for c3, c3v := range connections {
				if c1 != c2 && c1 != c3 && c2 != c3 {
					iter++
					if iter%1000 == 0 {
						fmt.Println(iter)
						// fmt.Println(c1v, c2v, c3v)
					}
					// fmt.Println(c1, c2, c3)

					// Deep copy of the graph
					graph := map[string][]string{}
					for k, v := range g {
						for _, vv := range v {
							if (k == c1v[0] && vv == c1v[1]) ||
								(k == c2v[0] && vv == c2v[1]) ||
								(k == c3v[0] && vv == c3v[1]) ||
								(vv == c1v[0] && k == c1v[1]) ||
								(vv == c2v[0] && k == c2v[1]) ||
								(vv == c3v[0] && k == c3v[1]) {
								continue
							}
							graph[k] = append(graph[k], vv)
						}
					}

					flatGraph := maps.Clone(nodes)

					island := 0
					islanSize := map[int]int{}
					for node, visited := range flatGraph {
						if !visited {
							island++
							traverseOnce(graph, node, func(currentNode string) {
								flatGraph[currentNode] = true
								islanSize[island]++
							})
						}
					}
					if len(islanSize) == 2 {
						fmt.Println(islanSize)
						return islanSize[1] * islanSize[2], 0
					}
				}

			}
		}
	}

	return 0, 0
}

func main() {
	start := time.Now()
	fmt.Println("Day 25: Snowverload")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) //
	fmt.Println("\tPart Two:", p2) //
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
