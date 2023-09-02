package main

import (
	"fmt"
	"sort"
)

func main() {

	var sum, countBill int
	fmt.Scan(&sum, &countBill)

	bills := make([]int, countBill)
	for i := 0; i < countBill; i++ {
		fmt.Scan(&bills[i])
	}

	sort.Ints(bills)

	sumArray := make([]int, sum+1)
	for i := range sumArray {
		sumArray[i] = -1
	}

	sumArray[0] = 0

	for i := 0; i < countBill; i++ {
		for j := bills[i]; j <= sum; j++ {
			if sumArray[j-bills[i]] != -1 {
				if sumArray[j] == -1 || sumArray[j] > sumArray[j-bills[i]]+1 {
					sumArray[j] = sumArray[j-bills[i]] + 1
				}
			}
		}
	}

	if sumArray[sum] == -1 {
		fmt.Println(-1)
		return
	}

	result := []int{}
	k := sum
	for k > 0 {
		for i := 0; i < countBill; i++ {
			if k >= bills[i] && sumArray[k-bills[i]] == sumArray[k]-1 {
				result = append(result, bills[i])
				k -= bills[i]
				break
			}
		}
	}

	fmt.Println(len(result))
	for i := 0; i < len(result); i++ {
		fmt.Printf("%d ", result[i])
	}
}
