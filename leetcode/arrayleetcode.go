package main

import (
	"fmt"
)

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

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	searchVal := 6 // should return 3
	start := 0
	end := len(nums) - 1

	value := -1

	for start <= end {
		mid := (start + end) / 2 // midIndex

		if nums[mid] == searchVal {
			value = mid
			break
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

	nums := []int{3, 4, 5, 1, 2}
	left := 0
	right := len(nums) - 1

	for left < right {
		midValue := (left + right) / 2

		// midValue is bigger than the right most value
		if nums[midValue] > nums[right] {
			left = midValue + 1
		} else {
			right = midValue
		}
	}

	fmt.Println(nums[left])
}
