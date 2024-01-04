package main

import (
	"fmt"
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// returns nearest integer intersection coordinates for two arbitrary lines or error
func intersectPoint(line1, line2 [2]image.Point) (*image.Point, error) {
	x1, y1 := line1[0].X, line1[0].Y
	x2, y2 := line1[1].X, line1[1].Y
	x3, y3 := line2[0].X, line2[0].Y
	x4, y4 := line2[1].X, line2[1].Y

	xi, yi, err := intersectAt(x1, y1, x2, y2, x3, y3, x4, y4)

	return &image.Point{X: int(xi), Y: int(yi)}, err
}

func TestIntersect(t *testing.T) {
	// 19, 13, 30 @ -2,  1, -2
	// x1, y1, z1 @ dx1,dy1,dz1
	//
	// 18, 19, 22 @ -1, -1, -2
	// x2, y2, z2 @ dx2,dy2,dz2

	areaFrom := 7
	areaTo := 27

	// Hailstone A: 19, 13, 30 @ -2, 1, -2
	// Hailstone B: 18, 19, 22 @ -1, -1, -2
	// Hailstones' paths will cross inside the test area (at x=14.333, y=15.333).
	xi, yi, err := intersectAt(19, 13, 17, 14, 18, 19, 17, 18)
	require.Nil(t, err)
	require.EqualValues(t, [2]float64{14.333333333333334, 15.333333333333332}, [2]float64{xi, yi})

	// Check Inside or Outside of the test area
	assert.True(t, insideArea(xi, yi, float64(areaFrom), float64(areaTo)))

	intersect, err := intersectPoint([2]image.Point{{19, 13}, {17, 14}}, [2]image.Point{{18, 19}, {17, 18}})
	require.Nil(t, err)
	require.EqualValues(t, &image.Point{14, 15}, intersect)

	// Hailstone A: 19, 13, 30 @ -2, 1, -2
	// Hailstone B: 20, 25, 34 @ -2, -2, -4
	// Hailstones' paths will cross inside the test area (at x=11.667, y=16.667).
	xi, yi, err = intersectAt(19, 13, 17, 14, 20, 25, 18, 23)
	require.Nil(t, err)
	require.EqualValues(t, [2]float64{11.666666666666666, 16.666666666666668}, [2]float64{xi, yi})

	// Check Inside or Outside of the test area
	assert.True(t, insideArea(xi, yi, float64(areaFrom), float64(areaTo)))

	// Hailstone A: 18, 19, 22 @ -1, -1, -2
	// Hailstone B: 20, 25, 34 @ -2, -2, -4
	// Hailstones' paths are parallel; they never intersect.
	xi, yi, err = intersectAt(18, 19, 17, 18, 20, 25, 18, 23)
	require.Error(t, err)
	require.EqualError(t, fmt.Errorf("Lines are parallel or coincident"), err.Error())

	// Check Intersection is in the future for both lines

	// Hailstone A: 19, 13, 30 @ -2, 1, -2
	// Hailstone B: 20, 19, 15 @ 1, -5, -3
	// Hailstones' paths crossed in the past for hailstone A.
	x1, y1, dx, dy := 19, 13, -2, 1
	x2, y2, dx2, dy2 := 20, 19, 1, -5
	xi, yi, err = intersectAt(19, 13, 17, 14, 20, 19, 21, 14)
	require.Nil(t, err)
	require.Equal(t, [2]float64{21.444444444444443, 11.777777777777779}, [2]float64{xi, yi})

	// Either one of these is Negative? Then intersection was in the past
	assert.Negative(t, (xi-float64(x1))*float64(dx))
	assert.Negative(t, (yi-float64(y1))*float64(dy))

	// All Positive? Intersect is in the future
	assert.Positive(t, (xi-float64(x2))*float64(dx2))
	assert.Positive(t, (yi-float64(y2))*float64(dy2))

	// Future intersect for BOTH vectors
	assert.True(t, futureIntersect(float64(x2), float64(y2), float64(dx2), float64(dy2), xi, yi))
	assert.False(t, futureIntersect(float64(x1), float64(y1), float64(dx), float64(dy), xi, yi))

	// Hailstone A: 20, 25, 34 @ -2, -2, -4
	// Hailstone B: 12, 31, 28 @ -1, -2, -1
	// Hailstones' paths will cross outside the test area (at x=-2, y=3).
	x1, y1, dx1, dy1 := 20, 25, -2, -2
	x2, y2, dx2, dy2 = 12, 31, -1, -2
	xi, yi, err = intersectAt(x1, y1, x1+dx1, y1+dy1, x2, y2, x2+dx2, y2+dy2)
	assert.NoError(t, err)
	assert.Equal(t, [2]float64{-2, 3}, [2]float64{xi, yi})

	// Future intersect for BOTH vectors
	assert.True(t, futureIntersect(float64(x1), float64(y1), float64(dx1), float64(dy1), xi, yi))
	assert.True(t, futureIntersect(float64(x2), float64(y2), float64(dx2), float64(dy2), xi, yi))

	// Confirming "outside the test area (at x=-2, y=3)"
	assert.False(t, insideArea(xi, yi, float64(areaFrom), float64(areaTo)))

	// The whole strategy
	require.Equal(t, 2, PartOne("../aoc-inputs/2023/24/input0.txt", 7, 27))
	require.Equal(t, 15593, PartOne("../aoc-inputs/2023/24/input.txt", 200000000000000, 400000000000000))
}
