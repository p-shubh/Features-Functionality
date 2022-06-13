package main

import "fmt"

func main() {

	var a []int

	var b []int

	var d []int

	c := [...]int{1, 23, 41, 4, 1, 4, 1, 6}

	for i := 0; i < len(c); i++ {
		if c[i] != 1 {
			a = append(a, c[i])
		}
	}

	for i := 0; i < len(c); i++ {
		if c[i] == 1 {
			b = append(b, 1)
		}
	}

	d = append(a, b...)

	fmt.Println(d)

}
