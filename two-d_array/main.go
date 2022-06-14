package main

import "fmt"

func main() {

	// Example 1:
	// Input: accounts = [[1,2,3],[3,2,1]]
	// Output: 6
	// Explanation:
	// 1st customer has wealth = 1 + 2 + 3 = 6
	// 2nd customer has wealth = 3 + 2 + 1 = 6
	// Both customers are considered the richest with a wealth of 6 each, so return 6.

	// Example 2:
	// Input: accounts = [[1,5],[7,3],[3,5]]
	// Output: 10
	// Explanation:
	// 1st customer has wealth = 6
	// 2nd customer has wealth = 10
	// 3rd customer has wealth = 8
	// The 2nd customer is the richest with a wealth of 10.

	// Example 3:
	// Input: accounts = [[2,8,7],[7,1,3],[1,9,5]]
	// Output: 17

	Input := [][]int{{1, 2, 3}, {3, 2, 1}}

	// var extract_1 int
	// var extract_2 []int

	var result []int

	// var i,j int

	for i := 0; i < len(Input); i++ {
		var sum int
		// var sum1 int

		for j := 0; j < len(Input[i]); j++ {
			// fmt.Println(Input[i][j])
			fmt.Println()

			sum = (Input[i][j]) + sum

			// extract_1 = sum

			// fmt.Println(extract_1)

		}

		fmt.Println(sum)

		result = append(result, sum)
		// fmt.Println(result)

	}
	fmt.Println(result)

	var larger = result[0]
	for k := 0; k < len(result); k++ {
		if larger < result[k] {
			larger = result[k]
		}
	}

	fmt.Println("The largest wealth is ", larger)

	Input1 := [][]int{{1, 2, 3}, {3, 2, 1}}

	// var extract_1 int
	// var extract_2 []int

	var result2 []int

	// var i,j int

	for i := 0; i < len(Input1); i++ {
		var sum2 int
		// var sum1 int

		for j := 0; j < len(Input1[i]); j++ {
			// fmt.Println(Input1[i][j])
			fmt.Println()

			sum2 = (Input1[i][j]) + sum2

			// extract_1 = sum2

			// fmt.Println(extract_1)

		}

		fmt.Println(sum2)

		result2 = append(result2, sum2)
		// fmt.Println(result2)

	}
	fmt.Println(result2)

	var larger2 = result2[0]
	for k := 0; k < len(result2); k++ {
		if larger2 < result2[k] {
			larger2 = result2[k]
		}
	}

	fmt.Println("The largest wealth is ", larger2)

}
