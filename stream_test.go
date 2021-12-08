package functional

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestStream(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	s := Stream(data)
	r, err := s.Next()
	fmt.Println(r, err)
	fmt.Println(s.Collect())
}

func TestMap(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	res := Map(Stream(data), func(i int) string { return strconv.Itoa(i + 1) }).Collect()

	fmt.Println(res)
	fmt.Println(reflect.TypeOf(res))
}

// func TestMapNew(t *testing.T) {
// 	data := []int{1, 2, 3, 4, 5}
// 	s := Stream(data).Map(func(i int) string { return strconv.Itoa(i + 1) })
// 	fmt.Println(s.Collect())
// }

func TestFilter(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := Filter(Stream(data), func(i int) bool { return i%2 == 0 }).Collect()
	fmt.Println(res)
}

func TestFilterAndMap(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := Map(Filter(Stream(data), func(i int) bool { return i%2 == 0 }), func(i int) string { return strconv.Itoa(i + 1) }).Collect()
	fmt.Println(res)
	fmt.Println(reflect.TypeOf(res))
}

func TestReduce(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res1 := ReduceFromZero(Filter(Stream(data), func(i int) bool { return i%2 == 0 }), func(i, j int) int { return i + j })
	fmt.Println(res1)

	res2 := Reduce(Filter(Stream(data), func(i int) bool { return i%2 == 0 }), "0", func(i string, j int) string { return i + ", " + strconv.Itoa(j) })
	fmt.Println(res2)
}
