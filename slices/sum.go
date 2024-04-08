package slices

func Sum(num []int) int {
	sum := 0

	for _, n := range num {
		sum += n
	}

	return sum
}

func SumAll(nums ...[]int) []int {
	var sums []int

	for _, val := range nums {
		sums = append(sums, Sum(val))
	}

	return sums
}

func SumAllTails(nums ...[]int) []int {
	var sums []int

	for _, val := range nums {
		tails := val[1:]

		if len(val) == 1 {
			tails = val[0:]
		}

		sums = append(sums, Sum(tails))
	}

	return sums
}
