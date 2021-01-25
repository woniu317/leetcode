package main

func main() {

}

func movesToMakeZigzag(nums []int) int {
	length := len(nums)
	rets := []int{0, 0}

	for i := range rets {
		for j := i; j < length; j += 2 {
			d := 0

			if j > 0 {
				d = max(d, nums[j]-nums[j-1]+1)
			}

			if j+1 < length {
				d = max(d, nums[j]-nums[j+1]+1)
			}

			rets[i] += d
		}
	}

	if rets[0] > rets[1] {
		return rets[1]
	}

	return rets[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
