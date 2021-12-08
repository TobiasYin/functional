package functional

func ReduceFromZero[T any, O any, I Iterator[T]](iter I, reduceFunc func (O, T) O) O {
	return Reduce(iter, *new(O), reduceFunc)
}

func Reduce[T any, O any, I Iterator[T]](iter I, start O, reduceFunc func (O, T) O) O {
	r := &reduceIter[T, O, I]{
		iter: iter, 
		sum: start,
		reduceFunc: reduceFunc,
	}
	return r.Do()
}

type reduceIter[T any, O any, I Iterator[T]] struct {
	iter I
	sum O
	reduceFunc func (O, T) O
}

func (r *reduceIter[T, O, I])Do() O {
	for {
		data, err := r.iter.Next()
		if err != nil {
			break
		}
		r.sum = r.reduceFunc(r.sum, data)
	}
	return r.sum
}

