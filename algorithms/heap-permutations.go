package algorithms

// This is Heap's algorithm to generate permutations
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func Generate(k int, a []int, results *[][]int) {
	if k == 1 {
		permutation := make([]int, len(a))
		copy(permutation, a)
		*results = append(*results, permutation)
		return
	}

	Generate(k-1, a, results)

	for i := 0; i < k-1; i += 1 {
		if k%2 == 0 {
			a[i], a[k-1] = a[k-1], a[i]
		} else {
			a[0], a[k-1] = a[k-1], a[0]
		}

		Generate(k-1, a, results)
	}
}
