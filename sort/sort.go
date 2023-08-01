// https://github.com/golang/go/issues/46477
// https://github.com/golang/go/issues/57433
// https://github.com/golang/go/issues/60546
// https://pkg.go.dev/golang.org/x/exp/slices
// https://pkg.go.dev/sort

package sort

func IsSorted[T any](slice []T, less func(a, b T) bool) bool {
	for i := len(slice) - 1; i > 0; i-- {
		if less(slice[i], slice[i-1]) {
			return false
		}
	}
	return true
}
