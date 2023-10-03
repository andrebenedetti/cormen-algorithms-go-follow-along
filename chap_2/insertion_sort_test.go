package chap2

import (
	"reflect"
	"testing"
)

type Case struct {
	input []int
	want  []int
}

func TestInsertionSort(t *testing.T) {
	cases := []Case{{
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

	for _, testCase := range cases {
		InsertionSort(&testCase.input)
		if !reflect.DeepEqual(testCase.input, testCase.want) {
			t.Fatalf("want: %v, got: %v ", testCase.want, testCase.input)
		}
	}

}

// Chapter 2 shows Insertion Sort's best case scenario
// should have O(n) (a.k.a linear) time complexity)
// Lets use Go's built-in benchmarking tools to verify that.

func buildSortedSlice(inputSize int) []int {
    sorted := make([]int, inputSize)
	for i := 0; i < inputSize; i++ {
		sorted[i] = i
	}
    return sorted
}

func BenchmarkBestCase1(b *testing.B) {
	input := buildSortedSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort(&input)
	}
}

func BenchmarkBestCase2(b *testing.B) {
	input := buildSortedSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort(&input)
	}
}

func BenchmarkBestCase3(b *testing.B) {
	input := buildSortedSlice(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort(&input)
	}
}

func BenchmarkBestCase4(b *testing.B) {
	input := buildSortedSlice(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort(&input)
	}
}

// We also learned that Insertion Sort's worst case
// scenario's time complexity is O(n^2). Let's check
// that

func buildWorstCaseInput(inputSize int)[]int {
	sortedDecreasing := make([]int, inputSize)
	for i := 0; i < inputSize; i++ {
		sortedDecreasing[i] = inputSize - i
	}
    return sortedDecreasing
}

func BenchmarkInsertionSortWorstCase1(b *testing.B) {
    input := buildWorstCaseInput(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort(&input)
	}
}

func BenchmarkInsertionSortWorstCase2(b *testing.B) {
    input := buildWorstCaseInput(20000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort(&input)
	}
}

func BenchmarkInsertionSortWorstCase3(b *testing.B) {
    input := buildWorstCaseInput(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort(&input)
	}
}

func BenchmarkInsertionSortWorstCase4(b *testing.B) {
    input := buildWorstCaseInput(40000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort(&input)
	}
}
