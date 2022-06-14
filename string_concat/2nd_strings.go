package main

import (
	"fmt"
	"strings"
)

func main() {
	var word1 = []string{"ab", "c"}
	var word2 = []string{"a", "bc"}

	fmt.Println(strings.Join(word1, "") == strings.Join(word2, ""))
}
