package main

import (
	"image"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	in := `1,0,1~1,2,1
	0,0,2~2,0,2
	0,2,3~2,2,3
	0,0,4~0,2,4
	2,0,5~2,2,5
	0,1,6~2,1,6
	1,1,8~1,1,9
	`
	stack := read(in)

	require.NotEmpty(t, stack)
	require.Equal(t, image.Rect(1, 0, 2, 3), stack[1][0].r)
	require.Equal(t, image.Rect(1, 1, 2, 2), stack[8][0].r)

	// Intersection is not empty
	require.False(t, stack[1][0].r.Intersect(stack[8][0].r).Empty())

	// it's 1,1-2,2
	require.Equal(t, image.Rect(1, 1, 2, 2), stack[1][0].r.Intersect(stack[8][0].r))
}

func TestDrop(t *testing.T) {
	in := `1,0,1~1,2,1
	0,0,2~2,0,2
	0,2,3~2,2,3
	0,0,4~0,2,4
	2,0,5~2,2,5
	0,1,6~2,1,6
	1,1,8~1,1,9
	`
	stack := read(in)
	stack, changed := drop(stack)

	assert.Equal(t, 5, changed)
	assert.Equal(t, 1, len(stack[5]))
	assert.Equal(t, 0, len(stack[6]))

	// t.SkipNow()

	stack, changed = drop(stack)
	require.Zero(t, changed, "the same stack shouldn't change")
	stack, changed = drop(stack)
	require.Zero(t, changed, "the same stack shouldn't change")
	stack, changed = drop(stack)
	require.Zero(t, changed, "the same stack shouldn't change")

	input, _ := os.ReadFile("input.txt")
	stack, _ = drop(read(string(input)))
	nodiff := 0
	diff := 0
	for ilevel := range stack {
		log.Println(ilevel)
		stack, changed = drop(stack)
		require.Zero(t, changed, "the same stack shouldn't change")

		for ibrick := range stack[ilevel] {
			log.Println(ibrick)

			checkStack := Stack{}
			for il, ll := range stack {
				for ib, bb := range ll {
					if il == ilevel && ib == ibrick {
						continue
					}
					checkStack[il] = append(checkStack[il], bb)
				}
			}

			require.Equal(t, 1, len(stack[ilevel])-len(checkStack[ilevel]))

			_, changed := drop(checkStack)
			if changed == 0 {
				nodiff++
			}
			diff += changed

		}
	}

	require.Equal(t, 439, nodiff)
	require.Equal(t, 43056, diff)
}
