package functional

func Filter[T any, I Iterator[T]](iter I, filter func (T) bool) *FilterIter[T, I] {
	return &FilterIter[T, I]{
		iter: iter, 
		filter: filter,
	}
}

type FilterIter[T any, I Iterator[T]] struct {
	iter I
	filter func (T) bool
}

func (f *FilterIter[T, I])Next() (T, error) {
	for {
		data, err := f.iter.Next()
		if err != nil {
			return *new(T), err
		}
		ok := f.filter(data)
		if ok {
			return data, nil
		}
	}
}


func (f *FilterIter[T, I]) Collect() []T {
	return Collect[T](f)
}