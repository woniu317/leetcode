package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 {
		return -1
	}
	for i := 0; i < len(gas); i++ {
		if testCircuit(i, gas, cost) {
			return i
		}
	}
	return -1
}

func testCircuit(i int, gas, cost []int) bool {
	start := i
	over := false
	nowGas := 0
	for !over || start != i {
		over = true
		nowGas += gas[start]
		nowGas -= cost[start]
		fmt.Println(">", start, i, nowGas)
		if nowGas < 0 {
			return false
		}
		start = (start + 1) % len(gas)
	}
	return true
}
