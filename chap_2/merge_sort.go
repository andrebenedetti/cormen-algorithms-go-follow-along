package chap2

import "fmt"

func merge(input []int, start int, middle int, end int) {
	// 0, 0, 0
	if start == end {
		return
	}
	// middle will always be rounded down
	// 0, 0, 1 possible [3, 2] should split into [3] [2]
	// 0, 1, 1 not possible
	// 0, 1, 2 possible  [3, 1, 2] should split into [3, 2] [1]
	// 3, 3, 4
	left := make([]int, middle-start)
	for i := 0; i < len(left); i++ {
		left[i] = input[start+i]
	}
	fmt.Printf("Left: %x\n", left)

	right := make([]int, end-middle+1)
	for i := 0; i < len(right); i++ {
		right[i] = input[middle+i]
	}

	fmt.Printf("Right: %x\n", right)
}

func Run() {
	fmt.Println("merge sort")

	input := []int{1, 2}
	start := 0
	end := len(input)
	middle := (start + end) / 2
	merge(input, 0, middle, end-1)
}
