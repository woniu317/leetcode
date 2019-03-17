package main

import (
	"fmt"
)

//性能:area>area1
func main() {
	//Input: [1,8,6,2,5,4,8,3,7]
	//Output: 49
	fmt.Println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea1([]int{7, 3, 8, 4, 5, 2, 6, 8, 1}))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea2(height []int) int {
	lenth := len(height)
	maxArea, left := 0, 0
	right := lenth - 1
	for left < right {
		area := (right - left) * min(height[right], height[left])
		maxArea = max(area, maxArea)
		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func maxArea(height []int) int {
	lenth := len(height)
	maxArea := 0
	leftMaxHeigh := 0
	for i := 0; i < lenth; i++ {
		leftHigh := height[i]
		if leftHigh > leftMaxHeigh {
			leftMaxHeigh = leftHigh
			for j := i + 1; j < lenth; j++ {
				area := leftHigh * (j - i)
				if leftHigh > height[j] {
					area = height[j] * (j - i)
				}
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func maxArea1(height []int) int {
	lenth := len(height)
	maxArea := 0
	leftMaxHeigh := 0
	if computeFromLeft(height) {
		for i := 0; i < lenth; i++ {
			leftHigh := height[i]
			if leftHigh > leftMaxHeigh {
				leftMaxHeigh = leftHigh
				for j := i + 1; j < lenth; j++ {
					area := leftHigh * (j - i)
					if leftHigh > height[j] {
						area = height[j] * (j - i)
					}
					if area > maxArea {
						maxArea = area
					}
				}
			}
		}
	} else {
		for i := lenth - 1; i >= 0; i-- {
			leftHigh := height[i]
			if leftHigh > leftMaxHeigh {
				leftMaxHeigh = leftHigh
				for j := i - 1; j >= 0; j-- {
					area := leftHigh * (i - j)
					if leftHigh > height[j] {
						area = height[j] * (i - j)
					}
					if area > maxArea {
						maxArea = area
					}
				}
			}
		}
	}
	return maxArea
}

func computeFromLeft(height []int) bool {
	num := 0
	maxHeight := 0
	for _, value := range height {
		if value > maxHeight {
			num++
			maxHeight = value
		}
	}
	rightNum := 0
	maxHeight = 0
	for j := len(height) - 1; j >= 0; j-- {
		if height[j] > maxHeight {
			rightNum++
			maxHeight = height[j]
		}
	}
	return num >= rightNum
}
