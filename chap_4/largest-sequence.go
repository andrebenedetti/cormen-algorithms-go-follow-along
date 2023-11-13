package chap_4

import (
	"math"
)

// Chap 4 explores a divide and conquer algorithm to find the largest-sum sequence
// in a list of numbers.
// Although its time complexity is not the best compared to a sliding window method,
// it is still better than the bruteforce approach that has quadratic time.

func largestMiddleCrossingSequence(seq []int, start int, end int) (int, int, int) {
	leftSum := math.Inf(-1)
	middle := (end + start) / 2
	localSum := 0
	leftIndex := middle - 1
	for i := middle - 1; i >= start; i-- {
		localSum += seq[i]

		if localSum >= int(leftSum) {
			leftSum = float64(localSum)
			leftIndex = i
		}
	}

	localSum = 0
	rightIndex := middle
	rightSum := math.Inf(-1)
	for i := middle; i < end; i++ {
		localSum += seq[i]

		if localSum >= int(rightSum) {
			rightSum = float64(localSum)
			rightIndex = i
		}
	}
	return leftIndex, rightIndex + 1, int(leftSum + rightSum)
}

// 'start' is the start index, included
// 'end' is the end index, excluded.
// Doesn't work if all numbers are < 0
func largestSequence(seq []int, start int, end int) (int, int, int) {
	if end-start == 1 {
		return start, end, seq[start]
	}

	middle := (start + end) / 2
	leftStart, leftEnd, leftSum := largestSequence(seq, start, middle)
	rightStart, rightEnd, rightSum := largestSequence(seq, middle, end)
	crossingStart, crossingEnd, crossingSum := largestMiddleCrossingSequence(seq, start, end)

	if leftSum > rightSum && leftSum > crossingSum {
		return leftStart, leftEnd, leftSum
	}

	if rightSum > leftSum && rightSum > crossingSum {
		return rightStart, rightEnd, rightSum
	}

	return crossingStart, crossingEnd, crossingSum
}
