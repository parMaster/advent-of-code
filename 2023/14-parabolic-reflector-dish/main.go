package main

import (
	"fmt"
	"image"
	"os"
	"slices"
	"strings"
)

type field map[image.Point]rune

// directions	↑ ← ↓ →
// directions	0 1 2 3
var xyDir = []image.Point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

func read(in string) (m field, w int, h int) {
	m = make(map[image.Point]rune)
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		w = len(l) - 1
		h = y
		for x, r := range strings.TrimSpace(l) {
			m[image.Point{x, y}] = r
		}
	}
	return
}

func cycle(field field, w, h int) field {
	var dots, rocks int
	// fmt.Println("North")
	for x := 0; x <= w; x++ {
		// down each column
		cur := image.Pt(x, 0)
		for y := 0; y <= h; y++ {
			if field[image.Point{x, y}] == '.' {
				dots++
			} else if field[image.Point{x, y}] == 'O' {
				rocks++
			}
			if field[image.Point{x, y}] == '#' || y == h {
				for i := 0; i < rocks; i++ {
					field[cur] = 'O'
					cur = cur.Add(xyDir[2])
				}
				for i := 0; i < dots; i++ {
					field[cur] = '.'
					cur = cur.Add(xyDir[2])
				}
				cur = cur.Add(xyDir[2])

				dots, rocks = 0, 0
			}
		}
	}
	// show(field, w, h)

	// fmt.Println("West")
	for y := 0; y <= h; y++ {
		cur := image.Pt(0, y)
		// right each row
		for x := 0; x <= w; x++ {
			if field[image.Point{x, y}] == '.' {
				dots++
			} else if field[image.Point{x, y}] == 'O' {
				rocks++
			}
			if field[image.Point{x, y}] == '#' || x == w {
				for i := 0; i < rocks; i++ {
					field[cur] = 'O'
					cur = cur.Add(xyDir[3])
				}
				for i := 0; i < dots; i++ {
					field[cur] = '.'
					cur = cur.Add(xyDir[3])
				}
				cur = cur.Add(xyDir[3])

				dots, rocks = 0, 0
			}
		}
	}
	// show(field, w, h)
	// fmt.Println("South")
	for x := 0; x <= w; x++ {
		// down each column
		cur := image.Pt(x, 0)
		for y := 0; y <= h; y++ {
			if field[image.Point{x, y}] == '.' {
				dots++
			} else if field[image.Point{x, y}] == 'O' {
				rocks++
			}
			if field[image.Point{x, y}] == '#' || y == h {
				for i := 0; i < dots; i++ {
					field[cur] = '.'
					cur = cur.Add(xyDir[2])
				}
				for i := 0; i < rocks; i++ {
					field[cur] = 'O'
					cur = cur.Add(xyDir[2])
				}
				cur = cur.Add(xyDir[2])

				dots, rocks = 0, 0
			}
		}
	}
	// show(field, w, h)
	// fmt.Println("East")
	for y := 0; y <= h; y++ {
		// right each row
		cur := image.Pt(0, y)
		for x := 0; x <= w; x++ {
			if field[image.Point{x, y}] == '.' {
				dots++
			} else if field[image.Point{x, y}] == 'O' {
				rocks++
			}
			if field[image.Point{x, y}] == '#' || x == w {
				for i := 0; i < dots; i++ {
					field[cur] = '.'
					cur = cur.Add(xyDir[3])
				}
				for i := 0; i < rocks; i++ {
					field[cur] = 'O'
					cur = cur.Add(xyDir[3])
				}
				cur = cur.Add(xyDir[3])

				dots, rocks = 0, 0
			}
		}
	}
	return field
}

func totalLoad(field field, w, h int) (sum int) {
	for x := 0; x <= w; x++ {
		for y := 0; y <= h; y++ {
			if field[image.Point{x, y}] == 'O' {
				sum += h - y + 1
			}
		}
	}
	return
}

// image -> string
func flatten(f field, w, h int) (s string) {
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			s += string(f[image.Pt(x, y)])
		}
	}
	return
}

func p2(f string) int {
	in, _ := os.ReadFile(f)
	field, w, h := read(string(in))
	// show(field, w, h)
	var cache []string

	cycleEnd := 0
	cycleStart := 0
	for {
		field := cycle(field, w, h)
		flat := flatten(field, w, h)
		if slices.Contains(cache, flat) {
			cycleStart = slices.Index(cache, flat)
			break
		}
		cache = append(cache, flat)
		cycleEnd++
	}

	// fmt.Println("Cycle found after", cycleEnd, "iterations - ", cycleStart, "was the same")
	// fmt.Println("1'000'000'000-", cycleStart, "%", cycleEnd, "-", cycleStart, " = ", (1e9-cycleStart)/(cycleEnd-cycleStart))

	field, w, h = read(string(in))
	for i := 0; i < 1e9-((1e9-cycleStart)/(cycleEnd-cycleStart))*(cycleEnd-cycleStart); i++ { // skipping as many cycles as we can
		field = cycle(field, w, h)
	}

	// show(field, w, h)
	load := totalLoad(field, w, h)
	return load
}

func p1(f string) int {
	in, _ := os.ReadFile(f)
	field, w, h := read(string(in))

	var dots, rocks int
	// fmt.Println("North")
	for x := 0; x <= w; x++ {
		// down every column
		cur := image.Pt(x, 0)
		for y := 0; y <= h; y++ {
			if field[image.Point{x, y}] == '.' {
				dots++
			} else if field[image.Point{x, y}] == 'O' {
				rocks++
			}
			if field[image.Point{x, y}] == '#' || y == h {
				for i := 0; i < rocks; i++ {
					field[cur] = 'O'
					cur = cur.Add(xyDir[2])
				}
				for i := 0; i < dots; i++ {
					field[cur] = '.'
					cur = cur.Add(xyDir[2])
				}
				cur = cur.Add(xyDir[2])

				dots, rocks = 0, 0
			}
		}
	}

	return totalLoad(field, w, h)
}

func main() {
	fmt.Println("Day 14: Parabolic Reflector Dish")
	fmt.Println("\tPart One:", p1("../aoc-inputs/2023/14/input.txt")) // 108840
	fmt.Println("\tPart Two:", p2("../aoc-inputs/2023/14/input.txt")) // 103445
}

func show(m map[image.Point]rune, w, h int) {
	fmt.Println(w+1, "x", h+1, ":")
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			fmt.Print(string(m[image.Pt(x, y)]))
		}
		fmt.Println()
	}
	fmt.Println()
}
