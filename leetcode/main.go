package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// edges := [][]int{
	// 	{1, 2},
	// 	{1, 3},
	// 	{2, 4},
	// 	{2, 5},
	// 	{3, 6},
	// }

	// root := RecreateTreeFromEdges(edges)
	// PrintIndentedTree(root, 0, -1, true)

	pathSum()

}

// ***************** Binary Search Tree ***************** //
func pathSum() {
	edges := [][]int{
		{8, 3},
		{8, 10},
		{3, 1},
		{3, 6},
		{6, 4},
		{4, 7},
		{10, 14},
		{10, 9},
		{14, 13},
	}

	// Build BST given edges
	var root *BSTNode
	for _, edge := range edges {
		parent, child := edge[0], edge[1]

		if root == nil {
			root = &BSTNode{Val: parent}
		}
		root = insertIntoBST(root, child)
	}

	targetSum := 27
	// run dfs from this node
	var runDfs func(*BSTNode, int) bool

	runDfs = func(rootNode *BSTNode, currSum int) bool {
		if rootNode == nil {
			return false
		}

		currSum += rootNode.Val

		// meaning this is child
		if rootNode.Left == nil && rootNode.Right == nil {
			return currSum == targetSum
		}

		return runDfs(rootNode.Left, currSum) || runDfs(rootNode.Right, currSum)
	}

	hasPathSum := runDfs(root, 0)
	fmt.Println(hasPathSum)

	// PrintIndentedTree(root, 0, -1, true)
}

type BSTNode struct {
	Val   int
	Left  *BSTNode
	Right *BSTNode
}

func KthSmallestElementInBST(root *BSTNode) {
	// the trick is to search through the entire left side and add it to stack then go to the root node and then the right side
	kthElement := 5
	stack := []*BSTNode{root}
	current := root

	count := 0

	for len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		count++
		if kthElement == count {
			fmt.Printf("Kth %d element is %d\n", kthElement, current.Val)
			break
		}

		// move to the right
		current = current.Right
	}

}

func InOrderTraversal(root *BSTNode) {
	if root == nil {
		return
	}

	InOrderTraversal(root.Left)
	fmt.Printf("%d ", root.Val)
	InOrderTraversal(root.Right)
}

func buildBSTGivenEdges() {
	edges := [][]int{
		{3, 2},
		{3, 4},
		{2, 1},
		{4, 5},
	}

	var root *BSTNode

	for _, edge := range edges {
		parent, child := edge[0], edge[1]

		if root == nil {
			root = &BSTNode{Val: parent}
		}
		root = insertIntoBST(root, child)
	}
	KthSmallestElementInBST(root)
	//InOrderTraversal(root)
}

func insertIntoBST(root *BSTNode, value int) *BSTNode {
	if root == nil {
		return &BSTNode{Val: value}
	}

	if value < root.Val {
		// insert to left
		root.Left = insertIntoBST(root.Left, value)
	} else if value > root.Val {
		// insert to right
		root.Right = insertIntoBST(root.Right, value)
	}

	return root
}

// Minimum Height Tree to find the root
func MinHeightTree() {
	edges := [][]int{
		// {1, 2},
		// {2, 3},
		// {3, 4},
		{3, 4},
		{3, 0},
		{3, 1},
		{3, 2},
		{4, 5},
	}
	graph := make(map[int][]int)
	edgeMap := make(map[int]int)

	for _, edge := range edges {
		edge0, edge1 := edge[0], edge[1]

		graph[edge0] = append(graph[edge0], edge1)
		graph[edge1] = append(graph[edge1], edge0)
		edgeMap[edge0]++
		edgeMap[edge1]++
	}

	nonLeafNodes := []int{}
	for node, val := range edgeMap {
		if val != 1 {
			nonLeafNodes = append(nonLeafNodes, node)
		}
	}

	// smallest height node
	minHeightNodes := make(map[int]int)
	for _, node := range nonLeafNodes {
		visited := make(map[int]bool)
		height := 0
		// run BFS from nodes and record the smallest height
		queue := []int{node} // FIFO
		visited[node] = true

		for len(queue) > 0 {
			levelSize := len(queue)
			for i := 0; i < levelSize; i++ { // process all current nodes before moving to the next node
				queueVal := queue[0]
				queue = queue[1:]

				for _, neighbor := range graph[queueVal] {
					if !visited[neighbor] {
						visited[neighbor] = true
						queue = append(queue, neighbor)
					}
				}
			}
			height++
		}

		minHeightNodes[node] = height - 1
	}

	fmt.Println(minHeightNodes)

	minDegreeNodes := []int{}
	minVal := -1
	for node, val := range minHeightNodes {
		// this calculates the max value in the hashmap
		if val > minVal {
			minDegreeNodes = []int{node} // we directly modify the maxDegreeNode array
			minVal = val
		} else if val == minVal {
			minDegreeNodes = append(minDegreeNodes, node)
		}
	}

	fmt.Println(minDegreeNodes)
}

type TreeNode1 struct {
	Value    int
	Children []*TreeNode1
}

// CreateTreeNode creates a new TreeNode1 with the given value
func CreateTreeNode(value int) *TreeNode1 {
	return &TreeNode1{
		Value:    value,
		Children: []*TreeNode1{},
	}
}

// func PrintIndentedTree(node *TreeNode, depth int) {
//     if node == nil {
//         return
//     }

//     fmt.Printf("%*s%d\n", depth*2, "", node.Value)
//     for _, child := range node.Children {
//         PrintIndentedTree(child, depth+1)
//     }
// }

// PrintTree prints the tree in a depth-first manner
func PrintIndentedTree(node *TreeNode1, depth int, parentValue int, isLastChild bool) {
	if node == nil {
		return
	}

	branch := "├─ "
	if isLastChild {
		branch = "└─ "
	}

	fmt.Printf("%*s%s%d (Parent: %d, Children: %v)\n", depth*4, "", branch, node.Value, parentValue, getChildValues(node.Children))

	for i, child := range node.Children {
		isLast := i == len(node.Children)-1
		PrintIndentedTree(child, depth+1, node.Value, isLast)
	}
}

func getChildValues(children []*TreeNode1) []int {
	values := make([]int, len(children))
	for i, child := range children {
		values[i] = child.Value
	}
	return values
}

// Edge case: same number nodes
func RecreateTreeFromEdges(edges [][]int) *TreeNode1 {
	nodeMap := make(map[int]*TreeNode1)
	inDegree := make(map[int]int)

	// Create Nodes and build map
	for _, edge := range edges {
		u, v := edge[0], edge[1]

		if _, exists := nodeMap[u]; !exists {
			nodeMap[u] = CreateTreeNode(u)
		}

		if _, exists := nodeMap[v]; !exists {
			nodeMap[v] = CreateTreeNode(v)
		}

		// Todo handle duplicate case

		nodeMap[u].Children = append(nodeMap[u].Children, nodeMap[v])
		inDegree[u] += 0
		inDegree[v]++

	}

	// find the root node whose inDegree is 0
	var root *TreeNode1
	for _, node := range nodeMap {
		if inDegree[node.Value] == 0 {
			root = node
			break
		}
	}

	fmt.Println("ROOT: ", root)

	return root
}

// ***************** Graph ***************** //
func distanceFromNode() {
	graph := [][]int{
		3: {5, 1},
		5: {6, 2},
		2: {7, 4},
		1: {0, 8},
		6: {},
		7: {},
		4: {},
		0: {},
		8: {},
	}

	distance := 2
	fromNode := 5
	childParentNodes := make(map[int]int)

	// Initialize all nodes with parent = nil
	for node := range graph {
		childParentNodes[node] = -1 // You can use any value that represents nil, -1 is just an example.
	}

	for node := range graph {
		for _, val := range graph[node] {
			childParentNodes[val] = node
		}
	}

	// initialize Queue and Run BFS

	queue := []int{}
	visited := make(map[int]bool)
	queue = append(queue, fromNode)
	level := 0

	for len(queue) > 0 && level <= distance {
		queueVal := queue[0]
		for i := 0; i < len(queue); i++ {
			queue = queue[1:]

			if !visited[queueVal] {
				visited[queueVal] = true

				// also append it's children to run BFS
				for _, neighbor := range graph[queueVal] {
					if !visited[neighbor] {
						queue = append(queue, neighbor)
					}
				}

				// get the queueVal parent and append it to the queue
				parent := childParentNodes[queueVal]
				if parent != -1 && !visited[parent] {
					queue = append(queue, parent)
				}
			}
			level++
		}

	}

	fmt.Println(queue)

}

func numOfIsland() {
	var grid = [][]string{
		{"1", "1", "1", "1", "0"},
		{"1", "1", "0", "1", "0"},
		{"1", "1", "0", "0", "0"},
		{"0", "0", "0", "0", "1"},
	}

	numOfIslands := 0
	numRows := len(grid)
	numCols := len(grid[0])

	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			// if the grid value is 1 do something
			if grid[r][c] == "1" {
				run_bfs(grid, r, c)
				numOfIslands += 1
			}
		}
	}

	fmt.Println(numOfIslands)
}

func run_bfs(grid [][]string, row, col int) {
	queue := [][]int{}
	queue = append(queue, []int{row, col})
	grid[row][col] = "2"

	for len(queue) > 0 {
		queueVal := queue[0]
		queue = queue[1:]

		directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

		for _, direction := range directions {
			// check if these direction is in bound
			// [0] is row and [1] is column
			row := queueVal[0] + direction[0]
			col := queueVal[1] + direction[1]
			if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[row]) && grid[row][col] == "1" {
				queue = append(queue, []int{row, col})
				grid[row][col] = "2"
			}
		}

	}
}

func courseSchedule_() {
	numCourses := 5
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {3, 4}}
	graph := make(map[int][]int)

	for _, edge := range edges {
		edge0, edge1 := edge[0], edge[1]
		graph[edge0] = append(graph[edge0], edge1)
	}

	visited := make([]bool, numCourses)

	isValid := true

	var dfs func(course int)
	dfs = func(course int) {
		visited[course] = true

		for _, neighbor := range graph[course] {
			if !visited[neighbor] {
				dfs(neighbor)
			} else {
				isValid = false // Detected a cycle
				break
			}
		}

		visited[course] = false
	}

	for course := 0; course < numCourses; course++ {
		if !visited[course] {
			dfs(course)
		}
	}

	fmt.Println(isValid)
}

type CloneNode struct {
	Val       int
	Neighbors []*CloneNode
}

func cloneGraph(node *CloneNode) *CloneNode {

	// run this in main() to test
	// node1 := &CloneNode{Val: 1}
	// node2 := &CloneNode{Val: 2}
	// node3 := &CloneNode{Val: 3}
	// node4 := &CloneNode{Val: 4}

	// node1.Neighbors = []*CloneNode{node2, node4}
	// node2.Neighbors = []*CloneNode{node1, node3}
	// node3.Neighbors = []*CloneNode{node2, node4}
	// node4.Neighbors = []*CloneNode{node1, node3}

	// clonedNode := cloneGraph(node1)
	// fmt.Println(clonedNode)

	visited := make(map[*CloneNode]*CloneNode)

	var clone func(*CloneNode) *CloneNode

	clone = func(original *CloneNode) *CloneNode {
		cloned, exist := visited[original]
		if exist {
			return cloned
		}

		newNode := &CloneNode{Val: original.Val}
		visited[original] = newNode

		for _, neighbor := range original.Neighbors {
			newNeighbor := clone(neighbor)
			newNode.Neighbors = append(newNode.Neighbors, newNeighbor)
		}

		return newNode
	}

	return clone(node)
}

func containsCycle() {
	graph := map[string][]string{
		"A": {"B"},
		"B": {"C"},
		"C": {"D"},
		"D": {},
	}

	var dfs func(course string)
	visited := make(map[string]bool)
	containsCycle := false

	dfs = func(course string) {
		visited[course] = true

		for _, neighbor := range graph[course] {
			if !visited[neighbor] {
				dfs(neighbor)
			} else {
				containsCycle = true
				break
			}
		}
	}

	for node := range graph {
		if !visited[node] {
			dfs(node)
		}
	}

	fmt.Println(containsCycle)
}

// BFS
func BFS() {
	graph := map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	queue := []string{"a"} // FIFO

	for len(queue) > 0 {
		queueVal := queue[0]
		queue = queue[1:]

		fmt.Println(queueVal)

		for _, node := range graph[queueVal] {
			queue = append(queue, node)
		}

	}
}

func hasPath() {
	// find the shortest distance from A to E
	var graph = map[string][]string{
		"A": {"B", "C"},
		"B": {"D"},
		"C": {"E"},
		"D": {"F"},
		"E": {},
		"F": {},
	}

	// check to see if a path exist from A to E
	end := "E"
	stack := []string{"B"}
	hasVisited := make(map[string]bool)

	for len(stack) > 0 {
		stackVal := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if stackVal == end {
			fmt.Println("Path exist")
			break
		}

		if !hasVisited[stackVal] {
			hasVisited[stackVal] = true

			for _, node := range graph[stackVal] {
				stack = append(stack, node)
			}
		}

	}
}

func shortestPath() {
	var graph = map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "D"},
		"C": {"A", "E"},
		"D": {"B", "E"},
	}

	destination := "E"
	queue := [][]string{{"A"}}
	hasVisited := make(map[string]bool)

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		node := path[len(path)-1]

		if node == destination {
			// Path found
			fmt.Println(path)
			break
		}

		if !hasVisited[node] {
			hasVisited[node] = true

			for _, val := range graph[node] {
				newPath := append([]string{}, path...)
				newPath = append(newPath, val)

				queue = append(queue, newPath)
			}
		}

	}
}

// DFS
func DFS() {
	graph := map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	stack := []string{}
	stack = append(stack, "a")

	for len(stack) > 0 {
		stackVal := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Println(stackVal)

		for _, node := range graph[stackVal] {
			stack = append(stack, node)
		}
	}
}

func countConnectedComponents() {
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

	count := 0
	hasVisited := make(map[string]bool)

	for node := range graph {
		stack := []string{node}

		if !hasVisited[node] {
			for len(stack) > 0 {
				stackVal := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if !hasVisited[stackVal] {
					hasVisited[stackVal] = true
					for _, node := range graph[stackVal] {
						stack = append(stack, node)
					}
				}
			}
			count++
		}
	}

	fmt.Println(count)
}

func buildGraphGivenEdgesAndFindRoot() {
	edges := [][]string{
		{"N1", "N3"},
		{"N3", "A"},
		{"A", "B"},
		{"A", "D"},
		{"B", "C"},
		{"D", "C"},
	}

	graph := make(map[string][]string)
	inDegree := make(map[string]int)
	for _, edge := range edges {
		edge0, edge1 := edge[0], edge[1]
		graph[edge0] = append(graph[edge0], edge1)
		inDegree[edge[1]]++

	}

	// Find the node with an in-degree of 0, which is the root
	for node, degree := range inDegree {
		if degree == 0 {
			fmt.Println("NODE: ", node)
			break
		}
	}

	fmt.Println(inDegree)
}

// ***************** Stack ***************** //
func ValidParenthesis() {
	s := "()[]" // true, ()[]{} = true, (] false

	hashMap := make(map[rune]rune)

	hashMap[')'] = '('
	hashMap['}'] = '{'
	hashMap[']'] = '['

	// LIFO
	stack := []rune{}
	output := true

	for _, val := range s { // TODO check if this is the correct loop
		newVal, ok := hashMap[val]
		if ok {
			stackVal := stack[len(stack)-1]
			if len(stack) != 0 && newVal == stackVal {
				stack = stack[:len(stack)-1]
			} else {
				output = false
				break
			}
		} else {
			stack = append(stack, val)
		}

	}

	if len(stack) != 0 {
		output = false
	}

	fmt.Println(output)
}

// ***************** String ***************** //

func longestRepeatedCharacterReplacement() {
	s := "AABABBA" // Output: 4, Replace the two 'A's with two 'B's or vice versa.
	k := 1
	hashmap := make(map[rune]int)
	start := 0
	maxFreq := float64(0)
	output := 0

	for i, val := range s {
		hashmap[val]++

		maxFreq := max(int(maxFreq), hashmap[val])
		freq := (i - start + 1) - maxFreq

		if freq <= k {
			output++
		} else {
			hashmap[val]--
			start++
		}

	}

	fmt.Println(output)
}

// Runtime: nlogn
func PermutationInString() {
	s1 := "ab"
	s2 := "eidbeoba" // output true s2 contains one permutation of s1 ("ba").

	s1 = SortString(s1)
	perm := false
	end := 0

	for end < len(s2)-1 {
		subStr := s2[end : end+2]

		sortedSubStr := SortString(subStr)

		if s1 == sortedSubStr {
			perm = true
			break
		}

		end++
	}

	fmt.Println(perm)
}

func ValidPalindrome() {
	str := "racecar"
	left := 0
	right := len(str) - 1
	isPalindrome := true

	for left != right {
		if str[left] != str[right] {
			isPalindrome = false
			break
		}
		left++
		right--
	}

	fmt.Println(isPalindrome)
}

func FindAllAnagram() {
	s := "baccbaacb"
	p := "abc"
	counter := len(p)

	wordHashMap := make(map[rune]int)
	anagrams := []string{}
	start := 0

	for _, val := range p {
		wordHashMap[val]++
	}

	for i, str := range s {
		if wordHashMap[str] > 0 {
			// this str is from word
			counter--
		}

		wordHashMap[str]--

		// determine when you should start checking for anagrams. check if the current window is greater than or equal to 3
		if i-start+1 >= len(p) {
			if counter == 0 {
				// Anagram found
				anagramStr := s[start : i+1]
				anagrams = append(anagrams, anagramStr)
			}

			wordHashMap[rune(s[start])]++
			if wordHashMap[rune(s[start])] > 0 {
				counter++
			}

			start++
		}

	}

	fmt.Println(anagrams)
}

func MinWindowSubString() {
	str := "ADOBECODEBANC"
	word := "ABC" // output: "BANC"
	res := ""
	minWindow := len(str)

	start := 0
	wordHashMap := make(map[rune]int)
	counter := len(word)
	for _, val := range word {
		wordHashMap[val]++
	}

	for i := 0; i < len(str); i++ {
		val, _ := wordHashMap[rune(str[i])]
		if val > 0 {
			counter--
		}
		wordHashMap[rune(str[i])]--

		for counter == 0 {
			// this means we have found a substring
			window := i - start + 1
			if minWindow > window {
				minWindow = window
				res = str[start : i+1]
			}

			wordHashMap[rune(str[start])]++
			if wordHashMap[rune(str[start])] > 0 {
				counter++
			}
			start += 1
		}

	}

	fmt.Println(res)
}

// Note: (SlidingWindow) confusing problem
func RepeatedDNASequence() {
	// 	Given a string s that represents a DNA sequence, return all the 10-letter-long sequences (substrings) that occur more than once
	// 	in a DNA molecule. You may return the answer in any order.
	//	Output: ["AAAAACCCCC","CCCCCAAAAA"]

	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	subStrs := []string{}
	count := 0

	hashMap := make(map[string]bool)

	for i := 10; i < len(s); i += 1 {
		if i > len(s)-1 {
			break
		}

		subStr := s[count:i]

		if hashMap[subStr] {
			subStrs = append(subStrs, subStr)
		}

		hashMap[subStr] = true
		count++
	}

	// fmt.Println(hashMap)
	fmt.Println(subStrs)
}

// Note: Very interesting problem (check what hashmap is storing & how to increment left pointer)
func longestSubString() {
	// find the length of the longest substring without repeating characters.
	// Sliding Window
	str := "pwwkew" // output = 3 abc
	left := 0
	longestSubStr := 0
	hashmap := make(map[rune]int)

	for i, s := range str {
		val, ok := hashmap[s]
		if ok { // if we've found a repetetive char
			// get the current window
			left = val + 1
		}

		if longestSubStr < i-left+1 {
			longestSubStr = i - left + 1
		}

		hashmap[s] = i
	}
	fmt.Println(longestSubStr)
}

func InitStringToMap() {
	alph := "abcdefghijklmnopqrstuvwz"
	s := "110#11#12"
	res := ""
	i := 0

	for i < len(s) {
		if i < len(s)-2 && s[i+2] == '#' {
			// get the substring
			strIdx, _ := strconv.Atoi(s[i : i+2]) // this will be a number
			subStr := alph[strIdx-1]

			res += string(subStr)
			i += 3
		} else {
			strIdx, _ := strconv.Atoi(string(s[i]))
			subStr := alph[strIdx-1]
			res += string(subStr)
			i++
		}
	}

	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ***************** Array ***************** //
// Note: SlidingWindow
func MinSizeSubArraySum() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3} // output: 2 The subarray [4,3] has the minimal length under the problem constraint.

	start := 0
	minSum := 0
	windowSize := len(nums)

	for i, num := range nums {
		minSum = minSum + num

		for minSum > target {
			minSum = minSum - nums[start]
			start = start + 1
		}

		if minSum == target {
			// we want to check if the length of the window
			if i-start+1 < windowSize {
				windowSize = i - start + 1
			}
		}
	}

	fmt.Println(windowSize)

}

func twoSum() {
	nums := []int{2, 7, 11, 15}
	target := 17

	numSet := make(map[int]int)
	result := []int{}

	for i, val := range nums {
		comp := target - val
		idx, ok := numSet[comp]
		if ok {
			result = append(result, idx, i)
			break
		}

		numSet[val] = i
	}

	fmt.Println(result)
}

func TimeToBuySellStock() {
	prices := []int{7, 6, 4, 3, 1} // output = 5

	left := 0
	right := 1

	profit := 0

	for right < len(prices) {
		if prices[left] < prices[right] {
			// profit
			if prices[right]-prices[left] > profit {
				profit = prices[right] - prices[left]
			}
		} else {
			// loss so shift left
			left = right
		}

		right++
	}

	fmt.Println(profit)
}

func containsDuplicate() {
	nums := []int{1, 2, 3, 1, 1, 1}
	maps := make(map[int]bool)
	containsDuplicate := false

	for _, val := range nums {
		if maps[val] {
			containsDuplicate = true
			break
		}

		maps[val] = true
	}

	fmt.Println(containsDuplicate)
}

// Note: redo
func ProductOfArrExceptSelf() {
	arr := []int{1, 2, 3, 4}
	leftProductArr := make([]int, len(arr))
	rightProductArr := make([]int, len(arr))

	leftProduct := 1
	for i := 0; i < len(arr); i++ {
		leftProductArr[i] = leftProduct
		leftProduct = leftProduct * arr[i]
	}

	// Calculate right products
	rightProduct := 1
	for i := len(arr) - 1; i >= 0; i-- {
		rightProductArr[i] = rightProduct
		rightProduct *= arr[i]
	}

	result := make([]int, len(arr))
	// Calculate the result
	for i := 0; i < len(arr); i++ {
		result[i] = leftProductArr[i] * rightProductArr[i]
	}

	fmt.Println(result)
}

// Note: nice problem
func MaxSubArray() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4} // 4,-1,2,1 = 6
	maxSub := nums[0]
	currSum := 0

	for _, val := range nums {
		if currSum < 0 {
			currSum = 0
		}

		currSum = val + currSum
		maxSub = max(maxSub, currSum)
	}

	fmt.Println(maxSub)
}

// Note: confusing around negative case edge case of 0 as well
func MaxProductSubarray() {
	nums := []int{-1, -2, -3}
	maxProd := nums[0]
	currMax := nums[0]
	currMin := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			currMax, currMin = currMin, currMax // Swap max and min when encountering a negative number
		}

		currMax = max(nums[i], currMax*nums[i])
		currMin = min(nums[i], currMin*nums[i])

		maxProd = max(maxProd, currMax)
	}

	fmt.Println(maxProd)
}

func FindMinInRotatedArray() {
	nums := []int{2, 3, 4, 5}

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (right + left) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	fmt.Println(nums[left])
}

func SearchInRotatedArray() {
}

func three_Sum() {
	nums := []int{-1, 0, 1, 2, -1, -4} // [[-1,-1,2],[-1,0,1]]
	target := 0
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		numSet := make(map[int]int)
		currTarget := target - nums[i]

		for j := i + 1; j < len(nums); j++ {
			complement := currTarget - nums[j]

			if idx, ok := numSet[complement]; ok {
				result = append(result, []int{nums[i], nums[j], nums[idx]})
			}

			numSet[nums[j]] = j
		}
	}

	fmt.Println(result)
}

// Dynamic programming
// TODO: understand this
func UniquePaths() {
	m := 3
	n := 7

	// Create a 2D slice dp to store the number of unique paths
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Initialize the first row and first column to 1
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// Calculate dp[i][j] for the rest of the cells
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	// The result is stored in dp[m-1][n-1]
	fmt.Println(dp[m-1][n-1])
}

// Note: What a problem!!! maybe revisit this
func DecodeWays() {
	s := "2212"

	n := len(s)

	// Handle edge case where the string is empty
	if n == 0 {
		fmt.Println("0")
		return
	}

	// Initialize a DP array to store the number of ways to decode substrings
	dp := make([]int, n+1)

	// Base case: There's one way to decode an empty string
	dp[0] = 1

	// Check the first character of the string
	if s[0] == '0' {
		fmt.Println("0") // If it's '0', there's no valid decoding
		return
	}

	dp[1] = 1 // There's one way to decode a single character string

	// Populate the DP array
	for i := 2; i <= n; i++ {
		// Check if the current character is '0', and if the previous character is '1' or '2'
		if s[i-1] == '0' && (s[i-2] == '1' || s[i-2] == '2') {
			dp[i] = dp[i-2]
		} else if s[i-2:i] >= "10" && s[i-2:i] <= "26" {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}

	fmt.Println(dp[n])
}

func HouseRobber() {
	house := []int{2, 7, 9, 3, 1}

	dp := make([]int, len(house))
	dp[0] = house[0]
	dp[1] = max(house[0], house[1])

	for i := 2; i < len(house); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+house[i])
	}

	fmt.Println(dp)
}

func HouseRobberII() {
	house := []int{2, 7, 9, 3, 1}

	robHelper := func(nums []int) int {
		dp := make([]int, len(house))
		dp[0] = house[0]
		dp[1] = max(house[0], house[1])

		for i := 2; i < len(house); i++ {
			dp[i] = max(dp[i-1], dp[i-2]+house[i])
		}

		return dp[len(nums)-1]
	}

	val := max(robHelper(house[:len(house)-1]), robHelper(house[1:]))

	fmt.Println(val)
}

// TODO do this
func CombinationSum() {
}

// NOTE: Very nice problem
func WordBreak() {
	s := "leet"
	dict := []string{"leet", "code"}

	left := 0
	right := 1

	res := ""

	for i := 0; i < len(s); i++ {
		letter := s[left:right]

		// check if val is present in dict
		for _, val := range dict {
			if val == letter {
				res += letter
				left = right
			}
		}

		right++
	}

	fmt.Println(res == s)
}

func longestIncreasingSubSequence() {
	nums := []int{3, 4, -1, 0, 6, 2, 3}
	result := make([]int, len(nums))
	maxLen := 1

	for i := 0; i < len(nums); i++ {
		result[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && result[j] >= result[i] {
				// longest increasing subsequence will be atleast longest increasing subsequence at j + 1
				result[i] = result[j] + 1
			}
		}

		if result[i] > maxLen {
			maxLen = result[i]
		}
	}

	fmt.Println(result)
	fmt.Println(maxLen)
}

func coinChange() {
	coins := []int{1, 2, 5}
	amount := 13
	val := []int{}

	for amount > 0 {
		max := -1

		// find the max value in coins
		for _, val := range coins {
			if val > max && amount >= val {
				max = val

			}
		}

		amount = amount - max
		val = append(val, max)
	}

	fmt.Println(val)
}

func ClimbingStairs() {
	// 3 stairs how many ways you can climb it
	n := 3

	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		if i != n {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}

	fmt.Println(dp[n-1])
}

// TODO do this
func threeSum() {
}

func longestPalindromicSubstring() {
	input := "babad"
	result := ""
	resLen := 0

	// this checks for palindrome from the center
	expandAroundCenter := func(left, right int) {
		for left >= 0 && right < len(input) && input[left] == input[right] {
			if right-left+1 > resLen {
				// we have found the longest palindrome
				result = input[left : right+1]
				resLen = right - left + 1
			}

			left--
			right++
		}
	}

	for i := 0; i < len(input); i++ {
		// Odd length palindrome
		expandAroundCenter(i, i)

		// Even length palindrome
		//expandAroundCenter(i, i+1)
	}

	fmt.Println(result)
}

// Note: pretty annoying problem
func Permutations() {
	nums := []int{1, 2, 3}
	result := [][]int{}
	result = append(result, []int{nums[0]})

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		newPerms := [][]int{}

		for _, perm := range result {
			for j := 0; j <= len(perm); j++ {
				newPerm := append([]int{}, perm[:j]...) // this is j exclusive
				newPerm = append(newPerm, num)
				newPerm = append(newPerm, perm[j:]...) // this is j inclusive
				newPerms = append(newPerms, newPerm)
				fmt.Println("NEW PERMS: ", newPerms)
			}
		}

		result = newPerms
	}

	fmt.Println(result)
}

func AddTwoNumbers() {
	num1 := 243 // reverse 342
	digits1 := []int{}

	num2 := 564 // reverse 465
	digits2 := []int{}

	// Extract the digits and put them into the array reversed .
	for num1 > 0 {
		digit := num1 % 10
		digits1 = append(digits1, digit)
		num1 = num1 / 10
	}

	// Extract the digits and put them into the array reversed .
	for num2 > 0 {
		digit := num2 % 10
		digits2 = append(digits2, digit)
		num2 = num2 / 10
	}

	reversedNum1 := 0
	numZeros := len(digits1) - 1

	for i := 0; i < len(digits1); i++ {
		num := digits1[i] * int(math.Pow(10, float64(numZeros)))
		reversedNum1 = reversedNum1 + num
		numZeros--
	}

	reversedNum2 := 0
	numZeros1 := len(digits2) - 1

	for i := 0; i < len(digits2); i++ {
		num := digits2[i] * int(math.Pow(10, float64(numZeros1)))
		reversedNum2 = reversedNum2 + num
		numZeros1--
	}

	fmt.Println(reversedNum1 + reversedNum2)
}

// NOTE very hard problem
func LongestCommonSubsequence() {
}

func containerWithMostWater() {
	input := []int{1, 8, 6, 2, 5, 4, 6, 3, 7}
	left := 0
	right := len(input) - 1
	maxArea := 1

	for left != right {
		area := (right - left) * int(math.Min(float64(input[left]), float64(input[right])))

		if area > maxArea {
			maxArea = area
		}

		if input[left] < input[right] {
			left++
		} else {
			right--
		}
	}

	fmt.Println(maxArea)
}

// Note: Very nice problem
// RUNTIME: 4^N (each number can have atmost 4 letter associated with it)
func PhoneNumberLetterComb() {
	letterNumComb := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	input := "23"
	result := []string{""}

	for _, val := range input {
		str, ok := letterNumComb[val]
		if !ok {
			fmt.Println("invalid number ")
			break
		}

		newRes := []string{}
		for _, res := range result {
			for _, s := range str {
				newRes = append(newRes, string(res)+string(s))
			}
		}

		result = newRes
	}

	fmt.Println(result)
}

// Note: Complex (DFS on island like problem)
func WordSearch() {
	board := [][]string{
		{"A", "B", "C", "E"},
		{"S", "F", "C", "S"},
		{"A", "D", "E", "E"},
	}

	word := "ABCCED"

	lenRows := len(board)
	lenColm := len(board[0])

	visited := make(map[string]bool)

	var dfs func(int, int, int) bool
	dfs = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}

		if r < 0 || r >= lenRows || c < 0 || c >= lenColm || visited[string(r)+string(c)] || board[r][c] != string(word[i]) {
			// out of bounds case or wrong character found case or already visited grid
			return false
		}

		visited[string(r)+string(c)] = true

		res := (dfs(r+1, c, i+1) || dfs(r-1, c, i+1) || dfs(r, c+1, i+1) || dfs(r, c-1, i+1))

		visited[string(r)+string(c)] = false
		return res
	}

	for r := 0; r < lenRows; r++ {
		for c := 0; c < lenColm; c++ {
			if dfs(r, c, 0) {
				fmt.Println("TRUE ")
				return
			}
		}
	}

	fmt.Println("FALSE")
}

// Note complex logic around subset
func Subsets() {
	nums := []int{1, 2, 3}

	res := [][]int{{}}

	fmt.Println(len(res))

	for i := 0; i < len(nums); i++ {
		currLen := len(res)

		for j := 0; j < currLen; j++ {
			newSubSet := append([]int{}, res[j]...) // creates a copy of the empty subset [], resulting in [].
			newSubSet = append(newSubSet, nums[i])  // append([], nums[i]) appends 1 to the empty slice, resulting in [1]

			res = append(res, newSubSet)

			fmt.Println(res)
		}

	}
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
