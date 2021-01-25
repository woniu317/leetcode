package main

import "sort"

type SnapshotArray struct {
	id   int
	snap map[int]map[int]int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{0, make(map[int]map[int]int, 0)}
}

func (ssArray *SnapshotArray) Set(index int, val int) {
	if ssArray.snap[index] == nil {
		ssArray.snap[index] = make(map[int]int, 0)
	}
	// 记录下一次快照前，改变过的index值
	ssArray.snap[index][ssArray.id] = val
}

func (ssArray *SnapshotArray) Snap() int {
	ssArray.id++
	// 返回上一次的snap_id
	return ssArray.id - 1
}

/*如果index当前snap_id不存在，则说明snap_id这次快照跟snap_id-1这次快照之间index的值没变化过
二分找到snpa_id前第一个存在的快照id
但是go没有unorderedmap，还要枚举allkeys排序再二分,还不如直接遍历O(N)
*/
func (ssArray *SnapshotArray) Get(index int, snap_id int) int {
	if value, ok := ssArray.snap[index][snap_id]; ok {
		return value
	}
	// 枚举所有key
	allkeys := make([]int, len(ssArray.snap[index]))
	i := 0
	for key, _ := range ssArray.snap[index] {
		allkeys[i] = key
		i++
	}
	//排序已备二分
	sort.Ints(allkeys)

	low := 0
	high := len(allkeys) - 1
	for low <= high {
		mid := (low + high) / 2
		if allkeys[mid] < snap_id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	//如果low==0，说明不存在这样的快照，返回初始值0
	if low == 0 {
		return 0
	} else {
		//找到第一个比目标值大的下标，需要-1
		return ssArray.snap[index][allkeys[low-1]]
	}
}

func (ssArray *SnapshotArray) Get2(index int, snap_id int) int {
	for snap_id >= 0 {
		if v, ok := ssArray.snap[index][snap_id]; ok {
			return v
		} else {
			snap_id--
		}
	}
	return 0
}
