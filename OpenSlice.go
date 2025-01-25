package main

func Len(s *Slice) int {
	return s.len
}

func Cap(s *Slice) int {
	return s.cap
}

func Append(s *Slice, newValues ...int) *Slice {
	old_len := s.len
	new_cap := s.cap

	if s.len+len(newValues) <= s.cap {
		s.len = s.len + len(newValues)
		copy(s.pointer[s.len:], newValues)
	} else {
		if new_cap < s.len+len(newValues) {
			new_cap = new_cap * 2
		}

		for new_cap < s.len+len(newValues) {
			new_cap += new_cap / 4
		}

		new_Slice := Init(s.len+len(newValues), new_cap)
		copy(new_Slice.pointer, s.pointer)
		copy(new_Slice.pointer[old_len:], newValues)
		//s = new_Slice
		return new_Slice
	}

	return s
}
