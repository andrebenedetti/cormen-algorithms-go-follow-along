package main

import (
	"fmt"

	"example.com/algorithms"
)

func main() {
	a := []int{1, 2, 3, 4}
	results := make([][]int, 0)

	algorithms.Generate(len(a), a, &results)

	fmt.Println(results)

}
