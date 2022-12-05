package util

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Intme(s string) int {
	i, err := strconv.ParseInt(strings.Trim(s, " "), 10, 64)
	if err != nil {
		log.Fatal(err)
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

func GetItems(filename string) []string {
	bcontent, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	content := string(bcontent)
	content = strings.Trim(content, " ")
	return strings.Split(content, "\n")
}

type Set interface {
	Add() bool
	Contains() bool
}

type RuneSet map[rune]struct{}

func (a RuneSet) Contains(item rune) bool {
	_, ok := a[item]
	return ok
}

func (a RuneSet) Add(item rune) bool {
	prevL := len(a)
	a[item] = struct{}{}
	return len(a) != prevL
}
