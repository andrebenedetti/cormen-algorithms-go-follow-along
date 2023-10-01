package main

import (
	"fmt"

	chap2 "example.com/chap_2"
)

func main() {
	slice := []int{5, 4, 6, 1, 3, 9, 9, 12, 5}
	chap2.InsertionSort(&slice)
	fmt.Println(slice)

}
