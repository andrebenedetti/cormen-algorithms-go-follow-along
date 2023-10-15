package chap2

func Merge(s1 []int, s2 []int) []int {
	result := make([]int, 0)

	i := 0
	j := 0
	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			result = append(result, s1[i])
			i++
		} else {
			result = append(result, s2[j])
			j++
		}
	}

	for i < len(s1) {
		result = append(result, s1[i])
		i++
	}

	for j < len(s2) {
		result = append(result, s2[j])
		j++
	}

	return result
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	r1 := MergeSort(arr[:len(arr)/2])
	r2 := MergeSort(arr[len(arr)/2:])
	return Merge(r1, r2)
}

func BetterMerge(arr *[]int, start int, middle int, end int) {
}
