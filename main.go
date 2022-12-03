package main

import (
	"fmt"
	"os"

	twentyone "github.com/bradj/AdventOfCode/2021"
	"github.com/bradj/AdventOfCode/util"
)

func main() {
	day := util.Intme(os.Args[1])
	part := util.Intme(os.Args[2])

	fmt.Printf("running day %d part %d\n", day, part)

	switch day {
	case 1:
		twentyone.D1()

		break
	case 2:
		if part == 1 {
			twentyone.D2p1()
		} else {
			twentyone.D2p2()
		}

		break
	case 3:
		if part == 1 {
			twentyone.D3p1()
		} else {
			twentyone.D3p2()
		}

		break
	case 4:
		twentyone.D4()

		break
	default:
		fmt.Printf("day %d part %d not found\n", day, part)
	}
}
