package numIslands

func numIslands(grid [][]byte) int {
	mm, n := len(grid), len(grid[0])
	visit := make([][]bool, mm)
	for i := range visit {
		visit[i] = make([]bool, n)
	}
	count := 0
	for i  := range grid {
		for j  := range grid[i] {
			if grid[i][j] == '1' && !visit[i][j]{
				help(grid, i, j ,visit)
				count++
			}
		}
	}
	return count
}

func help(grid [][]byte, i, j int, visit [][]bool) {
	if i >= len(grid) || i < 0 || j >= len(grid[0]) || j < 0 || grid[i][j] == '0' {
		return
	}
	if visit[i][j] {
		return
	}
	visit[i][j] = true
	help(grid, i+1, j+1, visit)
	help(grid, i, j+1, visit)
	help(grid, i+1, j, visit)
	help(grid, i, j, visit)

}
