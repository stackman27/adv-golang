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

    for i:=0; i < len(graph); i++ {
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

// Note: assuming all edges have same length of 1 
func shortestPath() {
	// Perform BFS (DFS will traverse unnecessary nodes)
	// Graph represented as an adjacency list with uniform edge lengths
	var graph = map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "D"},
		"C": {"A", "E"},
		"D": {"B", "E"},
	}

	visited := make(map[string]bool)
	queue := [][]string{{"A"}}
	//start := "A"
	end := "E"

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		node := path[len(path) - 1]
		fmt.Println(path)
		if node == end {
			fmt.Println(path)
			return  
		}

		if !visited[node] {
			visited[node] = true
			for _, neighbor := range graph[node] {
				// create a copy of the current explored path
				// This is done to ensure that we don't modify the current path when exploring neighbors.
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
  
				queue = append(queue, newPath) 
			}

		}


	}
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
					newCount ++ 
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

func connectedComponentsCount() {
	var graph = map[string][]string{
		"a": {"b", "c"},
		"b": {"a", "d"},
		"c": {"a", "e"},
		"d": {"b", "f"},
		"e": {"c"},
		"f": {"d"},
		"g": {"h"},
		"h": {"g"},
		"i": {},
	}

	visited := make(map[string]bool)
    count := 0

    for node := range graph {
		stack := []string{node}

        if !visited[node] {
       
            for len(stack) > 0 {
                current := stack[len(stack)-1]
                stack = stack[:len(stack)-1]

                if !visited[current] {
                    visited[current] = true
                    for _, neighbor := range graph[current] {
                        stack = append(stack, neighbor)
                    }
                }
            }

            count++
        }
    }

   fmt.Println(count)
}

// Note: This is for undirected graph meaning there is edges both ways 
func buildGraphGivenEdges() {
	edges := [][]string{
		 {"i", "j"},
		 {"k", "i"},
		 {"m", "k"},
		 {"k", "l"},
		 {"o", "n"},
	}

	graph := make(map[string][]string)

	for _, edge := range edges {
		a := edge[0]
		b := edge[1]

		_, ok := graph[a]
		if !ok {
			graph[a] = []string{}
		} 

		_, ok = graph[b]
		if !ok {
			graph[b] = []string{}
		} 

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
 
	}

	fmt.Println(graph)
	
}

// NOTE: Make sure the graph is acyclic (meaning no cycle)
func hasPath() {
	// DFS solution 
	var graph = map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}	

	source := "b"
	destination := "e"

	stack := []string{source} 	
	visited := make(map[string]bool)
	pathExist := false

	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if node == destination {
			pathExist = true
			break;  
		} 

		if !visited[node] {
			visited[node] = true
			for _, neighbor := range graph[node] {
				stack = append(stack, neighbor) 
			}
		}
	}

	fmt.Println(pathExist)
	// BFS solution
	// Note: Do the same thing but with a queue
	// visited := make(map[string]bool)
    // queue := []string{source}

    // for len(queue) > 0 {
    //     node := queue[0]
    //     queue = queue[1:]

    //     if node == destination {
    //         return true // Found a path
    //     }

    //     if !visited[node] {
    //         visited[node] = true
    //         for _, neighbor := range graph[node] {
    //             queue = append(queue, neighbor)
    //         }
    //     }
    // }

}

func BFS() {
	// RUNTIME: O(node + edge)
	graph := make(map[string][]string)
	graph = map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}	

	// Breadth first search 

	queue := []string{} 
	queue = append(queue, "a")

	for (len(queue) > 0) {
		val := queue[0]
		queue  = queue[1:]

		fmt.Println("Popped: ", val) 

		// for each child of a node 
		// push element to the queue
		for _, val := range graph[val] {
			queue = append(queue, val)
		}
	}
}

func DFS() {
	// RUNTIME: O(node + edge)
	graph := make(map[string][]string)
	graph = map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	// Depth first search

	stack := []string{} 
	stack = append(stack, "a")

	for len(stack) > 0 {
		val := stack[len(stack) - 1]
		stack  = stack[:len(stack) - 1]

		fmt.Println("Popped: ", val) 

		// for each child of a node 
		// push element to the stac
		for _, val := range graph[val] {
			stack = append(stack, val)
		}
	}
}



// TODO: Understand this
func courseSchedule() {
	/**
	There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] 
	indicates that you must take course bi first if you want to take course ai.

	For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
	Return true if you can finish all courses. Otherwise, return false.

	

	Example 1:

	Input: numCourses = 2, prerequisites = [[1,0]]
	Output: true
	Explanation: There are a total of 2 courses to take. 
	To take course 1 you should have finished course 0. So it is possible.

	*/

	numCourses := 5
	prerequisites := [][]int{{0, 1}, {0,2}, {1,3}, {1,4}, {3,4}} // [{1,0}] 1 = class, 0 = prereq

	isValid := true

	graph := make([][]int, numCourses)  
	inDegree := make([]int, numCourses)  // store the number of prereq a course needs 
	
	// Build the graph and calculate in-degrees
	for _, prerequisite := range prerequisites {
		course, prereq := prerequisite[0], prerequisite[1]
		graph[prereq] = append(graph[prereq], course)  
		inDegree[course]++ 
	}

	fmt.Println(inDegree)
	fmt.Println(graph)
	// Create a queue for topological sorting
	queue := make([]int, 0)

	// Initialize the queue with courses having no prerequisites
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	fmt.Println("QUEUE: ", queue)

	// Perform topological sorting
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]

		for _, nextCourse := range graph[course] {
			inDegree[nextCourse]--
			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	// Check if all courses have been taken (in-degree is 0)
	for _, degree := range inDegree {
		if degree > 0 {
			isValid= false 
			break;
		}
	}

	fmt.Println(isValid)
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
		{"1","1","1","1","0"},
		{"1","1","0","1","0"},
		{"1","1","0","0","0"},
		{"0","0","0","0","1"},
	}
  
	var island = 0
 
	for r := 0; r < len(grid); r++ {
		for c:= 0; c < len(grid[r]); c++ {

			// this grid is 1(island) and hasn't been visited
			if grid[r][c] == "1"  {
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
    
    q = append(q, []int{r,c})
    grid[r][c] = "2"
	
	// directions left, right, up, down
    offsets := [][]int{{-1,0}, {1,0}, {0,1}, {0,-1}}

    for len(q) > 0 {
        curr := q[0]
        q = q[1:]
        
		// check left,right,up,down from the current gridValue
        for _, offset := range offsets {
            x := curr[0] + offset[0]
            y := curr[1] + offset[1]
   
	 
            if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) && grid[x][y] == "1" {
                q = append(q, []int{x,y})
                grid[x][y] = "2"
            } 

        }
    }
}

func cloneGraph() {
	/*
	Description: https://leetcode.com/problems/clone-graph/
	*/

	 
}