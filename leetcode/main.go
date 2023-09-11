package main

import "fmt"

func main() {
	longestPalindromicSubstring()
}

func threeSum() {

}

func longestPalindromicSubstring() {
	input := "babad"
	result := ""
	resLen := 0

	for i := 0; i < len(input); i++ {
		// odd length
		l, r := i, i
		for l >= 0 && r < len(input) && input[l] == input[r] {
			if (r - l + 1) > resLen {
				result = input[l : r+1]
				resLen = r - l + 1
			}

			l--
			r++
		}
	}

	fmt.Println(result)
}
