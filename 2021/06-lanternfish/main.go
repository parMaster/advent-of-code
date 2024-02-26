package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func solve(filename string) (int64, int64) {
	data, _ := os.ReadFile(filename)
	fishStr := strings.Split(strings.TrimSpace(string(data)), ",")
	fish := make(map[int]int64, 9)
	for _, f := range fishStr {
		fi := int(f[0] - 0x30)
		if _, ok := fish[fi]; !ok {
			fish[fi] = 1
		} else {
			fish[fi]++
		}
	}

	day80Total := int64(0)
	day256Total := int64(0)
	for day := range 256 {
		spawnFish := fish[0]
		fish[0] = 0
		for i := 1; i < 9; i++ {
			fish[i-1] = fish[i]
		}
		fish[8] = spawnFish
		fish[6] += spawnFish
		if day == 79 {
			for _, f := range fish {
				day80Total += f
			}
		}
	}

	for _, f := range fish {
		day256Total += f
	}

	return day80Total, day256Total
}

func main() {
	start := time.Now()
	fmt.Println("Day 6: Lanternfish")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 374927
	fmt.Println("\tPart Two:", p2) // 1687617803407
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
