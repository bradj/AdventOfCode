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
	case 3:
		items := util.GetItems("2022/3.input")

		if part == 1 {
			twentytwo.D3p1(items)
		} else {
			twentytwo.D3p2(items)
		}
	case 4:
		items := util.GetItems("2022/4.input")

		if part == 1 {
			twentytwo.D4p1(items)
		} else {
			twentytwo.D4p2(items)
		}
	case 5:
		items := util.GetItems("2022/5.input")
		// items := util.GetItems("2022/5-test.input")

		if part == 1 {
			twentytwo.D5p1(items)
		} else {
			twentytwo.D5p2(items)
		}
	case 6:
		items := util.GetItems("2022/6.input")

		if part == 1 {
			twentytwo.D6p1(items)
		} else {
			twentytwo.D6p2(items)
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
