package arrays

func ArraySum(numbers []int) int {
	var sum int
	for i := range len(numbers) {
		sum += numbers[i]
	}
	return sum
}
