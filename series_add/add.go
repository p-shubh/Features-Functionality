package main

import "fmt"

func main() {

	var add int

	// 	Input: nums = [1,2,3,4]
	// Output: [1,3,6,10]
	// Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].

	input := []int{1, 2, 3, 4}

	var output []int

	for i := 0; i < len(input); i++ {

		add = input[i] + add
		output = append(output, add)

	}

	fmt.Println(output)

}
