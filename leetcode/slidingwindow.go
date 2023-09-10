package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)
func SlidingWindow1() {
	/*
	You are given an array prices where prices[i] is the price of a given stock on the ith day.
	You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
	Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.


	Example 1:

	Input: prices = [7,1,5,3,6,4]
	Output: 5 
	Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
	Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

	*/

	arr := []int{7,1,5,3,6,4}
	maxProfit := 0
	leftPtr := 0
	rightPtr := 1

	// profit = right - left 
	for rightPtr < len(arr) {
		profit := arr[rightPtr] - arr[leftPtr]
		
		// arr[leftPtr] is the lowest min value for buy
		if arr[leftPtr] > arr[rightPtr] {
			leftPtr = rightPtr
		}

		if profit > maxProfit {
			maxProfit = profit
		}
		rightPtr ++ 
	}

	fmt.Println(maxProfit)
}

func SlidingWindow2() {
	/*
	Given two strings s and t of lengths m and n respectively, return the minimum window 
	substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

	The testcases will be generated such that the answer is unique.

	

	Example 1:

	Input: s = "ADOBECODEBANC", t = "ABC"
	Output: "BANC"
	Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
	*/

	// 3 - 0 + 1 < 7 
	// charCount[s[0]]++ 
	s := "ABDCAB" 
	t := "ABC"

	charCount := make(map[rune]int)
 

	for _, char := range t {
		charCount[char]++
	}

	start, minLen := 0, len(s)+1
	counter := len(t)
	result := ""

	for i, char := range s {
		if charCount[char] > 0 {
			counter--
		}
		charCount[char]-- 

		// counter == 0 means we have found A match
		for counter == 0 {
			if i-start+1 < minLen { 
				minLen = i - start + 1
				result = s[start : i+1]
			}
		 
			charCount[rune(s[start])]++

			if charCount[rune(s[start])] > 0 {
				counter++
			}
			start++ 
		}
	}

	fmt.Println(result)	 
}	

// TODO: confusing logic (rewatch video)
func SlidingWindow3() {
	/* 
	https://www.youtube.com/watch?v=gqXU1UyA8pk&list=PLot-Xpze53leOBgcVsJBEGrHPd_7x_koV&index=4
	You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character.
	You can perform this operation at most k times.

	Return the length of the longest substring containing the same letter you can get after performing the above operations.

	Example 1:

	Input: s = "ABAB", k = 2
	Output: 4
	Explanation: Replace the two 'A's with two 'B's or vice versa.

	Input: s = "AABABBA", k = 1
	Output: 4
	Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
	The substring "BBBB" has the longest repeating letters, which is 4.
	There may exists other ways to achive this answer too.


	"ABAB" 
	*/

	s := "AABABBA"
	k := 1
	start := 0 
	 

	maxFreq := float64(0)
	result := 0

	hashmap := make(map[rune]int)

	for i, char := range s {
		hashmap[char]++

		// find the max of window length and char count 
		maxFreq := math.Max(maxFreq, float64(hashmap[char]))
		freq := (i-start+1) - int(maxFreq)

		if freq <= k  {
			result += 1
		} else {
			// we slide the window 
			hashmap[char]--
			start++
			 
		}	
	}

	fmt.Println(result)
}

func SlidingWindow4() {
	/*
	438. Find All Anagrams in a String
	Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.

	An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

	Example 1:

	Input: s = "cbaebabacd", p = "abc"
	Output: [0,6]
	Explanation:
	The substring with start index = 0 is "cba", which is an anagram of "abc".
	The substring with start index = 6 is "bac", which is an anagram of "abc".
	Runtime: O(n)
	*/

	s := "cbaebabacd"
	p := "abc"

	pHashMap := make(map[rune]int)

	for _, val := range p {
		pHashMap[val]++
	}

	counter := len(p)
	start := 0
	indexes := []int{}

	for i, val := range s {
		if pHashMap[val] > 0 {
			counter -- 
		}

		pHashMap[val]--

		if counter == 0 {
			fmt.Println("ANAGRAM FOUND")
			// this is an anagram 
			// get the start index  
			indexes = append(indexes, start)
			  			
			pHashMap[rune(s[start])]++  
			if pHashMap[rune(s[start])] > 0 {
				counter++
			}

			start++
		}

		// Shrink the window if it exceeds the length of p
		if i-start+1 >= len(p) {
			if pHashMap[rune(s[start])] > 0 {
				counter++
			}
			pHashMap[rune(s[start])]++  
			start++
		}
	}

	fmt.Println(indexes)
}

// LEETCODE HARD
func SlidingWindow5() {
	/*
	Sliding Window Maximum
	You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of 
	the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right 
	by one position.
	Return the max sliding window.

	Example 1:

	Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
	Output: [3,3,5,5,6,7]
	Explanation: 
	Window position                Max
	---------------               -----
	[1  3  -1] -3  5  3  6  7       3
	1 [3  -1  -3] 5  3  6  7       3
	1  3 [-1  -3  5] 3  6  7       5
	1  3  -1 [-3  5  3] 6  7       5
	1  3  -1  -3 [5  3  6] 7       6
	1  3  -1  -3  5 [3  6  7]      7
	*/

	nums := []int{1,3,-1,-3,5,3,6,7}
	k := 3

	result := []int{}
	deque := []int{} // values are always decreasing  

	for i:= 0; i <len(nums); i++ {
		if len(deque) > 0 && deque[0] <= i-k { 
			// Removes first element from the deque
			deque = deque[1:] // Remove elements outside of the window 
		}

		for len(deque) > 0 && nums[i] >= nums[deque[len(deque) - 1]] { 
			// Removes the last element from the deque
			deque = deque[:len(deque) -1] // remove elements smaller than current element 
			fmt.Println(deque)
		}
		 
		deque = append(deque, i)

		if i >= k - 1 {
			result = append(result, nums[deque[0]]) // front of deque contains maximum element for current window
		}
	}

	fmt.Println(result)
}

// MIGHT BE INCORRECT: RECHECK
func SlidingWindow6() {
	/*
	Frequency of the most frequent element
	The frequency of an element is the number of times it occurs in an array.

	You are given an integer array nums and an integer k. In one operation, you can choose an index of 
	nums and increment the element at that index by 1.

	Return the maximum possible frequency of an element after performing at most k operations.

	Example 1:

	Input: nums = [1,2,4], k = 5
	Output: 3
	Explanation: Increment the first element three times and the second element two times to make nums = [4,4,4].
	4 has a frequency of 3.
	O(nlogn)
	*/
 
	nums :=  []int{1, 2, 2, 3}
	k := 2

	sort.Ints(nums)  

	lastNum := nums[len(nums) - 1]
	output := 1

	for i := len(nums) - 1; i >= 0; i -- {
		if k > 0 {
			// last element  
			secondLastNum := nums[i-1]
			diff := (lastNum - secondLastNum)
 
			// increment the second last number by the remainig k 			
			k =  k - diff   
			output ++		 
		}
	}

	fmt.Println(output)
}

// TODO: Do this 
func SlidingWindow7() {
	/*
	Minimum number of flips to make a binary String alternative

	You are given a binary string s. You are allowed to perform two types of operations on the string in any sequence:

	Type-1: Remove the character at the start of the string s and append it to the end of the string.
	Type-2: Pick any character in s and flip its value, i.e., if its value is '0' it becomes '1' and vice-versa.
	Return the minimum number of type-2 operations you need to perform such that s becomes alternating.

	The string is called alternating if no two adjacent characters are equal.

	For example, the strings "010" and "1010" are alternating, while the string "0100" is not.
	

	Example 1: 

	Input: s = "111000"
	Output: 2
	Explanation: Use the first operation two times to make s = "100011".
	Then, use the second operation on the third and sixth elements to make s = "101010".

	*/
} 

// NOTE: CURRENT SOLUTION log(nLogn) can be done in log(n)
func SlidingWindow8() {
	/*
	Permutation in String

	Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

	In other words, return true if one of s1's permutations is the substring of s2.

	

	Example 1:

	Input: s1 = "ab", s2 = "eidbaooo"
	Output: true
	Explanation: s2 contains one permutation of s1 ("ba").
	Example 2:

	Input: s1 = "ab", s2 = "eidboaoo"
	Output: false
	*/

	s1 := "ab"
	s2 := "eidboaoo"

	s1Len := len(s1)
	start := 0
	end := s1Len
	found := false

	s1 = SortString(s1) 
 
	for end <= len(s2) {
		subset := s2[start:end]
		subset = SortString(subset)
		// get the subset value 
		if subset == s1 {
			found = true; 
			break;
		}

		start ++
		end ++
	}
 
	fmt.Println(found)
}



func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func SlidingWindow9() {
	/*
	Minimum Size Subarray Sum
	Given an array of positive integers nums and a positive integer target, return the minimal length of a 
	subarray
	whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

	

	Example 1:

	Input: target = 7, nums = [2,3,1,2,4,3]
	Output: 2
	Explanation: The subarray [4,3] has the minimal length under the problem constraint.
	Example 2:

	Input: target = 4, nums = [1,4,4]
	Output: 1
	Example 3:

	Input: target = 11, nums = [1,1,1,1,1,1,1,1]
	Output: 0
	*/

	nums := []int{1,1,1,1,1,1,1,1}
	target := 11

	start := 0
	sum := 0
	output := 0

	for i, val := range nums {
		if val == target {
			output = 1
			break;
		}

		sum = sum + val
		 
		if sum == target {
			// store the window length
			if output == 0 || i-start + 1 < output { 
				output = i-start + 1
			}
			sum = sum - nums[start] 
			start ++ 
		}  
		
	 
		if sum > target { 
			for sum > target { // learn more about this for loop here
				sum = sum - nums[start] 
				start ++ 
			}

			if sum == target {
				if output == 0 ||  i-start + 1 < output {
					output = i-start + 1
				}
			}
		}  
	}

	fmt.Println(output)
}

func SlidingWindow10() {
	/*
	Repeated DNA Sequences

	The DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T'.

	For example, "ACGAATTCCG" is a DNA sequence.
	When studying DNA, it is useful to identify repeated sequences within the DNA.

	Given a string s that represents a DNA sequence, return all the 10-letter-long sequences (substrings) that occur more than once 
	in a DNA molecule. You may return the answer in any order.

	Example 1:

	Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	Output: ["AAAAACCCCC","CCCCCAAAAA"]
	Example 2:

	Input: s = "AAAAAAAAAAAAA"
	Output: ["AAAAAAAAAA"]
	*/


	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTTT"
	
	 
	dnaHashSet := make(map[string]bool)
	duplicates := make(map[string]bool)

	// loop for every value in s until there is less than 10 chars left
	for i:= 0; i < len(s) - 10; i++{
		seq := s[i:i+10]
			
		if dnaHashSet[seq] {
			duplicates[seq] = true
		} else {
			dnaHashSet[seq] = true
		} 
	}

	fmt.Println(duplicates)
}