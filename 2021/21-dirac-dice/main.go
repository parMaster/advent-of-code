package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func roll3sum() func() int {
	roll := 0
	return func() int {
		sum := 0
		for range 3 {
			roll++
			if roll > 100 {
				roll = 1
			}
			fmt.Print(roll, " ")
			sum += roll
		}
		fmt.Println()
		return sum
	}
}

type Player struct {
	place int
	score int
}

func solve(file string) (p1 int, p2 int) {
	input, _ := os.ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	// fmt.Println(lines)
	players := []Player{}
	for _, line := range lines {
		parts := strings.Split(line, ":")
		place, _ := strconv.Atoi(strings.TrimSpace(parts[len(parts)-1]))
		players = append(players, Player{place: place, score: 0})
	}
	fmt.Println(players)
	roll := roll3sum()
	rolls := 0
	for rolls < 2000 {
		for i, player := range players {
			r := roll()
			rolls += 3
			fmt.Println("Rolls", rolls)
			fmt.Print("Player ", i, " p: ", player.place, " s: ", player.score, " rolls ", r)

			player.place += r
			if player.place > 10 {
				player.place = player.place % 10
				if player.place == 0 {
					player.place = 10
				}
			}
			player.score += player.place
			players[i] = player
			fmt.Println(" score", players[i].score)

			if player.score >= 1000 {
				if i == 0 {
					p1 = players[1].score * rolls
				} else {
					p1 = players[0].score * rolls
				}
				break
			}
		}
		if p1 > 0 {
			break
		}
	}

	return p1, 0
}

func main() {
	start := time.Now()
	fmt.Println("Day 21: Dirac Dice")
	p1, p2 := solve("input.txt")
	fmt.Println("\tPart One:", p1) // 598416
	fmt.Println("\tPart Two:", p2) //
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}
