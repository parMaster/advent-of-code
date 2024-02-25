package main

import (
	"encoding/json"
	"fmt"
	"image"
	"os"
	"strings"
	"time"
)

func readInput(filename string) ([]image.Point, []image.Point, int, int) {
	data, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	hvPoints := []image.Point{}
	dPoints := []image.Point{}
	w, h := 0, 0
	for _, line := range lines {
		c := []int{}
		json.Unmarshal([]byte("["+strings.Replace(line, " -> ", ",", 1)+"]"), &c)
		// 0 1    2 3
		// 0,9 -> 5,9
		w = max(w, c[2], c[0])
		h = max(h, c[3], c[1])

		// diagonal lines
		if c[0] != c[2] && c[1] != c[3] {
			length := max(c[2], c[0]) - min(c[2], c[0])
			for i := 0; i <= length; i++ {
				dx := 1
				if c[0] > c[2] {
					dx = -1
				}
				dy := 1
				if c[1] > c[3] {
					dy = -1
				}
				dPoints = append(dPoints, image.Pt(c[0]+i*dx, c[1]+i*dy))
			}
			continue
		}

		// horizontal or vertical lines
		if c[0] == c[2] {
			for i := min(c[1], c[3]); i <= max(c[1], c[3]); i++ {
				hvPoints = append(hvPoints, image.Pt(c[0], i))
			}
		}

		if c[1] == c[3] {
			for i := min(c[0], c[2]); i <= max(c[0], c[2]); i++ {
				hvPoints = append(hvPoints, image.Pt(i, c[1]))
			}
		}
	}

	return hvPoints, dPoints, w, h
}

func solve(filename string) (int, int) {

	hvPoints, dPoints, _, _ := readInput(filename)

	pm := map[image.Point]int{}
	for _, p := range hvPoints {
		pm[p] += 1
	}

	hvOverlaps := 0
	for _, v := range pm {
		if v > 1 {
			hvOverlaps++
		}
	}

	for _, p := range dPoints {
		pm[p] += 1
	}

	dOverlaps := 0
	for _, v := range pm {
		if v > 1 {
			dOverlaps++
		}
	}

	return hvOverlaps, dOverlaps
}

func main() {
	start := time.Now()
	fmt.Println("Day 5: Hydrothermal Venture")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 6687
	fmt.Println("\tPart Two:", p2) // 19851
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
