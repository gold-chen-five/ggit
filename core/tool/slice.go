package tool

func SliceInsert[T any](s []T, index int, value T) []T {
	if len(s) == index {
		return append(s, value)
	}
	s = append(s[:index+1], s[index:]...)
	s[index] = value
	return s
}
