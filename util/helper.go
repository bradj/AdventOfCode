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

func GetItems(filename string) []string {
	bcontent, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	content := string(bcontent)
	return strings.Split(content, "\n")
}

type Set interface {
	Add() bool
}

type RuneSet map[rune]bool

func (a RuneSet) Exists(item rune) bool {
	_, ok := a[item]
	return ok
}

func (a RuneSet) Add(item rune) bool {
	prevL := len(a)
	a[item] = true
	return len(a) != prevL
}
