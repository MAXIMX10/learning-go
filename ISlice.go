package main

type Slice struct {
	// 1 шаг
	pointer []int
	len     int
	cap     int
}

func (s *Slice) Get(idx int) int {

	if (idx < 0) || (idx >= s.len) {
		panic("Input value of index isn't correct")
	}

	return s.pointer[idx]

}

func (s *Slice) Set(idx, value int) {
	if (idx < 0) || (idx >= s.len) {
		panic("Input value of index is not correct")
	}
	s.pointer[idx] = value
}

func Init(length, capacity int) *Slice {

	var s Slice

	if length < 0 {
		panic("Length can't be less than 0")
	}
	if capacity < 0 {
		panic("Capacity can't be less than 0")
	}
	if length > capacity {
		panic("Capacity can't be less than length")
	}

	arr := make([]int, length, capacity)

	s.pointer = arr
	s.len = length
	s.cap = capacity

	return &s
}
