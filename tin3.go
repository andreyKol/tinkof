package main

import (
	"fmt"
	"sort"
)

func main() {
	var cards int
	fmt.Scan(&cards)
	values1 := make([]int, cards)
	values2 := make([]int, cards)

	for i := 0; i < cards; i++ {
		fmt.Scan(&values1[i])
	}
	for i := 0; i < cards; i++ {
		fmt.Scan(&values2[i])
	}

	result := checkWin(cards, values1, values2)
	fmt.Println(result)
}

func checkWin(n int, a, b []int) string {
	diffIndex := -1
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			diffIndex = i
			break
		}
	}

	if diffIndex == -1 {
		return "YES"
	}

	for i := n - 1; i >= 0; i-- {
		if a[i] != b[i] {
			newArray := append([]int{}, a[diffIndex:i+1]...)
			sort.Ints(newArray)
			copy(a[diffIndex:i+1], newArray)

			if comparingArrays(a, b) {
				return "YES"
			} else {
				return "NO"
			}
		}
	}
	return "NO"
}

func comparingArrays(slice1, slice2 []int) bool {

	if len(slice1) != len(slice2) {
		return false
	}

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
