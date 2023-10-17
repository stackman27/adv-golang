package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Stacks Leetcode

func ValidParentheses() {
	/*
		Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

		An input string is valid if:

		Open brackets must be closed by the same type of brackets.
		Open brackets must be closed in the correct order.
		Every close bracket has a corresponding open bracket of the same type.


		Example 1:

		Input: s = "()"
		Output: true
		Example 2:

		Input: s = "()[]{}"
		Output: true
		Example 3:

		Input: s = "(]"
		Output: false
	*/

	s := "()["
	closeToOpen := make(map[rune]rune)

	stack := []rune{} // Last in first out

	closeToOpen[')'] = '('
	closeToOpen[']'] = '['
	closeToOpen['}'] = '{'
	output := true

	for _, char := range s {
		// if the last element in stack is anything from the map
		_, ok := closeToOpen[char]
		if ok {
			if len(stack) != 0 && stack[len(stack)-1] == closeToOpen[char] {
				stack = stack[:len(stack)-1]
			} else {
				output = false
				break
			}
		} else {
			stack = append(stack, char)
		}
	}

	// return true only if stack is empty
	if len(stack) != 0 {
		output = false
	}

	fmt.Println(output)
}

func SimplifyPath() {
	/*
		Given a string path, which is an absolute path (starting with a slash '/') to a file or directory in a Unix-style file system, convert it to the simplified canonical path.

		In a Unix-style file system, a period '.' refers to the current directory, a double period '..' refers to the directory up a level, and any multiple consecutive slashes (i.e. '//')
		are treated as a single slash '/'. For this problem, any other format of periods such as '...' are treated as file/directory names.

		The canonical path should have the following format:

		The path starts with a single slash '/'.
		Any two directories are separated by a single slash '/'.
		The path does not end with a trailing '/'.
		The path only contains the directories on the path from the root directory to the target file or directory (i.e., no period '.' or double period '..')
		Return the simplified canonical path.



		Example 1:

		Input: path = "/home/"
		Output: "/home"
		Explanation: Note that there is no trailing slash after the last directory name.
		Example 2:

		Input: path = "/../"
		Output: "/"
		Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
		Example 3:

		Input: path = "/home//foo/"
		Output: "/home/foo"
		Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
	*/

	path := "/home/foo/../" // output => /home
	var curr string
	stack := []string{}

	for _, char := range path {
		if char == '/' {
			if curr == ".." {
				// we want to go back, pop the last element from stack
				if len(stack) != 0 {
					stack = stack[:len(stack)-1]
				}
			} else if curr != "" && curr != "." {
				stack = append(stack, curr)
			}
			curr = ""

		} else {
			// this is a text based directory
			curr += string(char)
		}
	}

	result := "/" + strings.Join(stack, "/") // Join and add the leading slash

	fmt.Println(result)
}

// TODO: look into backtracking solution: https://www.youtube.com/watch?v=s9fokUqJ76A&list=PLot-Xpze53lfxD6l5pAGvCD4nPvWKU8Qo&index=4
func GenerateParentheses() {
	/*
		Given n pairs of parentheses, write a function to generate all combinations
		of well-formed parentheses.

		Example 1:

		Input: n = 3
		Output: ["((()))","(()())","(())()","()(())","()()()"]
		Example 2:

		Input: n = 1
		Output: ["()"]

	*/
}

/*
	Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

	Implement the MinStack class:

	MinStack() initializes the stack object.
	void push(int val) pushes the element val onto the stack.
	void pop() removes the element on the top of the stack.
	int top() gets the top element of the stack.
	int getMin() retrieves the minimum element in the stack.
	You must implement a solution with O(1) time complexity for each function.
*/

// LAST IN FIRST OUT
type MinStack struct {
	Arr      []int
	MinValue []int
}

func NewMinStack() *MinStack {
	return &MinStack{
		Arr:      []int{},
		MinValue: []int{math.MaxInt},
	}
}

// push appends to the end of the stack
func (ms *MinStack) push(i int) {
	ms.Arr = append(ms.Arr, i)

	if i < ms.MinValue[len(ms.MinValue)-1] {
		ms.MinValue = append(ms.MinValue, i)
	}
}

// pop removes from the first element
func (ms *MinStack) pop() {
	// value we are popping is actually the minValue
	if ms.Arr[0] == 0 {
		// there is no value to pop
		return
	}

	// popped value is minimum, so we remove it from minValueStack as well
	if ms.Arr[len(ms.Arr)-1] == ms.MinValue[len(ms.MinValue)-1] {
		ms.MinValue = ms.MinValue[:len(ms.MinValue)-1]
	}

	ms.Arr = ms.Arr[:len(ms.Arr)-1]
}

func (ms *MinStack) top() int {
	return ms.Arr[len(ms.Arr)-1] // top of the stack
}

func (ms *MinStack) getMin() int {
	return ms.MinValue[len(ms.MinValue)-1]
}

func MinStackImpl() {
	newMinStack := NewMinStack()
	// arr [3,0,2,4] . minArr = [3, 0]
	newMinStack.push(3)
	newMinStack.push(0)
	newMinStack.push(2)
	newMinStack.push(4)

	newMinStack.pop()
	newMinStack.pop()

	// [1, 2] => [2]
	fmt.Println(newMinStack)
}

// TODO Figure this out
func DailyTemperatures() {
	/*
		Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the
		number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible,
		keep answer[i] == 0 instead.

		Example 1:

		Input: temperatures = [73,74,75,71,69,72,76,73]
		Output: [1,1,4,2,1,1,0,0]
		Example 2:

		Input: temperatures = [30,40,50,60]
		Output: [1,1,1,0]
		Example 3:

		Input: temperatures = [30,60,90]
		Output: [1,1,0]

	*/
}

func AsteroidCollision() {
	/*
		We are given an array asteroids of integers representing asteroids in a row.

		For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left).
		Each asteroid moves at the same speed.

		Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode.
		Two asteroids moving in the same direction will never meet.

		Example 1:

		Input: asteroids = [5,10,-5]
		Output: [5,10]
		Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.
		Example 2:

		Input: asteroids = [8,-8]
		Output: []
		Explanation: The 8 and -8 collide exploding each other.
	*/

	asteroids := []int{-1, 3, 4, -3}
	stack := []int{}

	for _, val := range asteroids {
		// checks if stack has value, the asteroid is negative and the last value in stack is positive
		for len(stack) != 0 && val < 0 && stack[len(stack)-1] > 0 {
			// we are about to have a collision

			diff := val + stack[len(stack)-1]

			// check if the val is greater than stack value if so pop it
			if diff < 0 {
				// pop value from stack
				stack = stack[:len(stack)-1]
			} else if diff > 0 {
				val = 0
			} else {
				val = 0
				stack = stack[:len(stack)-1]
			}
		}

		if val != 0 {
			stack = append(stack, val)
		}
	}

	fmt.Println(stack)
}

// TODO: Go over this problem again (VERY INTERESTING)
func CarFleet() {
	/*
		Description: https://leetcode.com/problems/car-fleet/
	*/

	position := []int{0, 2, 4, 1}
	speed := []int{4, 2, 1, 6}

	positionSpeedMap := make(map[int]int)

	for i, value := range position {
		positionSpeedMap[value] = speed[i]
	}

	// Step 1: Extract keys from the map and store them in a slice
	keys := make([]int, 0, len(positionSpeedMap))
	for key := range positionSpeedMap {
		keys = append(keys, key)
	}

	// Step 2: Sort the slice of keys
	sort.Ints(keys)

	stack := []int{}
	target := 10
	// reverse direction loop
	for i := len(keys) - 1; i >= 0; i-- {
		val := (target - keys[i]) / positionSpeedMap[keys[i]] // key = position, positionSpeedMap[key] = speed
		stack = append(stack, val)

		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			// the cars have collided
			stack = stack[:len(stack)-1]
		}
	}

	fmt.Println(len(stack))
}

func OnlineStockSpan() {
	/*
		Design an algorithm that collects daily price quotes for some stock and returns the span of that stock's price for the current day.

		The span of the stock's price in one day is the maximum number of consecutive days (starting from that day and going backward) for
		which the stock price was less than or equal to the price of that day.

		For example, if the prices of the stock in the last four days is [7,2,1,2] and the price of the stock today is 2, then the span of
		today is 4 because starting from today, the price of the stock was less than or equal 2 for 4 consecutive days.
		Also, if the prices of the stock in the last four days is [7,34,1,2] and the price of the stock today is 8, then the span of today
		is 3 because starting from today, the price of the stock was less than or equal 8 for 3 consecutive days.
		Implement the StockSpanner class:

		StockSpanner() Initializes the object of the class.
		int next(int price) Returns the span of the stock's price given that today's price is price.
		Example 1:

		Input
		["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
		[[], [100], [80], [60], [70], [60], [75], [85]]
		Output
		[null, 1, 1, 1, 2, 1, 4, 6]

		Explanation
		StockSpanner stockSpanner = new StockSpanner();
		stockSpanner.next(100); // return 1
		stockSpanner.next(80);  // return 1
		stockSpanner.next(60);  // return 1
		stockSpanner.next(70);  // return 2
		stockSpanner.next(60);  // return 1
		stockSpanner.next(75);  // return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.
		stockSpanner.next(85);  // return 6
	*/

	prices := []int{}
	spans := []int{}

	span := 1
	givenPrices := []int{100, 80, 60, 70, 60, 75, 85}

	for _, price := range givenPrices {
		// if the price is > last value in stack
		for len(prices) != 0 && price > prices[len(prices)-1] {

			// pop value from prices
			prices = prices[:len(prices)-1]
			span += spans[:len(spans)-1][0]
			spans = spans[:len(spans)-1]

		}

		prices = append(prices, price)
		spans = append(spans, span)

		fmt.Println(span)
	}
}

// Note: Really enjoyed solving this
func EvaluateReversePolishNotation() {
	/*
		You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

		Evaluate the expression. Return an integer that represents the value of the expression.

		Note that:

		The valid operators are '+', '-', '*', and '/'.
		Each operand may be an integer or another expression.
		The division between two integers always truncates toward zero.
		There will not be any division by zero.
		The input represents a valid arithmetic expression in a reverse polish notation.
		The answer and all the intermediate calculations can be represented in a 32-bit integer.


		Example 1:

		Input: tokens = ["2","1","+","3","*"]
		Output: 9
		Explanation: ((2 + 1) * 3) = 9
	*/

	tokens := []string{"4", "13", "5", "/", "+"}

	stack := []int{}

	for _, val := range tokens {
		if val == "+" || val == "*" || val == "-" || val == "/" {
			// we donot add to stack

			// get last two values from stack and do the operation
			if len(stack) >= 2 {
				lastStackValue := stack[len(stack)-1]
				secondLastStackValue := stack[len(stack)-2]
				var result int
				if val == "+" {
					// do add
					result = secondLastStackValue + lastStackValue
				} else if val == "*" {
					// do multiply
					result = secondLastStackValue * lastStackValue
				} else if val == "-" {
					// do subtract
					result = secondLastStackValue - lastStackValue
				} else if val == "/" {
					// do divide
					result = secondLastStackValue / lastStackValue
				} else {
					fmt.Println("Invalid operation ")
					return
				}

				// pop element from the stack
				stack = stack[:len(stack)-1]
				stack = stack[:len(stack)-1]

				// get the result and add it to stack
				stack = append(stack, result)
			}
		} else {

			val, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println("Encountrered error")
				return
			}

			stack = append(stack, val)
		}
	}

	fmt.Println(stack)
}

func RemoveKDigits() {
	/*
		Given string num representing a non-negative integer num, and an integer k, return the smallest possible
		integer after removing k digits from num.

		Example 1:

		Input: num = "1432219", k = 3
		Output: "1219"
		Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
		Example 2:

		Input: num = "10200", k = 1
		Output: "200"
		Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
	*/

	num := "1432219"
	stack := []int{} // Monotonic increasing stack
	limit := 3
	for _, val := range num {
		valNum, err := strconv.Atoi(string(val))
		if err != nil {
			fmt.Println("Error converting string to a num")
			return
		}

		if len(stack) > 0 {
			if valNum > stack[len(stack)-1] {
				stack = append(stack, valNum)
			} else {
				// we want to pop here
				// if the number is less than the stack last value
				for len(stack) != 0 && stack[len(stack)-1] >= valNum && limit > 0 {
					// pop the last stack Value
					stack = stack[:len(stack)-1]
					limit--
				}
				stack = append(stack, valNum)
			}
		} else {
			stack = append(stack, valNum)
		}
	}

	fmt.Println(stack)
}

// TODO do this
func MaximumFrequencyStack() {
	/*
		Design a stack-like data structure to push elements to the stack and pop the most frequent element from the stack.

		Implement the FreqStack class:

		FreqStack() constructs an empty frequency stack.
		void push(int val) pushes an integer val onto the top of the stack.
		int pop() removes and returns the most frequent element in the stack.
		If there is a tie for the most frequent element, the element closest to the stack's top is removed and returned.
	*/
}

func NextGreaterElement() {
	/*
		The next greater element of some element x in an array is the first greater element that is to the
		right of x in the same array.

		You are given two distinct 0-indexed integer arrays nums1 and nums2, where nums1 is a subset of nums2.

		For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j] and determine the
		next greater element of nums2[j] in nums2. If there is no next greater element, then the answer
		for this query is -1.

		Return an array ans of length nums1.length such that ans[i] is the next greater element as described
		above.

		Example 1:

		Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
		Output: [-1,3,-1]
		Explanation: The next greater element for each value of nums1 is as follows:
		- 4 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
		- 1 is underlined in nums2 = [1,3,4,2]. The next greater element is 3.
		- 2 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
	*/

	nums1 := []int{2, 4}
	nums2 := []int{1, 2, 3, 4}
	result := make([]int, len(nums1))

	for i := 0; i < len(nums1); i++ {
		result[i] = -1
	}

	hashmap := make(map[int]int)
	for i, val := range nums2 {
		hashmap[i] = val
	}

	fmt.Println("HASHMAP: ", hashmap)

	for i := range nums2 {
		for j := range nums1 {
			if hashmap[i] == nums1[j] {
				// check to see next value in nums2 is greater if not do nothing
				if hashmap[i+1] > hashmap[i] {
					result[j] = hashmap[i+1]
				}
			}
		}
	}

	fmt.Println(result)

}

// Look into tree traversal: https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/#
func BinarySearchTreeIterator() {

}

func BaseBallGame() {
	/*
		You are keeping the scores for a baseball game with strange rules. At the beginning of the game, you start with an empty record.

		You are given a list of strings operations, where operations[i] is the ith operation you must apply to the record and is one of the following:

		An integer x.
		Record a new score of x.
		'+'.
		Record a new score that is the sum of the previous two scores.
		'D'.
		Record a new score that is the double of the previous score.
		'C'.
		Invalidate the previous score, removing it from the record.
		Return the sum of all the scores on the record after applying all the operations.

		The test cases are generated such that the answer and all intermediate calculations fit in a 32-bit integer and that all operations are valid.



		Example 1:

		Input: ops = ["5","2","C","D","+"]
		Output: 30
		Explanation:
		"5" - Add 5 to the record, record is now [5].
		"2" - Add 2 to the record, record is now [5, 2].
		"C" - Invalidate and remove the previous score, record is now [5].
		"D" - Add 2 * 5 = 10 to the record, record is now [5, 10].
		"+" - Add 5 + 10 = 15 to the record, record is now [5, 10, 15].
		The total sum is 5 + 10 + 15 = 30.
	*/
}

func RemoveAdjacentDuplicateString() {
	/*
		You are given a string s consisting of lowercase English letters. A duplicate removal consists of choosing two adjacent and equal letters
		and removing them.

		We repeatedly make duplicate removals on s until we no longer can.
		Return the final string after all such duplicate removals have been made. It can be proven that the answer is unique.


		Example 1:

		Input: s = "abbaca"
		Output: "ca"
		Explanation:
		For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.
		The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".
		Example 2:

		Input: s = "azxxzy"
		Output: "ay"
	*/

	s := "azxxzy"

	stack := []rune{}

	for _, val := range s {
		if len(stack) > 0 && val == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, val)
		}
	}

	fmt.Println(string(stack))
}

func Pattern_123() {
	/*
		Given an array of n integers nums, a 132 pattern is a subsequence of three integers nums[i],
		nums[j] and nums[k] such that i < j < k and nums[i] < nums[k] < nums[j].

		Return true if there is a 132 pattern in nums, otherwise, return false.

		Example 1:

		Input: nums = [1,2,3,4]
		Output: false
		Explanation: There is no 132 pattern in the sequence.
		Example 2:

		Input: nums = [3,1,4,2]
		Output: true
		Explanation: There is a 132 pattern in the sequence: [1, 4, 2].
	*/

	nums := []int{3, 1, 4, 2}
	found := false

	fmt.Println(nums)
	fmt.Println(found)

}
