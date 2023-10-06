package main

import (
	"fmt"

	chap2 "example.com/chap_2"
)

func main() {
	input := []int{4, 3, 2}
	chap2.MergeSort(input, 0, len(input))
	fmt.Println(input)
}

// 3, 2, 1

// [3] ..  [2, 1]
// ou [3, 2] .. [1]
