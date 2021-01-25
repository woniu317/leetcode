package main

import "fmt"

func main() {
	// fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
	// fmt.Println(computeArea(-2, -2, 2, 2, 3, 3, 4, 4))
	fmt.Println(computeArea(-2, -2, 2, 2, -1, 4, 1, 6))
}
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	totalArea := area(A, B, C, D) + area(E, F, G, H)
	if !across(A, B, C, D, E, F, G, H) {
		return totalArea
	}
	xd, yd := down(A, B, E, F)
	xu, yu := up(C, D, G, H)
	acrossArea := area(xd, yd, xu, yu)
	return totalArea - acrossArea
}

func across(A int, B int, C int, D int, E int, F int, G int, H int) bool {
	// Y
	if B >= H || D <= F {
		return false
	}
	// X
	if A >= G || E >= C {
		return false
	}
	return true
}

func area(A, B, C, D int) int {
	fmt.Println(A, B, C, D)
	return (C - A) * (D - B)
}

func down(A, B, E, F int) (x, y int) {
	x = A
	if A < E {
		x = E
	}
	y = B
	if B < F {
		y = F
	}
	return
}

func up(C, D, G, H int) (x, y int) {
	x = C
	if C > G {
		x = G
	}
	y = D
	if D > H {
		y = H
	}
	return
}
