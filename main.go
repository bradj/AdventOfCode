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
		D1()

		break
	case 2:
		if part == 1 {
			D2p1()
		} else {
			D2p2()
		}

		break
	case 3:
		if part == 1 {
			D3p1()
		} else {
			D3p2()
		}

		break
	case 4:
		D4()

		break
	default:
		fmt.Printf("day %d part %d not found\n", day, part)
	}
}
