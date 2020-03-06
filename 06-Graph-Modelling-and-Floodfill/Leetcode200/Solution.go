package main

var (
	Grid  [][]byte
	num   int
	color byte
	H     int
	W     int
)

func numIslands(grid [][]byte) int {
	num = 0
	color = 2
	Grid = grid
	H = len(grid)
	if H == 0 {
		return 0
	}
	W = len(grid[0])
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if grid[i][j] == '1' {
				AreaOfIsland(i, j)
				num += 1
			}
		}
	}
	return num
}

func AreaOfIsland(i, j int) {
	Grid[i][j] = color
	if i != 0 {
		if Grid[i-1][j] == '1' {
			AreaOfIsland(i-1, j)
		}
	}
	if i < H-1 {
		if Grid[i+1][j] == '1' {
			AreaOfIsland(i+1, j)
		}
	}
	if j != 0 {
		if Grid[i][j-1] == '1' {
			AreaOfIsland(i, j-1)
		}
	}
	if j < W-1 {
		if Grid[i][j+1] == '1' {
			AreaOfIsland(i, j+1)
		}
	}
}

func main() {
	// 注意测试样例是 '1' 不是 1
}
