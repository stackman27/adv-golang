package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	longestRepeatedCharacterReplacement()
}

// String

// TODO understand this
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

// Arrays
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
