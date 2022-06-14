package main

import "fmt"

func main() {

	var add int
	var add2 int
	var add3 int

	// 	Example 1:
	// Input: nums = [1,2,3,4]
	// Output: [1,3,6,10]
	// Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].

	// Example 2:
	// Input: nums = [1,1,1,1,1]
	// Output: [1,2,3,4,5]
	// Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].

	// Example 3:
	// Input: nums = [3,1,2,10,1]
	// Output: [3,4,6,16,17]

	input1 := []int{1, 2, 3, 4}

	var output []int //it should be in slice because the value attached to it is in slice condition

	for i := 0; i < len(input1); i++ {

		add = input1[i] + add //here the input is in slice condition do the add no become slice
		output = append(output, add)
		// fmt.Print(output)
	}

	fmt.Println("output:=", output)

	// =====================================================

	input2 := []int{1, 1, 1, 1}

	var output2 []int

	for i := 0; i < len(input2); i++ {

		add2 = input2[i] + add2
		output2 = append(output2, add2)
		// fmt.Print(output2)
	}

	fmt.Println("output2:=", output2)

	// =====================================================

	input3 := []int{3, 1, 2, 10, 1}

	var output3 []int

	for i := 0; i < len(input3); i++ {

		add3 = input3[i] + add3
		output3 = append(output3, add3)
		// fmt.Print(output3)
	}

	fmt.Println("output3:=", output3)

}
