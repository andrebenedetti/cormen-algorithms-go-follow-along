package chap2

import "fmt"

// Assumes the slice of the input bounded
// by start, middle and end are locally sorted */
func Merge(input []int, start int, middle int, end int) {
	fmt.Printf("Merging. Start: %d, middle: %d, end: %d\n", start, middle, end)

	left := make([]int, middle-start)
	copy(left, input[:middle])
	right := make([]int, end-middle)
	copy(right, input[middle:end])

	fmt.Printf("Merging %x and %x\n", left, right)
	l := 0
	r := 0
	result := make([]int, end-start)
	for l < len(left) || r < len(right) {
		if l >= len(left) {
			input[start+l+r] = right[r]
			result[r+l] = right[r]
			r++
		} else if r >= len(right) {
			input[start+l+r] = left[l]
			result[r+l] = left[l]
			l++
		} else if left[l] < right[r] {
			input[start+l+r] = left[l]
			result[r+l] = left[l]
			l++
		} else {
			input[start+l+r] = right[r]
			result[r+l] = right[r]
			r++
		}
	}
	fmt.Printf("Result: %x\n", result)
}

// 'start' is included, 'end' is excluded
func MergeSort(input []int, start int, end int) {
	fmt.Printf("Merge sort: start: %d, end: %d\n", start, end)
	// do nothing if start == and, as a slice with one

	// element is already sorted.
	if start < end {
		middle := (start + end) / 2
		MergeSort(input, start, middle)
		MergeSort(input, middle+1, end)
		Merge(input, start, middle, end)
	}
}
