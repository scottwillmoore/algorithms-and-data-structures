// https://pkg.go.dev/golang.org/x/exp/constraints

package constraint

type Float interface {
	~float32 | ~float64
}

type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Integer interface {
	SignedInteger | UnsignedInteger
}

type Number interface {
	Integer | Float
}

type String interface {
	~string
}

type Comparable interface {
	comparable
}

type Orderable interface {
	Number | String
}
