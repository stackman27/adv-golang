package main

import (
	"fmt"
	"sort"
	"strconv"
)

// three sum

// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
// Note: The solution set must not contain duplicate triplets.

// Example:
// Given array nums = [-1, 0, 1, 2, -1, -4],
// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

func threeSumMain() {

	arr := []int{-1,0,1,2,-1,-4}

	// naive approach 3 for loops 

	// smart solution 
	triplets := make([][]int, 0)	
	sort.Ints(arr)

	if len(arr) < 3 {
		println(triplets)
		return 
	}

	// pick a number then solve 2 sum problem in space in front of it 
	for i := 0; i <len(arr); i++ {

		// skip duplicates if we've seen it in the past 
		if i > 0 && arr[i] == arr[i -1] {
			continue
		}

		need := 0 - arr[i]
		start := i + 1 
		end := len(arr) - 1 

		for start < end { 
			if arr[start] + arr[end] == need {
				triplets = append(triplets, []int{arr[i], arr[start], arr[end]})

				end -- 
				for start < end && arr[end] == arr[end+1] {
					end --
				}
			} else if arr[start] + arr[end] > need {
				end --
			} else {
				start ++ 
			}
		}
	}

	fmt.Println(triplets)

}

 

// TODO: not sure what's wrong with this algorithm
func MaxProductSubArray() {
	/*
	Maximum Product Subarray
	Given an integer array nums, find a subarray
	that has the largest product, and return the product.

	The test cases are generated so that the answer will fit in a 32-bit integer.

	

	Example 1:

	Input: nums = [2,3,-2,4]
	Output: 6
	Explanation: [2,3] has the largest product 6.
	Example 2:

	Input: nums = [-2,0,-1]
	Output: 0
	Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
	*/

	nums := []float64{-1,-2,-3} 
	product := float64(1)
	output := float64(0)

	for i := 0; i < len(nums); i++ {
		product = product * nums[i]
		if product > output {
			output = product
		}
	}	
	
 
	fmt.Println(output)

}

//TODO
func ProductArrayExceptSelf() {
	/*	
	Product of Array Except Self
	Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of 
	nums except nums[i].

	The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

	You must write an algorithm that runs in O(n) time and without using the division operation.

	

	Example 1:

	Input: nums = [1,2,3,4]
	Output: [24,12,8,6]
	Example 2:

	Input: nums = [-1,1,0,-3,3]
	Output: [0,0,9,0,0]
	*/
 
}

func BinarySearch1() {
	/*	 
	Search in Rotated Sorted Array

	There is an integer array nums sorted in ascending order (with distinct values).

	Prior to being passed to your function, nums is possibly rotated at an unknown pivot index 
	k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). 
	For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

	Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.
	You must write an algorithm with O(log n) runtime complexity.

	

	Example 1:

	Input: nums = [4,5,6,7,0,1,2], target = 0
	Output: 4
	Example 2:

	Input: nums = [4,5,6,7,0,1,2], target = 3
	Output: -1
	Example 3:

	Input: nums = [1], target = 0
	Output: -1

	*/

	nums := []int{4,5,6,7,0,1,2} 
	searchVal := 6 // should return 3
	start := 0
	end := len(nums) -1 

	value := -1

	for start <= end {
		mid := (start + end) / 2 // midIndex

		if nums[mid] == searchVal {
			value = mid
			break;
		}

		if nums[start] <= nums[mid] {
			// search left 

			// now its just sorted binary search 
			if searchVal > nums[start] && searchVal < nums[mid] {
				// we go to left 
				end = mid 
			} else {
				// we go right
				start = mid + 1
			}
		} else {
			// search right 
			
			// now its just sorted binary search 
			if searchVal > nums[start] && searchVal < nums[mid] {
				// we go to left 
				end = mid 
			} else {
				// we go right
				start = mid + 1
			}
		}
	}

	fmt.Println(value)
}

func BinarySearch() {
	/*
	Find Min in rotated SubArray
	Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] 
	might become:

	[4,5,6,7,0,1,2] if it was rotated 4 times.
	[0,1,2,4,5,6,7] if it was rotated 7 times.
	Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
	Given the sorted rotated array nums of unique elements, return the minimum element of this array.

	You must write an algorithm that runs in O(log n) time.

	Example 1:

	Input: nums = [3,4,5,1,2]
	Output: 1
	Explanation: The original array was [1,2,3,4,5] rotated 3 times.
	*/

	nums := []int{3,4,5,1,2}
	left := 0
	right := len(nums) - 1
	
	for left < right {
		midValue := (left + right) / 2 
		
		// midValue is bigger than the right most value 
		if nums[midValue] > nums[right] {
			left = midValue	+ 1		
		} else {
			right = midValue
		}	
	}


	fmt.Println(nums[left])
}

func initStringToMap() {

	/***
Example 1:

Input: s = "10#11#12"
Output: "jkab"
Explanation: "j" -> "10#" , "k" -> "11#" , "a" -> "1" , "b" -> "2".
Example 2:

Input: s = "1326#"
Output: "acz"
Example 3:

Input: s = "25#"
Output: "y"
Example 4:

Input: s = "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"
Output: "abcdefghijklmnopqrstuvwxyz"

**/

chars := "abcdefghijklmnopqrstuvwxyz"
d := ""
i := 0 

s := "10#1"

// 4 - 2 s[2] = # print 1
// 2 - 2 
//
for i < len(s) {
	if i < len(s) - 2 && string(s[i+2]) == "#" {
		alphaIdx, _ := strconv.Atoi(s[i:i+2])
		d += string(chars[alphaIdx - 1])
		i += 3
	} else {
		alphaIdx, _ := strconv.Atoi(string(s[i]))
		d += string(chars[alphaIdx -1])
		i+= 1
	}
}

println(d)
}

func longestSubstring() {
	/*
		Given a string, find the length of the longest substring without repeating characters.

		Examples:
			Given abcabcbb, the answer is abc, which the length is 3.

		1. When `i = 0`, we have the character 'a'. We haven't seen 'a' before, so we just record 'a' in `lastOccurred` with a value of 0, which is the current index.

		2. When `i = 1` (b), we record 'b' in `lastOccurred` with a value of 1. We also have a new maxLength = `i - start + 1 = 2`.

		3. When `i = 2` (c), we record 'c' in `lastOccurred` with a value of 2 and maxLength = `i - start + 1 = 3`.

 
		4. When `i = 3` (a), we notice that 'a' already exists in `lastOccurred` as it occurred at index 0. So it's a repeating character. We move our `start` to `lastOccurred['a'] (which is 0) + 1 = 1` (the next character after the previous 'a'). maxLength stays the same because `i - start + 1 = 3 - 1 + 1 = 3` is not greater than `maxLength = 3`.

		5. When `i = 4` (b), 'b' is also repeating, found at index 1. We reset `start` to `lastOccurred['b'] + 1 = 1 + 1 = 2` (the next character after the earlier 'b'). 

		5. Now if we continue to follow this process, when `i = 7`, we come across 'b' which has already been recorded at index 1. So we update `start` again to `lastOccurred['b'] + 1 = 2`.
	*/


	str := "abcebcbb" 
	lastOccoured := make(map[rune]int)
	start, maxLength := 0, 0

	for i, char := range str {
		lastVal, ok := lastOccoured[char] 
		if ok {
			// this value has occoured in the past
			 start = lastVal + 1
		}

		// store the window length 
		if i - start + 1 > maxLength {
			maxLength = i-start + 1
		}

		lastOccoured[char] = i

	}

	fmt.Println(maxLength)
}
 