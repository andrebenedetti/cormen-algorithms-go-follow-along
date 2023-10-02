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

func BenchmarkInsertionSortBestCase(b *testing.B) {
	inputSize := 1000000
	alreadySorted := make([]int, inputSize)
	for i := 0; i < inputSize; i++ {
		alreadySorted[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort(&alreadySorted)
	}
}

func BenchmarkInsertionSortWorstCase(b *testing.B) {
	inputSize := 1000000
	sortedDecreasing := make([]int, inputSize)
	for i := inputSize - 1; i >= 0; i-- {
		sortedDecreasing[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort(&sortedDecreasing)
	}
}
