package main

import "fmt"

func main() {

	// temp := 0

	// var a string

	var i int

	for i = 2; i <= 100; i++ {
		count := 0

		for j := 1; j <= i; j++ {
			if i%j == 0 && i%1 == 0 { // && i%1 == i
				count = count + 1 // fmt.Println(i)
			}

		}
		if count == 2 {
			fmt.Println(i)

		}

	}

	// temp = temp + 1

}
