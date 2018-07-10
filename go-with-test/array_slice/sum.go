package array_slice

func Sum(numbers []int) int {
	sum := 0
	// for _, := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	for _, num := range numbers {
		sum += num
	}

	return sum
}

func SumFixed(numbers [5]int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
