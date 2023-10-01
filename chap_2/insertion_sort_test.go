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
	}

	for _, testCase := range cases {
		InsertionSort(&testCase.input)
		if !reflect.DeepEqual(testCase.input, testCase.want) {
			t.Fatalf("want: %v, got: %v ", testCase.want, testCase.input)
		}
	}

}
