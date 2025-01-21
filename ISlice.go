package main

type Slice struct {
	// 1 шаг
	pointer []int
	len     int
	cap     int
}

func (s *Slice) Init(length, capacity int) *Slice {

	if (length > 0) && (capacity > 0) && (length <= capacity) {
		arr := make([]int, 10, 10)                 // считаем за базовый массив
		arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // для проверок
		//arr_pointer := &arr
		s.pointer = arr
		s.len = length
		s.cap = capacity

		for i := range length {
			s.pointer[i] = 0
		}

	} else {
		panic("Input length or capacity is incorrect")
	}
	return s // Но при выводе результата make выведется только слайс (без длины и ёмкости)
}

func (s *Slice) Get(idx int) int {
	if (idx >= 0) && (idx < s.len) {
		return s.pointer[idx]
	} else {
		panic("Input value of index is not correct")
	}

}

func (s *Slice) Set(idx, value int) {
	if (idx >= 0) && (idx < s.len) {
		s.pointer[idx] = value
	} else {
		panic("Input value of index is not correct")
	}
}
