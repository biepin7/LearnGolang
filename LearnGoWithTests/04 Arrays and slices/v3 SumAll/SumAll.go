package main

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums = make([]int, lengthOfNumbers)

	// 容易出错，因为 切片有容量大小
	// for i, numbers := range numbersToSum {
	// 	sums[i] = sum(numbers)
	// }

	// 使用 append 函数可以避免

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, sum(numbers))
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
