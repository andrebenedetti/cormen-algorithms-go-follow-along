package chap2

import "fmt"

// Assumes the slice of the input bounded
// by start, middle and end are locally sorted */
func Merge(i *[]int, start int, middle int, end int) {
	left := make([]int, middle-start)
	copy(left, (*i)[:middle])
	right := make([]int, end-middle)
	copy(right, (*i)[middle:end])

	l := 0
	r := 0
	for l < len(left) || r < len(right) {
		if l >= len(left) {
			(*i)[start+l+r] = right[r]
			r++
		} else if r >= len(right) {
			(*i)[start+l+r] = left[l]
			l++
		} else if left[l] < right[r] {
			(*i)[start+l+r] = left[l]
			l++
		} else {
			(*i)[start+l+r] = right[r]
			r++
		}
	}
}

// 'start' is included, 'end' is excluded
func MergeSort(input []int, start int, end int) {
	// do nothing if start == and, as a slice with one
	// element is already sorted.
	if start < end {
		fmt.Printf("Merging with start %d and end %d\n", start, end)
		// Split in half
		// middle is rounded down

		// 0 (inclusive) + 4 (not inclusive) / 2 = 2
		//
		middle := (start + end) / 2
		fmt.Printf("Middle %d\n", middle)
		MergeSort(input, start, middle)
		MergeSort(input, middle+1, end)
	}
}
