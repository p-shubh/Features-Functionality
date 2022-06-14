package main

import "fmt"

func main() {

	Input := [][]int{{1, 2, 3}, {3, 2, 1}}

	fmt.Println("the largest wealth is ", largestWealth(Input))

}

func largestWealth(inputs [][]int) int {

	ans := 0

	for i := 0; i < len(inputs); i++ {
		// now store the elements
		elements := inputs[i]
		total := 0
		for j := 0; j < len(elements); j++ {
			total = total + elements[j]
		}

		if ans < total {
			ans = total
		}
	}

	return ans
}
