# Funtional Api for Golang
Functional Programming support for golang.(Streaming API)

The package can only be used with go 1.18. Do not try in lower version of go.

## Examples

### Map
```go
data := []int{1, 2, 3, 4, 5}
res := Map(Stream(data), func(i int) string { return strconv.Itoa(i + 1) }).Collect()

fmt.Println(res)
fmt.Println(reflect.TypeOf(res))

//output:
//[2 3 4 5 6]
//[]string
```

### Filter

```go
data := []int{1, 2, 3, 4, 5, 6, 7, 8}
res := Filter(Stream(data), func(i int) bool { return i%2 == 0 }).Collect()
fmt.Println(res)

//output:
//[2 4 6 8]
```

#### combine use map and filter
```go
data := []int{1, 2, 3, 4, 5, 6, 7, 8}
res := Map(Filter(Stream(data), func(i int) bool { return i%2 == 0 }), func(i int) string { return strconv.Itoa(i + 1) }).Collect()
fmt.Println(res)
fmt.Println(reflect.TypeOf(res))

//output:
//[3 5 7 9]
//[]string
```

### Reduce
```go
data := []int{1, 2, 3, 4, 5, 6, 7, 8}
res1 := ReduceFromZero(Filter(Stream(data), func(i int) bool { return i%2 == 0 }), func(i, j int) int { return i + j })
fmt.Println(res1)

res2 := Reduce(Filter(Stream(data), func(i int) bool { return i%2 == 0 }), "0", func(i string, j int) string { return i + ", " + strconv.Itoa(j) })
fmt.Println(res2)
//output:
//20
//0, 2, 4, 6, 8
```