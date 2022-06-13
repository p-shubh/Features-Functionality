package main

import "fmt"

func main() {

	fmt.Println("Please! put the year ")

	var year int

	fmt.Scan(&year)

	if year%4 == 0 || year%400 == 0 {
		fmt.Println("Yes", year, "is a leap year")
	} else {
		fmt.Println(year, " is not a leap year")
	}
}
