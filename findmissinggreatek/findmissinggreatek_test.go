package findmissinggreatek

import "testing"

// 要求，在一个无序数组中arr中，找出比k大的最小的不在arr中存在的整数

func TestFindMissing(t *testing.T) {
	for i, tc := range []struct {
		target int
		arr    []int
		k      int
	}{
		{
			target: 9,
			arr:    []int{8, 7, 6, 10, 11, 32, 12, 13},
			k:      5,
		},
	} {
		ret := findMissingLatestK(tc.arr, tc.k)
		if tc.target != ret {
			t.Errorf("test(NO.%d)=%d,want matches=%d", i, ret, tc.target)
			continue
		}
	}
}

func findMissingLatestK(arr []int, k int) int {
	for i := range arr {
		resetPositionOnce(i, arr, k)
	}

	return fetchLatestK(arr, k)
}

func fetchLatestK(arr []int, k int) int {
	ret := k + len(arr) + 1

	for i := range arr {
		if arr[i] != k+i+1 {
			ret = k + i + 1
			break
		}
	}

	return ret
}

func resetPositionOnce(i int, arr []int, target int) {
	k := arr[i]
	for k > target && k <= target+len(arr) {
		p := k - target - 1
		if arr[p] == k {
			return
		}
		arr[p], k = k, arr[p]
	}
}
