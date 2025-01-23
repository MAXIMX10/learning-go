package main

func Len(s *Slice) int {
	return s.len
}

func Cap(s *Slice) int {
	return s.cap
}

func Append(s *Slice, newValues ...int) *Slice {

	new_len := s.len + len(newValues)
	new_cap := s.cap

	if new_len > s.cap {
		new_cap = s.cap * 2
		if new_len > new_cap {
			new_cap = new_len
		}
	}

	new_Slice := Init(new_len, new_cap)
	copy(new_Slice.pointer, s.pointer)

	for i := range len(newValues) {
		new_Slice.Set(s.len+i, newValues[i])
	}

	return new_Slice

}
