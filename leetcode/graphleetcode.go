package main

import (
	"fmt"
	"math"
)

// GRAPH problems
func shortestPathDijkstras() {
	// Graph represented as an adjacency list with weighted edges
	var graph = map[string]map[string]float64{
		"A": {"B": 2, "C": 4},
		"B": {"A": 2, "C": 1, "D": 7},
		"C": {"A": 4, "B": 1, "D": 3},
		"D": {"B": 7, "C": 3},
	}

	start := "A"
	shortest := make(map[string]float64)
	visited := make(map[string]bool)
	inf := math.Inf(1)

	for node := range graph {
		shortest[node] = inf
	}

	shortest[start] = 0

	for i := 0; i < len(graph); i++ {
		var current string
		minDist := inf
		for node, dist := range shortest {
			if !visited[node] && dist < minDist {
				minDist = dist
				current = node
			}
		}

		visited[current] = true

		for neighbor, weight := range graph[current] {
			if !visited[neighbor] {
				newDistance := shortest[current] + weight

				if newDistance < shortest[neighbor] {
					shortest[neighbor] = newDistance
				}
			}
		}
	}

	fmt.Println(shortest)
}

func largestComponent() {
	var graph = map[string][]string{
		"1": {"0"},
		"0": {"1", "5", "8"},
		"5": {"0", "8"},
		"8": {"0", "5"},
		"4": {"2", "3"},
		"2": {"4", "3"},
		"3": {"4", "2"},
	}

	visited := make(map[string]bool)
	count := 0

	for node := range graph {
		if !visited[node] {
			stack := []string{node}
			newCount := 0
			for len(stack) > 0 {
				current := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if !visited[current] {
					visited[current] = true
					newCount++
					for _, neighbor := range graph[current] {
						stack = append(stack, neighbor)
					}
				}
			}
			if newCount > count {
				count = newCount
			}
		}
	}

	fmt.Println(count)
}

// Note: Maybe revisit BFS algorithm
func NumberofIslands() {
	/*
		Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

		An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
		You may assume all four edges of the grid are all surrounded by water.

		Example 1:

		Input: grid = [
			["1","1","1","1","0"],
			["1","1","0","1","0"],
			["1","1","0","0","0"],
			["0","0","0","0","0"]
		]
		Output: 1

		Input: grid = [
			["1","1","0","0","0"],
			["1","1","0","0","0"],
			["0","0","1","0","0"],
			["0","0","0","1","1"]
		]
		Output: 3
	*/

	var grid = [][]string{
		{"1", "1", "1", "1", "0"},
		{"1", "1", "0", "1", "0"},
		{"1", "1", "0", "0", "0"},
		{"0", "0", "0", "0", "1"},
	}

	var island = 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {

			// this grid is 1(island) and hasn't been visited
			if grid[r][c] == "1" {
				// breadth first search on this cell
				island += 1
				bfs(grid, r, c)

			}
		}
	}
	fmt.Println(island)
}

func bfs(grid [][]string, r, c int) {
	q := [][]int{}

	q = append(q, []int{r, c})
	grid[r][c] = "2"

	// directions left, right, up, down
	offsets := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		// check left,right,up,down from the current gridValue
		for _, offset := range offsets {
			x := curr[0] + offset[0]
			y := curr[1] + offset[1]

			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) && grid[x][y] == "1" {
				q = append(q, []int{x, y})
				grid[x][y] = "2"
			}

		}
	}
}
