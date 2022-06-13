package main

import "fmt"

func main() {

	fmt.Println("Print the series : 1 1 2 3 5 8 13 ....")

	series(8)

	// count := 7

}

func series(a int) {

	var j int = 1

	var store int = 1

	var print int

	for i := 0; i < a; i++ {
		fmt.Print(store, ",")
		print = store + j

		// store = 1 + 1 = 2
		// if j == 1 {
		// 	j = store
		// } // j = 2

		j = store
		store = print

	}

	fmt.Print("....")
	fmt.Println()

}
