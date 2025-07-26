package arrays

func Sum(numbers []int) int {
	var sum int
	// Defines if it will count the first value on the collection
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func SumAll(sumAll ...[]int) []int {
	var sums []int
	for _, slice := range sumAll {
		sums = append(sums, Sum(slice))
	}
	return sums
}

func SumAllTails(sumAllTails ...[]int) []int {
	var tailSums []int

	for _, slice := range sumAllTails {
		if len(slice) == 0 {
			tailSums = append(tailSums, 0)
		} else {
			// Ignore the first value
			tailSlice := slice[1:]
			tailSums = append(tailSums, Sum(tailSlice))
		}
	}
	return tailSums
}
