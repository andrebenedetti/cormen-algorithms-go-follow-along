package main

import (
	"fmt"

	chap2 "example.com/chap_2"
)

func main() {
	s1 := []int{5, 3, 1, 2, 3, 6, 7, 8, 10}
	fmt.Println(chap2.MergeSort(s1))
}
