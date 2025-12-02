package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func solve(file string) (p1, p2 int) {
	f, _ := os.ReadFile(file)
	for r := range strings.SplitSeq(string(f), ",") {

		limits := strings.Split(r, "-")
		lStart, _ := strconv.Atoi(limits[0])
		lEnd, _ := strconv.Atoi(limits[1])

		for n := lStart; n <= lEnd; n++ {
			l := int(math.Log10(float64(n))) + 1

			// p1
			factor := int(math.Pow(10, float64(l/2)))
			if n/factor == n%factor {
				p1 += n
			}

			// p2
			for lRange := 1; lRange <= l/2; lRange++ {
				if l%lRange != 0 {
					continue
				}

				factor := int(math.Pow(10, float64(lRange)))

				n2 := n
				sequence := n2 % factor
				repeats := 0
				for n2 != 0 {
					if sequence != n2%factor {
						repeats = 0
						break
					}
					repeats++
					n2 = n2 / factor
				}
				if repeats >= 2 {
					p2 += n
					// fmt.Println("P2+: n, sequence, repeats", n, sequence, (l / lRange))
					break
				}

			}
		}
	}
	return p1, p2
}

func main() {
	start := time.Now()
	fmt.Println("Day 02: Gift Shop")
	p1, p2 := solve("../aoc-inputs/2025/02/input.txt")
	// p1, p2 := solve("input-pub.txt")
	fmt.Println("\tPart One:", p1) // 40055209690
	fmt.Println("\tPart Two:", p2) // 50857215650
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
