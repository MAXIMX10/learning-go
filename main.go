package main

import (
	"fmt"
)

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
	fmt.Printf("a: \n")
	//a := []int{0, 1, 2, 3}
	a := Init(4, 4)
	a.Set(0, 0)
	a.Set(1, 1)
	a.Set(2, 2)
	a.Set(3, 3)
	fmt.Println(a)
	a = Append(a, 1)
	fmt.Println(a)

	fmt.Printf("b: \n")
	b := Init(4, 4)
	b.Set(0, 0)
	b.Set(1, 1)
	b.Set(2, 2)
	b.Set(3, 3)
	//b := []int{0, 1, 2, 3}
	fmt.Println(b)
	b = Append(b, 1, 1)
	fmt.Println(b)

	fmt.Printf("c: \n")
	c := Init(4, 4)
	c.Set(0, 0)
	c.Set(1, 1)
	c.Set(2, 2)
	c.Set(3, 3)
	//c := []int{0, 1, 2, 3}
	fmt.Println(c)
	c = Append(c, 1, 1, 1, 1, 1, 1, 1)
	fmt.Println(c)
}
