package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll takes a varying number of slices, returning a new slice containing the totals for each slice passed in.
// length and capacity - create array with length 2 (elements initialized by default to 0 for int
func SumAll(slicesToSum ...[]int) []int {
	var resultSlice []int
	for _, collection := range slicesToSum {
		resultSlice = append(resultSlice, Sum(collection))
	}
	return resultSlice
}

func AppendAllTails(numbersToSum ...[]int) []int {
	var resultSlice []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			resultSlice = append(resultSlice, 0)
		} else {
			tail := numbers[1:]
			resultSlice = append(resultSlice, Sum(tail))
		}
	}
	return resultSlice
}
