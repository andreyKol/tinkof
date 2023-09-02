package main

import (
	"fmt"
	"sort"
)

func main() {
	var Revolvers, Money int

	if _, err := fmt.Scanln(&Revolvers, &Money); err != nil {
		fmt.Println("Пустые значения:", err)
		return
	}

	if (Revolvers < 1) || (Revolvers > 2e5) || (Money < 1) || (Money > 1e9) {
		fmt.Println("Неверные входные данные")
		return
	}

	revolversPrice := make([]int, Revolvers)

	for i := 0; i < Revolvers; i++ {
		fmt.Scan(&revolversPrice[i])
		if (revolversPrice[i]) < 1 || (revolversPrice[i] > 1e9) {
			fmt.Println("Неверные входные данные")
			return
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(revolversPrice)))

	for _, price := range revolversPrice {
		if price <= Money {
			fmt.Println(price)
			return
		}
	}

	fmt.Println(0)
}
