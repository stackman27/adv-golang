package main

import (
	"fmt"
	"sort"
)

func funccommonLeetCode() {

}

// COMMON LEETCODE  Resource : https://www.youtube.com/watch?v=Peq4GCPNC5c

func validAnagram() {
	s1 := "danger"
	s2 := "garden"

	hashmap := make(map[rune]int)
	counter := len(s1)
	isAnagram := false

	for _, val := range s1 {
		hashmap[val]++
	}

	for _, val := range s2 {
		_, ok := hashmap[val]
		if !ok {
			// not an anagram
			break
		}

		hashmap[val]--
		counter--

		if counter == 0 {
			// it is an anagram
			isAnagram = true
			break
		}
	}

	fmt.Println(isAnagram)
	fmt.Println(hashmap)
}

func FindFirstAndLastBS() {
	arr := []int{2, 4, 5, 5, 5, 7}
	target := 5

	// left value
	val1 := findFirstAndLastPositionHelper(arr, target, true)
	val2 := findFirstAndLastPositionHelper(arr, target, false)

	fmt.Println(val1, val2)
}

func findFirstAndLastPositionHelper(arr []int, target int, leftBias bool) int {
	// perform binary search
	left := 0
	right := len(arr) - 1
	midIdx := -1

	for left < right {
		//mid := (left+right) / 2 // this can overflow (float point division)
		mid := left + (right-left)/2

		val := arr[mid]

		if val == target {
			midIdx = mid
			if leftBias {
				// leftBias = true, BS will continue searching to the left
				right = mid - 1
			} else {
				// leftBias = true, BS will continue searching to the left
				left = mid + 1
			}
		} else if val < target {
			// search right
			left = mid + 1
		} else {
			// search left
			right = mid
		}
	}

	return midIdx
}

func FindKthLargestElement() {
	arr := []int{4, 2, 9, 7, 5, 6, 7, 1, 3}
	k := 4

	sort.Ints(arr)

	forthLargest := arr[len(arr)-k]

	fmt.Println(forthLargest)
}

type TreeNode struct {
	val       int
	leftNode  *TreeNode
	rightNode *TreeNode
}

func checkBTSymmetric() {
	// check BT is mirror of itself

	newTreeNode := &TreeNode{
		val: 1,
		leftNode: &TreeNode{
			val:       3,
			leftNode:  &TreeNode{val: 3},
			rightNode: &TreeNode{val: 4},
		},

		rightNode: &TreeNode{
			val:       3,
			leftNode:  &TreeNode{val: 4},
			rightNode: &TreeNode{val: 3},
		},
	}

	// Perform BFS
	queue := []*TreeNode{newTreeNode.leftNode, newTreeNode.rightNode} // Initialize the queue with two root nodes
	isSymmetrical := true

	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]

		queue = queue[2:]

		// Check if the values of the current nodes are equal
		if left == nil && right == nil {
			continue // Both nodes are nil, continue to the next level
		}
		if left == nil || right == nil {
			isSymmetrical = false
			break
		}
		if left.val != right.val {
			isSymmetrical = false
			break
		}

		// Enqueue left and right child nodes in opposite order
		queue = append(queue, left.leftNode, right.rightNode)
		queue = append(queue, right.leftNode, left.rightNode)
	}

	fmt.Println("ISSYMMETRICAL: ", isSymmetrical)
}

type Parenthesis struct {
	Str   string
	Open  int
	Close int
}

func generateParenthesis() {
	n := 2
	if n == 0 {
		return
	}

	result := []string{}
	stack := []Parenthesis{{"(", 1, 0}}

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if p.Open == n && p.Close == n {
			result = append(result, p.Str)
		}

		if p.Open < n {
			stack = append(stack, Parenthesis{p.Str + "(", p.Open + 1, p.Close})
		}

		if p.Close < p.Open {
			stack = append(stack, Parenthesis{p.Str + ")", p.Open, p.Close + 1})
		}
	}

	fmt.Println(result)
}

// Note: pretty shitty problem: https://www.youtube.com/watch?v=lJwbPZGo05A&t=1s
func GasStation() {
	// check the solution exist first
}

// TODO understand this better
func courseSchedule1() {
	prerequisites := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {3, 4}}
	numCourses := 4
	isPossible := true

	graph := make(map[int][]int)
	// convert the edges into a graph structure
	for _, prerequisite := range prerequisites {
		child, parent := prerequisite[0], prerequisite[1]
		graph[parent] = append(graph[parent], child)
	}

	// Create an array to keep track of visited nodes during DFS
	visited := make([]bool, numCourses)

	// Create an array to keep track of nodes being visited in the current DFS path
	inPath := make([]bool, numCourses)

	// Perform DFS from each node
	stack := []int{}

	for i := 0; i < numCourses; i++ {
		stack = append(stack, i)
		inPath[i] = true

		for len(stack) > 0 {
			node := stack[len(stack)-1]
			visited[node] = true
			inPath[node] = true

			for _, neighbor := range graph[node] {
				if inPath[neighbor] {
					// Cycle detected, return false
					isPossible = false
					break
				}
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}

		inPath[i] = false

	}

	fmt.Println(isPossible)
}

func shortestSubstring() {

}
