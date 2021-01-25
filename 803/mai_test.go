package main

import (
	"reflect"
	"testing"
)

func TestHitBricks(t *testing.T) {
	for i, tc := range []struct {
		grid [][]int
		hits [][]int
		ret  []int
	}{
		{[][]int{{1}, {1}, {1}, {1}, {1}}, [][]int{{3, 0}, {4, 0}, {1, 0}, {2, 0}, {0, 0}}, []int{1, 0, 1, 0, 0}},
		{[][]int{{1, 0, 1}, {1, 1, 1}}, [][]int{{0, 0}, {0, 2}, {1, 1}}, []int{0, 3, 0}},
	} {
		result := hitBricks(tc.grid, tc.hits)
		if !reflect.DeepEqual(result, tc.ret) {
			t.Errorf("No.%d's result error,ret=%v", i, result)
		}
	}
}

func hitBricks(grid [][]int, hits [][]int) []int {
	n := len(grid[0])

	stable := make([][]int, 0, len(grid))
	for range grid {
		stable = append(stable, make([]int, n))
	}

	ans := make([]int, len(hits))
	for k := range hits {
		if grid[hits[k][0]][hits[k][1]] == 0 {
			continue
		}

		// init stable
		for i := range grid {
			for j := range grid[i] {
				stable[i][j] = 0
			}
		}

		grid[hits[k][0]][hits[k][1]] = 0
		for i := range grid {
			if i == 0 {
				for j := range grid[i] {
					if grid[i][j] == 1 {
						stable[i][j] = 1
					}
				}
				continue
			}

			// 根据上、左标记当前稳定性
			for j := range grid[i] {
				if grid[i][j] != 1 {
					continue
				}

				// 上稳定
				if grid[i-1][j] == 1 {
					stable[i][j] = 1
					continue
				}

				if j > 0 && stable[i][j-1] == 1 {
					stable[i][j] = 1
					continue
				}
			}

			// 根据右侧标记稳定性
			for j := n - 1; j >= 0; j-- {
				if grid[i][j] != 1 || stable[i][j] == 1 {
					continue
				}

				if j < n-1 && stable[i][j+1] == 1 {
					stable[i][j] = 1
					continue
				}

				ans[k]++
				grid[i][j] = 0
			}
		}
	}

	return ans
}
