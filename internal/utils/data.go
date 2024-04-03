package utils

import "strconv"

func AtoiButGood(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
