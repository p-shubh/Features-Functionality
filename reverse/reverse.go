package main

import "fmt"

func main() {

	var reverse int = 23

	// var remainder int
	// var b int

	temp := reverse

	var remainder int

	if temp > 0 {
		remainder = temp % 10
		temp = temp / 10

		reverse = remainder*10 + temp
	}

	fmt.Println(reverse)

}
