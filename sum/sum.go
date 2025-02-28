package sum

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllTails(tailsToSum ...[]int) []int {
	numbersToSum := make([][]int, len(tailsToSum))
	for i, tail := range tailsToSum {
		if len(tail) == 0 {
			numbersToSum[i] = tail
		} else {
			numbersToSum[i] = tail[1:]
		}
	}
	return SumAll(numbersToSum...)
}
