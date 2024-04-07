package slices

func Sum(num []int) int {
	sum := 0

	for _, n := range num {
		sum += n
	}

	return sum
}
