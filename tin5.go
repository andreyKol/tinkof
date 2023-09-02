package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var n, m int

type edgeType struct {
	u, v, w int
}

var edges []edgeType
var parent []int

func findParent(x int) int {
	if parent[x] != x {
		parent[x] = findParent(parent[x])
	}
	return parent[x]
}

func merge(x, y int) {
	px, py := findParent(x), findParent(y)
	if px != py {
		parent[px] = py
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < m; i++ {
		scanner.Scan()
		u, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		w, _ := strconv.Atoi(scanner.Text())
		edges = append(edges, edgeType{u, v, w})
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	parent = make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	totalCount := 0

	for _, edge := range edges {
		if findParent(edge.u) != findParent(edge.v) {
			totalCount += edge.w
			merge(edge.u, edge.v)
		}
	}

	fmt.Println(totalCount)
}
