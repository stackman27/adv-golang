package main

import (
	"fmt"
	"math"
	"strings"
)


func reverseString() {
	str := "abcde"
	var val string

	for i:=len(str) - 1; i >= 0; i-- { 
		val += str[i:i+1] // s[i] is a byte but if we do s[i:i+1] it is a slice
	}

	fmt.Println(val)
}

func twoStringArePalindrome() {
	var str string = "madam"
	str = strings.ToLower(str)
	isPalindrome := true 
	for i:= 0; i < len(str)/2; i ++ {
		if str[i] !=  str[len(str) - i - 1] {
			isPalindrome = false
		}
	}

	fmt.Println(isPalindrome)
}

func slidingWindowDynamic() {

	minLength := -1000
	arr := []int{1,2,3,4,5,6}
	x := 3
	
	// log(n) solution 

	start := 0 
	end := 0
	currentSum := 0 

	for (end < len(arr)) {
		currentSum = currentSum + arr[end]
		end = end + 1 

		for start < end && currentSum >= x {
			currentSum = currentSum - arr[start]
			start = start + 1
			

			minLength = int(math.Min(float64(minLength), float64(end-start+1)))
 
		}
	} 

	fmt.Println(minLength)
}

func findDups() {
	/**
		Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
		Find all the elements that appear twice in this array.

		Example:

		Input:
		[4,3,2,7,8,2,3,1]

		Output:
		[2,3] 
	**/

	 
	nums := []int{4,3,2,7,8,2,3,1}

	// nlogn solution 
	// sort.Slice(arr, func(i, j int) bool{
	// 	return arr[i] < arr[j]
	// })

	// fmt.Println(arr)


	// sameArr := []int{}

	// for i := 1; i <len(arr); i++ {
	// 	// we can do this if we can sort it 
	// 	if arr[i] == arr[i-1] {
	// 		sameArr = append(sameArr, arr[i])
	// 	}
	// }
	
	// log(n) solution 
	nums_map := make(map[int]bool)
	duplicates := []int{}

	for _, val := range nums {
		if nums_map[val] {
			duplicates = append(duplicates, val)
		} 
		nums_map[val] = true
 
	} 

	fmt.Println(duplicates)

}


