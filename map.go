package functional

func Map[T any, O any, I Iterator[T]](iter I, mapFunc func (T) O) *MapIter[T, O, I] {
	return &MapIter[T, O, I]{
		iter: iter, 
		mapFunc: mapFunc,
	}
}

type MapIter[T any, O any, I Iterator[T]] struct {
	iter I
	mapFunc func (T) O
}

func (m *MapIter[T, O, I])Next() (O, error) {
	data, err := m.iter.Next()
	if err != nil {
		return *new(O), err
	}
	return m.mapFunc(data), nil
}


func (m *MapIter[T, O, I]) Collect() []O {
	return Collect[O](m)
}