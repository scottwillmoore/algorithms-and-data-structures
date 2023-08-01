package sort

func SelectionSort[T any](slice []T, less func(a, b T) bool) {
	for i := 0; i < len(slice); i++ {
		jMin := i
		for j := i + 1; j < len(slice); j++ {
			if less(slice[j], slice[jMin]) {
				jMin = j
			}
		}
		if i != jMin {
			slice[i], slice[jMin] = slice[jMin], slice[i]
		}
	}
}
