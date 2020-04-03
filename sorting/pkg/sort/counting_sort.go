package sort

// CountingSort ...
func CountingSort(a []int) {
	maxValue := getMaxValue(a)

	count := make([]int, maxValue+1)

	for i := 0; i < len(a); i++ {
		count[a[i]]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	b := make([]int, len(a))

	for i := len(a) - 1; i >= 0; i-- {
		count[a[i]]--
		b[count[a[i]]] = a[i]
	}

	for i := 0; i < len(a); i++ {
		a[i] = b[i]
	}
}

func getMaxValue(a []int) int {
	maxValue := 0

	for i := 0; i < len(a); i++ {
		if a[i] > maxValue {
			maxValue = a[i]
		}
	}

	return maxValue
}
