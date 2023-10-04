package chap2

import "fmt"

/* Assumes the slice of the input bounded
by start, middle and end are locally sorted */

func Merge(input []int, start int, middle int, end int) {

}

// 'start' is included, 'end' is excluded
func MergeSort(input []int, start int, end int) {
	// do nothing if start == and, as a slice with one
	// element is already sorted.
	if start < end {
		fmt.Printf("Merging with start %d and end %d\n", start, end)
		// Split in half
		// middle is rounded down
		middle := (start + end) / 2
		fmt.Printf("Middle %d\n", middle)
		MergeSort(input, start, middle)
		MergeSort(input, middle+1, end)
	}
}
