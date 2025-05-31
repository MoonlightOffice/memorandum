package main

func QuickSort[T any](x []T, less func(i, j int, x []T) bool) {
	if len(x) <= 1 {
		return
	}

	pivot := len(x) / 2
	left, right := []T{}, []T{}

	for i := range x {
		if i == pivot {
			continue
		}

		if less(i, pivot, x) {
			left = append(left, x[i])
		} else {
			right = append(right, x[i])
		}
	}

	QuickSort(left, less)
	QuickSort(right, less)

	copy(x, append(append(left, x[pivot]), right...))
}

// Gemini-generated

func QuickSort2[T any](x []T, less func(i int, j int) bool) []T {
	quickSort2Helper(x, 0, len(x)-1, less)
	return x
}

func quickSort2Helper[T any](x []T, low, high int, less func(i int, j int) bool) {
	if low < high {
		pi := partition2(x, low, high, less)
		quickSort2Helper(x, low, pi-1, less)
		quickSort2Helper(x, pi+1, high, less)
	}
}

func partition2[T any](x []T, low, high int, less func(i int, j int) bool) int {
	pivot := high
	i := low - 1

	for j := low; j < high; j++ {
		if less(j, pivot) {
			i++
			x[i], x[j] = x[j], x[i]
		}
	}

	x[i+1], x[high] = x[high], x[i+1]
	return i + 1
}
