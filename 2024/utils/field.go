package utils

type Field [][]rune

var dir = [][2]int{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}

// recursive search around for term
func (f Field) search(term string, y, x int) (res int) {

	r := rune(term[0])
	if y < 0 || y >= len(f) || x < 0 || x >= len(f[y]) {
		return
	}

	if f[y][x] == r {
		if len(term) == 1 {
			return 1
		}
		for _, d := range dir {
			ny := y + d[0]
			nx := x + d[1]
			res += f.search(term[1:], ny, nx)
		}
	}

	return res
}
