package sort

func InsertionSort[T any](slice []T, less func(a, b T) bool) {
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0 && less(slice[j], slice[j-1]); j-- {
			slice[j-1], slice[j] = slice[j], slice[j-1]
		}
	}
}
