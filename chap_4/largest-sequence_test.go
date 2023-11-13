package chap_4

import "testing"

type input struct {
	start    int
	end      int
	seq      []int
	expected []int
}

func TestLargestMiddleCrossingSequence(t *testing.T) {
	testCases := []input{
		{
			seq:      []int{1, 2, 3, 4, 5},
			expected: []int{0, 5, 15},
		},
		{
			seq:      []int{-1, 3, -2, 4, 5},
			expected: []int{1, 5, 10},
		},
		{
			seq:      []int{-1, -1, -1, -1, -1},
			expected: []int{1, 3, -2},
		},
	}

	for _, testCase := range testCases {
		start, end, sum := largestMiddleCrossingSequence(testCase.seq, 0, len(testCase.seq))
		if start != testCase.expected[0] {
			t.Fatalf("Expected start to be %d. Got %d", testCase.expected[0], start)
		}
		if end != testCase.expected[1] {
			t.Fatalf("Expected end to be %d. Got %d", testCase.expected[1], end)
		}
		if sum != testCase.expected[2] {
			t.Fatalf("Expected sum to be %d. Got %d", testCase.expected[2], sum)
		}
	}
}

func TestLargestSequence(t *testing.T) {
	testCases := []input{
		{
			seq:      []int{5, 5, 5, -5, -5, 5, 5},
			expected: []int{0, 7, 15},
		},
		{
			seq:      []int{5, 5, 5, -5, -5, 5, 5},
			expected: []int{0, 7, 15},
		},
		{
			seq:      []int{2, -2, -3, -2, 1},
			expected: []int{0, 1, 2},
		},
		{
			seq:      []int{1, -2, -3, -2, 2},
			expected: []int{4, 5, 2},
		},
	}

	for _, testCase := range testCases {
		start, end, sum := largestSequence(testCase.seq, 0, len(testCase.seq))
		if start != testCase.expected[0] {
			t.Fatalf("Expected start to be %d. Got %d", testCase.expected[0], start)
		}
		if end != testCase.expected[1] {
			t.Fatalf("Expected end to be %d. Got %d", testCase.expected[1], end)
		}
		if sum != testCase.expected[2] {
			t.Fatalf("Expected sum to be %d. Got %d", testCase.expected[2], sum)
		}
	}
}
