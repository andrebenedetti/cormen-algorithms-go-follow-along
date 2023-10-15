package chap2

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	for _, testCase := range GetTestCases() {
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

// How InsertionSort performed in this case:
// 1777 ns/op    0 B/op        0 allocs/op
// How MergeSort performed
// 245823 ns/op  163921 B/op   3012 allocs/op
func BenchmarkMergeSortBestCase1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := buildSortedSlice(1000)
		b.StartTimer()
		MergeSort(input)
	}
}

// We should see gains in the worst case when comparing to InsertionSort
// In fact, the implementation of MergeSort being tested herealready performs 10x better for this input size.
// We'll also try to improve this MergeSort implementation after this.
// InsertionSort
// 36343372 ns/op               0 B/op          0 allocs/op
// MergeSort
// 2872028 ns/op         3282112 B/op      33851 allocs/op
func BenchmarkMergeSortWorstCase1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := buildWorstCaseInput(10000)
		b.StartTimer()
		MergeSort(input)
	}
}

// Let's verify that our enhanced merge sort works as expected
func TestBetterMergeSort(t *testing.T) {
	for _, testCase := range GetTestCases() {
		BetterMergeSort(&testCase.input, 0, len(testCase.input))
		if !reflect.DeepEqual(testCase.input, testCase.want) {
			t.Fatalf("want: %v, got: %v ", testCase.want, testCase.input)
		}
	}
}

// Let's see if it got better
// 2872028 ns/op         3282112 B/op      33851 allocs/op (FIRST VERSION)
// 1343224 ns/op         1110798 B/op      19998 allocs/op (THIS VERSION)
// Yay! More than 2x improvement by reducing the number of allocations!
func BenchmarkBetterMergeSortWorstCase1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := buildWorstCaseInput(10000)
		b.StartTimer()
		BetterMergeSort(&input, 0, len(input))
	}
}

// What if we grow the input size?
// Still, it is performing worse than the InsetionSort
func BenchmarkBetterMergeSortWorstCase2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := buildWorstCaseInput(40000)
		b.StartTimer()
		BetterMergeSort(&input, 0, len(input))
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
