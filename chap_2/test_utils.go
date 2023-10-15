package chap2

type Case struct {
	input []int
	want  []int
}

func GetTestCases() []Case {
	return []Case{{
		input: []int{1, 2, 3, 4, 5},
		want:  []int{1, 2, 3, 4, 5},
	},
		{
			input: []int{},
			want:  []int{},
		},
		{
			input: []int{9, 8, 7, 5, -1, -2, 4, 0, 2, 2, 2},
			want:  []int{-2, -1, 0, 2, 2, 2, 4, 5, 7, 8, 9},
		},
		{
			input: []int{9, 9, 9, 9, 9},
			want:  []int{9, 9, 9, 9, 9},
		},
	}
}

func buildSortedSlice(inputSize int) []int {
	sorted := make([]int, inputSize)
	for i := 0; i < inputSize; i++ {
		sorted[i] = i
	}
	return sorted
}

func buildWorstCaseInput(inputSize int) []int {
	sortedDecreasing := make([]int, inputSize)
	for i := 0; i < inputSize; i++ {
		sortedDecreasing[i] = inputSize - i
	}
	return sortedDecreasing
}
