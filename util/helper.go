package util

import (
	"strconv"
	"strings"
)

func Intme(s string) int {
	i, err := strconv.ParseInt(strings.Trim(s, " "), 10, 64)

	if err != nil {
		panic(err)
	}

	return int(i)
}

func Sum(ints []int) int {
	result := 0

	for _, v := range ints {
		result += v
	}

	return result
}
