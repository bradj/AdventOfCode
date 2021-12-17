package main

import "strconv"

func intme(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}

	return int(i)
}

func sum(ints []int) int {
	result := 0

	for _, v := range ints {
		result += v
	}

	return result
}
