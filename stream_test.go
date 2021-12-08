package functional

import (
	"fmt"
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
	m := Map(Stream(data), func(i int) string { return strconv.Itoa(i + 1) })
	fmt.Println(m.Collect())
}

// func TestMapNew(t *testing.T) {
// 	data := []int{1, 2, 3, 4, 5}
// 	s := Stream(data).Map(func(i int) string { return strconv.Itoa(i + 1) })
// 	fmt.Println(s.Collect())
// }

func TestFilter(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	m := Filter(Stream(data), func(i int) bool { return i%2 == 0 })
	fmt.Println(m.Collect())
}

func TestReduce(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	m := ReduceFromZero(Filter(Stream(data), func(i int) bool { return i%2 == 0 }), func(i, j int) int { return i + j })
	fmt.Println(m)

	m2 := Reduce(Filter(Stream(data), func(i int) bool { return i%2 == 0 }), "0", func(i string, j int) string { return i + ", " + strconv.Itoa(j) })
	fmt.Println(m2)
}
