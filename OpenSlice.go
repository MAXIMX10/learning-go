package main

func Len(s *Slice) int {
	return s.len
}

func Cap(s *Slice) int {
	return s.cap
}

func Append(s *Slice, newValues ...int) *Slice {
	newLen := s.len + len(newValues)
	oldCap := s.cap
	newCap := nextCap(newLen, oldCap)

	if newLen <= oldCap {
		for j := range len(newValues) - 1 {
			s.pointer[s.len+j] = newValues[j]
		}

		s.len = newLen

		return s
	}

	newSlice := Init(newLen, newCap)

	for i := range s.len {
		newSlice.pointer[i] = s.pointer[i]
	}
	for j := range len(newValues) {
		newSlice.pointer[s.len+j] = newValues[j]
	}

	return newSlice
}

func nextCap(newLen, oldCap int) int {
	newCap := oldCap
	doubleCap := newCap + newCap

	if newLen > doubleCap {
		return newLen
	}

	const threshold = 256

	if oldCap < threshold {
		return doubleCap
	}

	for {
		newCap += (newCap + 3*threshold) >> 2

		if uint(newCap) >= uint(newLen) {
			break
		}
	}

	if newCap <= 0 {
		return newLen
	}

	return newCap
}
