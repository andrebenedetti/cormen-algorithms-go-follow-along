package chap2

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	cases := []Case{
		{
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
		result := MergeSort(testCase.input)
		if !reflect.DeepEqual(result, testCase.want) {
			t.Fatalf("want: %v, got: %v ", testCase.want, testCase.input)
		}
	}

}

// We should expect Merge Sort's best performance in the best case
// to be worse than Insertion Sort's best case, since the latter is
// O(n) while Merge Sort is still O(n log(n))
// By running this test, we can see that merge sort performs way worse for
// this input size in the best case scenario
func BenchmarkMergeSortBestCase1(b *testing.B) {
	input := buildSortedSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSort(input)
	}
}

// We should see gains in the worst case when comparing to InsertionSort, but we do not
// In fact, the implementation of MergeSort being tested here is way worse than the InsertionSort
// implementation. Let's see if we can improve our MergeSort to actually make it perform better.
// The benchmark says that this implementation does 33851 allocs/op, while our insertion sort
// does 0. This should be the source of the problems.
func BenchmarkMergeSortWorstCase1(b *testing.B) {
	input := buildWorstCaseInput(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSort(input)
	}
}

// func BenchmarkInsertionSortWorstCase2(b *testing.B) {
// 	input := buildWorstCaseInput(20000)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		InsertionSort(&input)
// 	}
// }

// func BenchmarkInsertionSortWorstCase3(b *testing.B) {
// 	input := buildWorstCaseInput(30000)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		InsertionSort(&input)
// 	}
// }

// func BenchmarkInsertionSortWorstCase4(b *testing.B) {
// 	input := buildWorstCaseInput(40000)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		InsertionSort(&input)
// 	}
// }
