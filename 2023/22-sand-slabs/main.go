package main

import (
	"encoding/json"
	"fmt"
	"image"
	"os"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

type Brick struct {
	r      image.Rectangle
	h      int
	stands int
}

// Stack: map[Z(level)][]rectangles present at this level(1..MZ)
// so we can fit new rectangle to the level, or find the ones it intersects
type Stack map[int][]Brick

func read(in string) Stack {
	a := Stack{}
	for _, l := range strings.Fields(strings.TrimSpace(in)) {
		c := []int{}
		json.Unmarshal([]byte("["+strings.Replace(l, "~", ",", 1)+"]"), &c)
		// 0 1 2 3 4 5
		// 1,1,8~1,1,9
		r := image.Rect(c[0], c[1], c[3]+1, c[4]+1)
		a[c[2]] = append(a[c[2]], Brick{r, c[5] + 1 - c[2], 0})
	}
	return a
}

func drop(s Stack) (newStack Stack, changed int) {
	newStack = Stack{} // key is a brick's bottom coordinate
	tops := Stack{}    // key is a brick's top coordinate

	levels := maps.Keys(s)
	slices.Sort(levels)
	for level := 1; level <= levels[len(levels)-1]; level++ {
		// for all bricks that start on this level:
		for _, brick := range s[level] {

			newLevel := level
			levelFound := false
			for !levelFound && newLevel > 0 {
				for _, nlBrick := range tops[newLevel] {
					if !brick.r.Intersect(nlBrick.r).Empty() {
						// we`re here every time nlBrick is this bricks stand
						brick.stands++
						levelFound = true
					}
				}
				if !levelFound {
					newLevel--
				}
			}
			newLevel++

			if level != newLevel {
				changed++
			}
			// brick falls to stopLevel in a newStack (bottom cube)
			newStack[newLevel] = append(newStack[newLevel], brick)
			// same but key is a top cube coortinate
			tops[newLevel+(brick.h-1)] = append(tops[newLevel+(brick.h-1)], brick)
		}
	}

	return newStack, changed
}

func main() {
	input, _ := os.ReadFile("../aoc-inputs/2023/22/input.txt")
	stack, _ := drop(read(string(input)))
	p1 := 0
	p2 := 0
	for ilevel := range stack {
		for ibrick := range stack[ilevel] {
			checkStack := Stack{}

			// deep copy map of slices without one brick
			for il, ll := range stack {
				for ib, bb := range ll {
					if il == ilevel && ib == ibrick {
						continue
					}
					checkStack[il] = append(checkStack[il], bb)
				}
			}

			// will it change without one brick?
			_, changed := drop(checkStack)
			if changed == 0 {
				p1++
			}
			p2 += changed
		}
	}

	fmt.Println("Day 22: Sand Slabs")
	fmt.Println("\tPart One:", p1) // 439
	fmt.Println("\tPart Two:", p2) // 43056

}
