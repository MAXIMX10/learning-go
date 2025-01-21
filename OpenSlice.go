package main

func Len(s Slice) int { // взял тип Slice вместо указателя на него, чтобы передавать сразу структуру, а не ее адрес
	return s.len
}

func Cap(s Slice) int {
	return s.cap
}
