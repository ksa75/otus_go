//package main
//
//import "fmt"
//
//func main() {
//	points := Points{1, 2, 3, 4}
//	points.Print()
//	newPoints := Scale(points, 2)
//	newPoints.Print()
//}
//
//type Point int
//
//type Points []Point
//
//func (p Points) Print() {
//	fmt.Println(p)
//}
//
//func Scale[S ~[]E, E Integer](s S, c E) S {
//	r := make([]E, len(s))
//	for i, v := range s {
//		r[i] = v * c
//	}
//	return r
//}
//
//type Unsigned interface {
//	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
//}
//type Signed interface {
//	~int | ~int8 | ~int16 | ~int32 | ~int64
//}
//type Integer interface {
//	Signed | Unsigned
//}
