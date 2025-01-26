package main

type ISlice interface {
	Get(idx int) int                  // s[1] 																		2 шаг
	Set(idx, value int)               // s[1] = 1 																	2 шаг
	Init(length, capacity int) *Slice // make([]int, 4, 8)															2 шаг
	Cut(i, j int) *Slice              // s[1:4]																		3 шаг
}

type OpenSlice interface {
	Len(s *Slice) int                             // 																2 шаг
	Cap(s *Slice) int                             // 																2 шаг
	Append(slice *Slice, newValues ...int) *Slice // append(a, 1, 2, 3), 											3 шаг
}

func main() {

}
