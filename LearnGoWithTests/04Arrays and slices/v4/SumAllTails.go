package main

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	// 这种写法，在传入 int[]{}时候，会导致 slice bounds out of range
	// for _, numbers := range numbersToSum {
	// 	tail := numbers[1:]
	// 	sums = append(sums, sum(tail))
	// }

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, sum(tail))
		}
	}

	return sums
}

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
