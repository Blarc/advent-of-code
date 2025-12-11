package utils

import "strconv"

func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func ToIntSlice(s []string) []int {
	r := make([]int, len(s))
	for i, v := range s {
		r[i], _ = strconv.Atoi(v)
	}
	return r
}
