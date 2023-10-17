package main

import (
	"fmt"
	"math"
)

// BINARY SEARCH LEETCODE

// check this out: https://www.youtube.com/watch?v=YUF3_eBdzsk&list=PLot-Xpze53leNZQd0iINpD-MAhMOMzWvO
func SplitArrayLargestSum() {
	/*
	Given an integer array nums and an integer k, split nums into k non-empty subarrays such that the largest sum of any subarray is minimized.

	Return the minimized largest sum of the split.

	A subarray is a contiguous part of the array.

	Example 1:

	Input: nums = [7,2,5,10,8], k = 2
	Output: 18
	Explanation: There are four ways to split nums into two subarrays.
	The best way is to split it into [7,2,5] and [10,8], where the largest sum among the two subarrays is only 18.
	*/
}

func TypicalBinarySearch() { 
	/*
	Given an array of integers nums which is sorted in ascending order, and an integer target, 
	write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

	You must write an algorithm with O(log n) runtime complexity.

	

	Example 1:

	Input: nums = [-1,0,3,5,9,12], target = 9
	Output: 4
	Explanation: 9 exists in nums and its index is 4
	*/

	nums := []int{-1,0,3,5, 9,12} 
	target := 9 

	left := 0
	right := len(nums) 

	targetIdx := -1


	for left < right {
		mid :=  (right + left) / 2
 
		if target == nums[mid] {
			targetIdx = mid 
			break;
		}

		if target > nums[mid] {
			// search right
			left = mid + 1 
		} else {
			// search left 
			right = mid 
		}
	}

	fmt.Println(targetIdx)
}

func ValidPerfectSquare() {
	/* 
	Given a positive integer num, return true if num is a perfect square or false otherwise.

	A perfect square is an integer that is the square of an integer. In other words, it is the product of some integer with itself.

	You must not use any built-in library function, such as sqrt.

	Example 1:

	Input: num = 16
	Output: true
	Explanation: We return true because 4 * 4 = 16 and 4 is an integer.
	*/

	num := float64(16)
	isValid := false 

	left := float64(1) 
	right := float64(num) 

	for left <= right {
		mid := (left + right) / 2
		perfSquare := mid * mid 

	 
		if perfSquare > num {
			// Search left 
			right = mid -1
		} else if perfSquare < num {
			// Search right
			left = mid + 1
		} else {
			isValid = true 
			break;
		}

	}

	fmt.Println(isValid)
}

func ArrangingCoins() {
	/*
	Description: https://leetcode.com/problems/arranging-coins/
	*/

	n := 5 
	rowsBuilt := 0
	// BRUTEFORCE SOLUTION = O(N)
	// for n > 0 {  
	// 	n = n - i
	  
	// 	if n < 0 {
	// 		break;
	// 	}
	// 	rowsBuilt++
	// 	i++
	// }


	// OPTIMAL SOLUTION = log(n)
	// using Gauss adding formula (n(n+1)) / 2

	left := 1
	right := n
	for left <= right {
		mid := (left + right) / 2
		midSum := (mid * (mid+1)) / 2 //gives us how many rows we need to fill for this mid 
		
		if midSum > n{
			right = mid -1
		} else {
			left= mid + 1 
			// we have enough to fill everything 
			rowsBuilt = int(math.Max(float64(mid), float64(rowsBuilt)))
		}
	}
	 

	fmt.Println(rowsBuilt)
}

// TODO do this explanation: https://www.youtube.com/watch?v=o-YDQzHoaKM&list=PLot-Xpze53leNZQd0iINpD-MAhMOMzWvO&index=6
func FindKClosestElements() {
	/*
	Given a sorted integer array arr, two integers k and x, return the k closest integers to x in the array. 
	The result should also be sorted in ascending order.

	An integer a is closer to x than an integer b if:

	|a - x| < |b - x|, or
	|a - x| == |b - x| and a < b
	

	Example 1:

	Input: arr = [1,2,3,4,5], k = 4, x = 3
	Output: [1,2,3,4]
	Example 2:

	Input: arr = [1,2,3,4,5], k = 4, x = -1
	Output: [1,2,3,4]

	*/
}

type TimeMap struct {
	Data map[string][]KeyValue // map string with an array of keyValue
}

type KeyValue struct {
	Timestamp int 
	Values string
}

// GetTimeMap returns the value given a key and timestamp, if timestamp doesnot exist return the most recent one 
func (tm *TimeMap) GetTimeMap(key string, timestamp int) string{
	vals, ok := tm.Data[key]
	if !ok {
		return "Error, key doesnot exist"
	}

	found := false 

	for _, val := range vals {
		if val.Timestamp == timestamp {
			// exact match 
			found = true
			return val.Values
		} 
	}
	
	// TODO: Make this binary search (i believe timestamp is already sorted)
	if !found {
		// we didn't find the value so return the most recent one 
		closestTimestamp := -1
		closestValue := ""
		for _, val := range vals {
			if val.Timestamp <= timestamp && val.Timestamp > closestTimestamp {
				closestTimestamp = val.Timestamp
                closestValue = val.Values
			}
		}
		return closestValue
	}

	return ""
}

func (tm *TimeMap) SetTimeMap(key string, value string, timestamp int) {

	kv := KeyValue{
		Timestamp: timestamp,
		Values: value,
	}
  
	tm.Data[key] = append(tm.Data[key], kv)
}

func NewTimeMap() TimeMap{
	return TimeMap{Data: make(map[string][]KeyValue)}
}

// NOTE: IMPORTANT PROBLEMMM
func TimeBasedKVStorage() {
	/* 
	Design a time-based key-value data structure that can store multiple values for the same key at different 
	time stamps and retrieve the key's value at a certain timestamp.

	Implement the TimeMap class:

	TimeMap() Initializes the object of the data structure.
	- void set(String key, String value, int timestamp) Stores the key key with the value value at the given time timestamp.
	- String get(String key, int timestamp) Returns a value such that set was called previously, with timestamp_prev <= timestamp. 
	- If there are multiple such values, it returns the value associated with the largest timestamp_prev. If there are no values, it returns "".
	

	Example 1:

	Input
	["TimeMap", "set", "get", "get", "set", "get", "get"]
	[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
	Output
	[null, null, "bar", "bar", null, "bar2", "bar2"]

	Explanation
	TimeMap timeMap = new TimeMap();
	timeMap.set("foo", "bar", 1);  // store the key "foo" and value "bar" along with timestamp = 1.
	timeMap.get("foo", 1);         // return "bar"
	timeMap.get("foo", 3);         // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
	timeMap.set("foo", "bar2", 4); // store the key "foo" and value "bar2" along with timestamp = 4.
	timeMap.get("foo", 4);         // return "bar2"
	timeMap.get("foo", 5);         // return "bar2"

	*/ 

	timeMap := NewTimeMap()

	timeMap.SetTimeMap("Sishir", "foo", 123)
	timeMap.SetTimeMap("Sishir", "bar", 1234)

	val := timeMap.GetTimeMap("Sishir", 1234)
	fmt.Println(val)
}


// NOTE: We did this in ArrayLeetcode section
func FindMinInRotatedSortedArray() {
	/*
	Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

	[4,5,6,7,0,1,2] if it was rotated 4 times.
	[0,1,2,4,5,6,7] if it was rotated 7 times.
	Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

	Given the sorted rotated array nums of unique elements, return the minimum element of this array.

	You must write an algorithm that runs in O(log n) time. 
	*/
}


func Search2DMatrix() {
	/*
	You are given an m x n integer matrix matrix with the following two properties:

	Each row is sorted in non-decreasing order.
	The first integer of each row is greater than the last integer of the previous row.
	Given an integer target, return true if target is in matrix or false otherwise.

	You must write a solution in O(log(m * n)) time complexity.

	
	*/

	matrix := [][]int{
        {1, 3, 4, 5},
        {6, 7, 8},
        {9, 10, 11},
    }
	
	searchValue := 11
	found := false
	for _, vals := range matrix {
	 
		// binary search on vals
		left := 0 
		right := len(vals) -1
		
		for left <= right {
			mid := (right + left) / 2
 

			if  searchValue > vals[mid] {
				// search to the right 
				left = mid + 1 
			} else if searchValue < vals[mid] {
				// search left
				right = mid 
			} else {

			if vals[mid] == searchValue {
				// value found 
				found = true
				break;
			}
			}

		}
	}

	fmt.Println(found)
}

// TODO
func MaxNumberofRemovalChars() {
	/*
	You are given two strings s and p where p is a subsequence of s. You are also given a distinct 0-indexed integer array removable containing a subset of indices of s (s is also 0-indexed).

	You want to choose an integer k (0 <= k <= removable.length) such that, after removing k characters from s using the first k indices in removable, p is still a subsequence of s. 
	More formally, you will mark the character at s[removable[i]] for each 0 <= i < k, then remove all marked characters and check if p is still a subsequence.

	Return the maximum k you can choose such that p is still a subsequence of s after the removals.

	A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted 
	without changing the relative order of the remaining characters.


	Example 1:

	Input: s = "abcacb", p = "ab", removable = [3,1,0]
	Output: 2
	Explanation: After removing the characters at indices 3 and 1, "abcacb" becomes "accb".
	"ab" is a subsequence of "accb".
	If we remove the characters at indices 3, 1, and 0, "abcacb" becomes "ccb", and "ab" is no longer a subsequence.
	Hence, the maximum k is 2.
	*/

}

func FindFirstAndLastPosition() {
	/*
	Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
	If target is not found in the array, return [-1, -1].
	You must write an algorithm with O(log n) runtime complexity.


	Example 1:

	Input: nums = [5,7,7,8,8,10], target = 8
	Output: [3,4]
	Example 2:

	Input: nums = [5,7,7,8,8,10], target = 6
	Output: [-1,-1]
	Example 3:

	Input: nums = [], target = 0
	Output: [-1,-1]
	*/
	nums := []int {5,8,7,8,8,8,10} // 8,8,10
	//first := FindFirstAndLastPositionHelper(nums, 8, true)
	last := FindFirstAndLastPositionHelper(nums, 8, false)

	fmt.Println(last)
}

func FindFirstAndLastPositionHelper(nums []int, target int, leftBias bool) int{
	left := 0 
	right := len(nums)
	i := -1

	for left < right {
		mid :=  (right + left) / 2
  
		if target > nums[mid] {
			// search right
			left = mid + 1 
		} else if target < nums[mid] {
			// search left 
			right = mid
		} else {
			i = mid 
			if leftBias { // leftBias = true, BS will continue searching to the left 
				right = mid - 1  
			} else { // leftBias = false, BS will continue searching to the right 
				left = mid + 1 
			}
	 
		}
	}

	return i
}

func SearchInsertPosition() {
	/*
	Given a sorted array of distinct integers and a target value, return the index if the target is found. 
	If not, return the index where it would be if it were inserted in order.
	You must write an algorithm with O(log n) runtime complexity.

	Example 1:

	Input: nums = [1,3,5,6], target = 5
	Output: 2
	Example 2:

	Input: nums = [1,3,5,6], target = 2
	Output: 1
	Example 3:

	Input: nums = [1,3,5,6], target = 7
	Output: 4
	*/

	nums := []int{1,3,5,6}
	target := 7
	left := 0 
	right := len(nums)
	val := -1

	for left < right {
		mid := (left+right) / 2 

		if nums[mid] == target {
			val = mid
			break;
		} else if target < nums[mid] {
			// search left 
			right = mid
		} else {
			// search right 
			left = mid + 1
		}

		val = left
	}

	fmt.Println(val)
}