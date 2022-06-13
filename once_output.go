package main

import "fmt"

func main() {

	a := []int{2, 5, 5, 2, 7, 7, 8}

	// out(a)

	result := 0

	for i := 0; i < len(a); i++ {
		var count int = 0
		for j := 0; j < len(a); j++ {

			if a[i] == a[j] && i != j {

				count = count + 1
			}

		}

		if count == 0 {
			result = a[i]
		}

	}

	fmt.Println(result)

}
