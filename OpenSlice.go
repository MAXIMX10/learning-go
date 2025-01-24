package main

func Len(s *Slice) int {
	return s.len
}

func Cap(s *Slice) int {
	return s.cap
}

func Append(s *Slice, newValues ...int) *Slice {
	old_len := s.len
	s.len = s.len + len(newValues)
	new_cap := s.cap

	if s.len > s.cap {
		if s.cap == 0 {
			new_cap = s.len
		} else {
			for new_cap < s.len {
				if new_cap < 256 {
					new_cap = new_cap * 2
				} else {
					new_cap = new_cap / 4
				}
			}
		}
	}

	new_Slice := Init(s.len, new_cap)
	copy(new_Slice.pointer, s.pointer)
	copy(new_Slice.pointer[old_len:], newValues)
	s = new_Slice

	return s
}
