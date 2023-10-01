package chap2

/*
The books highlights that this algorithm
should sort *in place*, so we are using *[]int.
*/
func InsertionSort(s *[]int) {
	for i, val := range *s {
		if i == 0 {
			continue
		}

		j := i - 1
		for j >= 0 && (*s)[j] > val {
			(*s)[j+1] = (*s)[j]
			j--
		}
		(*s)[j+1] = val
	}
}
