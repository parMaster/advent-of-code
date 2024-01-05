package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Vector struct {
	x, y, z, dx, dy, dz float64
}

func read(file string) []Vector {
	in, _ := os.ReadFile(file)

	v := []Vector{}
	for _, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		vc := Vector{}
		fmt.Sscanf(strings.TrimSpace(l), "%f, %f, %f @ %f, %f, %f", &vc.x, &vc.y, &vc.z, &vc.dx, &vc.dy, &vc.dz)
		v = append(v, vc)
	}

	return v
}

// returns intersection coordinates for two arbitrary lines or error
// result is float
func intersectAt(x1, y1, x2, y2, x3, y3, x4, y4 int) (float64, float64, error) {
	var m1, m2 float64

	if x2-x1 != 0 {
		m1 = float64(y2-y1) / float64(x2-x1)
	} else {
		m1 = math.MaxFloat64
	}

	if x4-x3 != 0 {
		m2 = float64(y4-y3) / float64(x4-x3)
	} else {
		m2 = math.MaxFloat64
	}

	if m1 == m2 {
		return 0, 0, fmt.Errorf("Lines are parallel or coincident")
	}

	b1 := float64(y1) - m1*float64(x1)
	b2 := float64(y3) - m2*float64(x3)

	xi := (b2 - b1) / (m1 - m2)
	yi := m1*float64(xi) + b1

	return xi, yi, nil
}

// returns true if {dx,dy} and {xi, yi} are on the same side from {x,y} -
// in other words - that vector is pointing to the intersection point, not from it
func futureIntersect(x, y, dx, dy, xi, yi float64) bool {
	if (xi-float64(x))*float64(dx) >= 0 && (yi-float64(y))*float64(dy) >= 0 {
		return true
	}
	return false
}

func insideArea(xi, yi, areaFrom, areaTo float64) bool {
	if xi >= float64(areaFrom) && xi <= float64(areaTo) && yi >= float64(areaFrom) && yi <= float64(areaTo) {
		return true
	}
	return false
}

func PartOne(file string, areaFrom, areaTo float64) int {
	v := read(file)
	crossings := 0
	for i, v1 := range v {
		for _, v2 := range v[i:] {

			xi, yi, err := intersectAt(int(v1.x), int(v1.y), int(v1.x+v1.dx), int(v1.y+v1.dy), int(v2.x), int(v2.y), int(v2.x+v2.dx), int(v2.y+v2.dy))
			if err != nil {
				continue
			}

			if !futureIntersect(v1.x, v1.y, v1.dx, v1.dy, xi, yi) ||
				!futureIntersect(v2.x, v2.y, v2.dx, v2.dy, xi, yi) {
				continue
			}

			if !insideArea(xi, yi, areaFrom, areaTo) {
				continue
			}

			crossings++
		}
	}

	return crossings
}

func main() {

	fmt.Println("Day 24: Never Tell Me The Odds")
	fmt.Println("\tPart One:", PartOne("../aoc-inputs/2023/24/input.txt", 200000000000000, 400000000000000)) // 15593

}
