package main

import (
	"fmt"
)

func main() {

	// 	Example 1:
	// Input: word1 = ["ab", "c"], word2 = ["a", "bc"]
	// Output: true
	// Explanation:
	// word1 represents string "ab" + "c" -> "abc"
	// word2 represents string "a" + "bc" -> "abc"
	// The strings are the same, so return true.

	// Example 2:
	// Input: word1 = ["a", "cb"], word2 = ["ab", "c"]
	// Output: false

	// Example 3:
	// Input: word1  = ["abc", "d", "defg"], word2 = ["abcddefg"]
	// Output: true

	var word1 = []string{"ab", "c"}
	var word2 = []string{"a", "bc"}

	// for extraction

	var a string
	var b string

	for i := 0; i < len(word1); i++ {
		a = a + word1[i]
		// fmt.Println(a)
	}

	for j := 0; j < len(word2); j++ {
		b = b + word2[j]
		// fmt.Println(b)
	}

	fmt.Println(a == b)

}
