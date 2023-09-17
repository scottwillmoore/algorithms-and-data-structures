// https://docs.oracle.com/javase/8/docs/api/java/util/List.html
// https://en.wikipedia.org/wiki/List_(abstract_data_type)
// https://pkg.go.dev/container/list

package list

type List[T any] interface {
	Add(index int, element T)
	Get(index int) T
	Length() int
	Remove(index int) T
	Set(index int, element T)
}

type GoList[T any] struct {
	slice []T
}
