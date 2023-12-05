package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var games map[int]map[string]int

func PartOne(file string, r, g, b int) (sum int) {

	bag := map[string]int{
		"red":   r,
		"green": g,
		"blue":  b,
	}

	input, _ := os.ReadFile(file)

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, l := range lines {

		sum += func() int {
			game, _ := strconv.Atoi(strings.Split(l, ":")[0][5:])
			draws := strings.Split(strings.Split(l, ":")[1], ";")

			for _, draw := range draws {

				pairs := strings.Split(draw, ",")

				for _, pair := range pairs {
					pair = strings.TrimSpace(pair)
					p := strings.Split(pair, " ")
					number, _ := strconv.Atoi(p[0])
					color := p[1]
					if number > bag[color] {
						return 0
					}
				}
			}
			return game
		}()
	}

	return
}

func PartTwo(file string) (sum int) {
	input, _ := os.ReadFile(file)

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, l := range lines {

		sum += func() int {
			play := map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			}

			draws := strings.Split(strings.Split(l, ":")[1], ";")

			for _, draw := range draws {

				pairs := strings.Split(draw, ",")

				for _, pair := range pairs {
					pair = strings.TrimSpace(pair)
					p := strings.Split(pair, " ")
					number, _ := strconv.Atoi(p[0])
					color := p[1]
					play[color] = max(play[color], number)
				}
			}
			return play["red"] * play["green"] * play["blue"]
		}()
	}

	return
}

func main() {
	fmt.Println("Day 2: Cube Conundrum\n\tPart One:", PartOne("../aoc-inputs/2023/02/input1.txt", 12, 13, 14)) // 3059
	fmt.Println("\tPart Two:", PartTwo("../aoc-inputs/2023/02/input1.txt"))                                    //65371
}
