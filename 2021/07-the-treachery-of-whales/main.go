package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func solve(filename string) (int, int) {
	data, _ := os.ReadFile(filename)
	crabsStr := strings.Split(strings.TrimSpace(string(data)), ",")
	crabs := []int{}
	minpos, maxpos := int(^uint(0)>>1), 0
	for _, c := range crabsStr {
		crabPos, _ := strconv.Atoi(c)
		crabs = append(crabs, crabPos)
		minpos = min(minpos, crabPos)
		maxpos = max(maxpos, crabPos)
	}

	minFuel := maxpos*len(crabs) + 1
	minFuel2 := int(^uint(0) >> 1)
	for i := 0; i < maxpos; i++ {
		fuel := 0
		fuel2 := 0

		for _, c := range crabs {
			fuelRequired := i - c
			if fuelRequired < 0 {
				fuelRequired = -fuelRequired
			}
			fuel += fuelRequired
			// sum of an arithmetic progression
			// Sn = n/2[2a + (n − 1) × d]
			fuel2 += int(float64(fuelRequired) / 2 * float64(2+(fuelRequired-1)*1))
		}
		if fuel < minFuel {
			minFuel = fuel
		}
		if fuel2 < minFuel2 {
			minFuel2 = fuel2
		}
	}
	return minFuel, minFuel2
}

func main() {
	start := time.Now()
	fmt.Println("Day 7: The Treachery of Whales")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 343605
	fmt.Println("\tPart Two:", p2) // 96744904
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
