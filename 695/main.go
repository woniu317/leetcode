package main

func main() {

}

func maxAreaOfIsland(nums [][]int) int {
	ret := 0
	for i := range nums {
		for j := range nums[i] {
			tmp := 1
			if getNums(nums, i, j) == 1 {
				nums[i][j] = 0
				dfs(nums, i, j, &tmp)
			}
			if tmp > ret {
				ret = tmp
			}
		}
	}
	return ret
}

func dfs(grid [][]int, x, y int, area *int) {
	if getNums(grid, x-1, y) == 1 {
		*area++
		grid[x-1][y] = 0
		dfs(grid, x-1, y, area)
	}
	if getNums(grid, x+1, y) == 1 {
		*area++
		grid[x+1][y] = 0
		dfs(grid, x+1, y, area)
	}
	if getNums(grid, x, y-1) == 1 {
		*area++
		grid[x][y-1] = 0
		dfs(grid, x, y-1, area)
	}
	if getNums(grid, x, y+1) == 1 {
		*area++
		grid[x][y+1] = 0
		dfs(grid, x, y+1, area)
	}
}

func getNums(nums [][]int, x, y int) int {
	if x < 0 || x >= len(nums) || y < 0 || y >= len(nums[0]) {
		return 0
	}
	return nums[x][y]
}
