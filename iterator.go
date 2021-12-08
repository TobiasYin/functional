package functional

import "errors"

var EndIterError = errors.New("end iterator error")

type any interface {}

type Iterator[T any] interface {
	Next() (T, error)
	Collect() []T 
}

func Collect[T any, I Iterator[T]](iter I) []T {
	res := make([]T, 0, 4)
	for {
		data, err := iter.Next()
		if err != nil {
			break
		}
		res = append(res, data)
	}
	return res
}
