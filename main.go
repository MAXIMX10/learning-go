package main

type Slice struct {
	// 1 шаг
	pointer *[]int // считаем за массив
	len     int
	cap     int
	//Проверка новой ветки
}

func (s *Slice) Init(length, capacity int) *Slice {
	panic("Implement me")
}

type ISlice interface {
	Init(length, capacity int) *Slice // make([]int, 4, 8)	2 шаг
	Get(idx int) int                  // s[1] 				2 шаг
	Set(idx, value int)               // s[1] = 1 			2 шаг
	Cut(i, j int) *Slice              // s[1:4]				3 шаг
}

type OpenSlice interface {
	Append(slice *Slice, newValues ...int) *Slice // append(a, 1, 2, 3), 	3 шаг
	Len()                                         // 																2 шаг
	Cap()                                         // 																2 шаг
}

func Append(slice *Slice, newValues ...int) *Slice {
	panic("implement me")
}

func Len() {}

func Cap() {}

func main() {
	a := Init(4, 8)
	a.Cut(1, 4) // [0 0 0]

	Append(a, 1, 2, 3)
	Len(a)
	Cap(a)
}
