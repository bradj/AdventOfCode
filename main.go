package main

import (
	"fmt"
	"os"

	twentyone "github.com/bradj/AdventOfCode/2021"
	twentytwo "github.com/bradj/AdventOfCode/2022"
	"github.com/bradj/AdventOfCode/util"
)

func year21(day int, part int) {
	switch day {
	case 1:
		twentyone.D1()
	case 2:
		if part == 1 {
			twentyone.D2p1()
		} else {
			twentyone.D2p2()
		}
	case 3:
		if part == 1 {
			twentyone.D3p1()
		} else {
			twentyone.D3p2()
		}
	case 4:
		twentyone.D4()
	default:
		fmt.Printf("day %d part %d not found\n", day, part)
	}
}

func year22(day int, part int) {
	switch day {
	case 1:
		if part == 1 {
			twentytwo.D1p1()
		} else {
			twentytwo.D1p2()
		}
	case 2:
		if part == 1 {
			twentytwo.D2p1()
		} else {
			twentytwo.D2p2()
		}
	default:
		fmt.Printf("day %d part %d not found\n", day, part)
	}
}

func main() {
	year := util.Intme(os.Args[1])
	day := util.Intme(os.Args[2])
	part := util.Intme(os.Args[3])

	fmt.Printf("running year %d day %d part %d\n", year, day, part)

	switch year {
	case 2021:
		year21(day, part)
	case 2022:
		year22(day, part)
	default:
		fmt.Printf("year %d not found\n", year)
	}
}
