package functional


func Stream[T any](data []T) *StreamIter[T] {
	return &StreamIter[T]{data: data, now: 0}
}

type StreamIter[T any] struct{
	data []T
	now int
}

func(s *StreamIter[T]) Next() (T, error) {
	if s.now < len(s.data) {
		res := s.data[s.now]
		s.now += 1
		return res, nil
	}
	return *new(T), EndIterError
}

func (s *StreamIter[T]) Collect() []T {
	return Collect[T](s)
}


// Not Supported for now
// func (s *StreamIter[T]) Map[O any](mapFunc func(T) O) *MapIter[T, O, Iterator[T]]{
// 	return NewMapIter(s, mapFunc)
// }