package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadReach(t *testing.T) {

	in, _ := os.ReadFile("../aoc-inputs/2023/21/input0.txt")
	lines, sx, sy := readLines(string(in), 0)

	// showLines(lines, sx, sy)

	require.Equal(t, 11, len(lines))
	require.Equal(t, 5, sx)
	require.Equal(t, 5, sy)

	require.Equal(t, 16, reach(lines, sx, sy, 6)[6])

	in, _ = os.ReadFile("../aoc-inputs/2023/21/input0.txt")
	lines, sx, sy = readLines(string(in), 3) // 3x3 tiles
	// showLines(lines, sx, sy)
	require.Equal(t, 33, len(lines))
	require.Equal(t, 16, sx)
	require.Equal(t, 16, sy)

	in, _ = os.ReadFile("../aoc-inputs/2023/21/input.txt")
	lines, sx, sy = readLines(string(in), 3) // 3x3 tiles
	// showLines(lines, sx, sy)

	in, _ = os.ReadFile("../aoc-inputs/2023/21/input.txt")
	lines, sx, sy = readLines(string(in), 0)
	require.Equal(t, 131, len(lines))
	require.Equal(t, 131, len(lines[0]))
	require.Equal(t, 65, sx)
	require.Equal(t, 65, sy)
	// correctness on a single tile
	require.Equal(t, 3646, reach(lines, sx, sy, 64)[64])

	lines, sx, sy = readLines(string(in), 5) // 5x5 tiles
	require.Equal(t, 131*5, len(lines))
	require.Equal(t, 131*5/2, sx)
	require.Equal(t, 131*5/2, sy)
	// correctness on a tiled plane
	require.Equal(t, 3646, reach(lines, sx, sy, 64)[64])
}

func TestPredict(t *testing.T) {

	a := []int{
		3759,  // 0*131+65
		33496, // 1*131+65
		92857, // 2*131+65
	}

	// a[3*131+65] = predict[a]

	// how many garden plots could the Elf reach in exactly 26501365 steps?
	// 26501365 = 202300 * 131 + 65

	// log.Println(predict(a))

	for N := 2; N < 202300; N++ {
		a = append(a, predict(a[max(0, len(a)-4):]))
		if N%1000 == 0 || N > 202290 {
			log.Println(N, a[len(a)-1])
		}
	}

	log.Println(a[len(a)-1])
}

// 2023/12/25 00:39:32 2 181842
// 2023/12/25 00:39:32 3 300451
// 2023/12/25 00:39:32 4 448684
// 2023/12/25 00:39:32 5 626541
// 2023/12/25 00:39:32 6 834022
// 2023/12/25 00:39:32 7 1071127
// 2023/12/25 00:39:32 8 1337856
// 2023/12/25 00:39:32 9 1634209
