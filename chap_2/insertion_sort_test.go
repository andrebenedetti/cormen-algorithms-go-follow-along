package chap2

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	for _, testCase := range GetTestCases() {
		InsertionSort(&testCase.input)
		if !reflect.DeepEqual(testCase.input, testCase.want) {
			t.Fatalf("want: %v, got: %v ", testCase.want, testCase.input)
		}
	}

}

// Chapter 2 shows Insertion Sort's best case scenario
// should have O(n) (a.k.a linear) time complexity)
// Lets use Go's built-in benchmarking tools to verify that.

// 1777 ns/op      0 B/op     0 allocs/op
func BenchmarkBestCase1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := buildSortedSlice(1000)
		b.StartTimer()
		InsertionSort(&input)
	}
}

func BenchmarkBestCase2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := buildSortedSlice(10000)
		b.StartTimer()
		InsertionSort(&input)
	}
}

func BenchmarkBestCase3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := buildSortedSlice(100000)
		b.StartTimer()
		InsertionSort(&input)
	}
}

func BenchmarkBestCase4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := buildSortedSlice(1000000)
		b.StartTimer()
		InsertionSort(&input)
	}
}

// We also learned that Insertion Sort's worst case
// scenario's time complexity is O(n^2). Let's check
// that

func BenchmarkInsertionSortWorstCase1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := buildWorstCaseInput(10000)
		b.StartTimer()
		InsertionSort(&input)
	}
}

func BenchmarkInsertionSortWorstCase2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := buildWorstCaseInput(20000)
		b.StartTimer()
		InsertionSort(&input)
	}
}

func BenchmarkInsertionSortWorstCase3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := buildWorstCaseInput(30000)
		b.StartTimer()
		InsertionSort(&input)
	}
}

func BenchmarkInsertionSortWorstCase4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := buildWorstCaseInput(40000)
		b.StartTimer()
		InsertionSort(&input)
	}
}
func BenchmarkInsertionSortWorstCase5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := buildWorstCaseInput(400000)
		b.StartTimer()
		InsertionSort(&input)
	}
}
