package main

import (
	"fmt"
)

type ISlice interface {
	Init(length, capacity int) *Slice // make([]int, 4, 8)															2 шаг
	Get(idx int) int                  // s[1] 																		2 шаг
	Set(idx, value int)               // s[1] = 1 																	2 шаг
	Cut(i, j int) *Slice              // s[1:4]																		3 шаг
}

type OpenSlice interface {
	Append(slice *Slice, newValues ...int) *Slice // append(a, 1, 2, 3), 											3 шаг
	Len()                                         // 																2 шаг
	Cap()                                         // 																2 шаг
}

func main() {
	/*a := Init(4, 8)
	a.Cut(1, 4) // [0 0 0]

	Append(a, 1, 2, 3)
	Len(a)
	Cap(a)*/
	var s Slice
	fmt.Println(s.Init(4, 5))
	fmt.Println(s.Get(3))
	s.Set(3, 10)
	fmt.Println(s.Get(3))
	fmt.Println(s.pointer)
	fmt.Println(Len(s)) // вместо &s
	fmt.Println(Cap(s)) // вместо &s
}
