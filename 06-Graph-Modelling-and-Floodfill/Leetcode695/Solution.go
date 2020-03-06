package main

import "fmt"

var (
	Grid  [][]int
	max   int
	sum   int
	color int
	H     int
	W     int
)

func maxAreaOfIsland(grid [][]int) int {
	max = 0
	sum = 0
	color = 2
	Grid = grid
	H = len(grid)
	W = len(grid[0])
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if grid[i][j] == 1 {
				AreaOfIsland(i, j)
				if max < sum {
					max = sum
				}
				sum = 0
			}
		}
	}
	return max
}

func AreaOfIsland(i, j int) {
	sum += 1
	Grid[i][j] = color
	if i != 0 {
		if Grid[i-1][j] == 1 {
			AreaOfIsland(i-1, j)
		}
	}
	if i < H-1 {
		if Grid[i+1][j] == 1 {
			AreaOfIsland(i+1, j)
		}
	}
	if j != 0 {
		if Grid[i][j-1] == 1 {
			AreaOfIsland(i, j-1)
		}
	}
	if j < W-1 {
		if Grid[i][j+1] == 1 {
			AreaOfIsland(i, j+1)
		}
	}
}

func main() {
	grid := [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}
	grid1 := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	grid2 := [][]int{
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 1}}
	fmt.Println(maxAreaOfIsland(grid))
	fmt.Println(maxAreaOfIsland(grid1))
	fmt.Println(maxAreaOfIsland(grid2))
}
