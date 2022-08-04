package common

type Iterable[T any] interface {
	HasNext() bool
	Next() T
}
