// https://github.com/golang/go/wiki/SliceTricks

package list

type ArrayList[T any] struct {
	slice []T
}

func New[T any]() ArrayList[T] {
	return ArrayList[T]{
		slice: make([]T, 0),
	}
}

func (list *ArrayList[T]) Add(index int, element T) {
	switch {
	case 0 <= index && index < list.Length():
		list.slice = append(list.slice, *new(T))
		copy(list.slice[index+1:], list.slice[index:])
		list.slice[index] = element

	case index == list.Length():
		list.slice = append(list.slice, element)

	default:
		panic("index out of range")
	}
}

func (list *ArrayList[T]) Get(index int) T {
	switch {
	case 0 <= index && index < list.Length():
		return list.slice[index]
	default:
		panic("index out of range")
	}
}

func (list *ArrayList[T]) Length() int {
	return len(list.slice)
}

func (list *ArrayList[T]) Remove(index int) T {
	if index <= 0 || list.Length() <= index {
		panic("index out of range")
	}

	element := list.slice[index]

	copy(list.slice[index:], list.slice[index+1:])
	list.slice = list.slice[:len(list.slice)-1]

	return element
}

func (list *ArrayList[T]) Set(index int, element T) {
	if index < 0 || list.Length() <= index {
		panic("index out of range")
	}

	list.slice[index] = element
}
