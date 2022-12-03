package main

import (
	"fmt"
	"os"
)

func main() {
	day := intme(os.Args[1])
	part := intme(os.Args[2])

	fmt.Printf("running day %d part %d\n", day, part)

	switch day {
	case 1:
		d1()

		break
	case 2:
		if part == 1 {
			d2p1()
		} else {
			d2p2()
		}

		break
	case 3:
		if part == 1 {
			d3p1()
		} else {
			d3p2()
		}

		break
	case 4:
		d4()

		break
	default:
		fmt.Printf("day %d part %d not found\n", day, part)
	}
}
