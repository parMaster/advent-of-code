package main

import (
	"encoding/json"
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
	"time"
)

func solve(filename string) (int, string) {

	paper := map[image.Point]bool{}

	text, _ := os.ReadFile(filename)
	parts := strings.Split(strings.TrimSpace(string(text)), "\n\n")
	points := strings.Split(strings.TrimSpace(string(parts[0])), "\n")
	folds := strings.Split(strings.TrimSpace(string(parts[1])), "\n")

	w, h := 0, 0
	for _, point := range points {
		n := []int{}
		json.Unmarshal([]byte("["+point+"]"), &n)
		paper[image.Pt(n[0], n[1])] = true
		w = max(w, n[0])
		h = max(h, n[1])
	}

	if w%2 != 0 {
		w++
	}
	if h%2 != 0 {
		h++
	}

	// fmt.Println("Paper size:", w, h, "Points:", len(paper))

	p1 := 0

	for i, foldPos := range folds {
		f := strings.Fields(foldPos)[2]
		fold, _ := strconv.Atoi(f[2:])

		if f[0] == 'x' {
			for x := 0; x < fold; x++ {
				for y := 0; y <= h; y++ {

					if _, ok := paper[image.Pt(w-x, y)]; ok {
						paper[image.Pt(x, y)] = true
					}

				}
			}
			w = fold - 1

		} else {
			for x := 0; x <= w; x++ {
				for y := 0; y < fold; y++ {

					if _, ok := paper[image.Pt(x, h-y)]; ok {
						paper[image.Pt(x, y)] = true
					}

				}
			}
			h = fold - 1
		}

		if i == 0 {
			for x := 0; x <= w; x++ {
				for y := 0; y <= h; y++ {
					if _, ok := paper[image.Pt(x, y)]; ok {
						p1++
					}
				}
			}
		}
	}

	p2 := render(paper, w, h)

	return p1, p2
}

func render(paper map[image.Point]bool, w, h int) string {
	res := "\n"
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			if _, ok := paper[image.Pt(x, y)]; ok {
				res += "#"
			} else {
				res += "."
			}
		}
		res += "\n"
	}
	return res
}

func main() {
	start := time.Now()
	fmt.Println("Day 13: Transparent Origami")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 763
	fmt.Println("\tPart Two:", p2) // RHALRCRA
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
