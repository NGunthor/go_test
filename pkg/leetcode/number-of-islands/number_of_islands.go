package number_of_islands

// NumsIslands is a main algorithm function
func NumIslands(grid [][]byte) int {
	count := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 49 {
				count++
				alg(x, y, grid)
			}
		}
	}
	return count
}

func alg(x int, y int, grid [][]byte) {
	h := len(grid)
	l := len(grid[0])
	if grid[y][x] == 49 {
		grid[y][x] = 48
		if x + 1 < l {
			alg(x + 1, y, grid)
		}
		if x - 1 > -1 {
			alg(x - 1, y, grid)
		}
		if y + 1 < h {
			alg(x, y + 1, grid)
		}
		if y - 1 > -1 {
			alg(x, y - 1, grid)
		}
	}
}
