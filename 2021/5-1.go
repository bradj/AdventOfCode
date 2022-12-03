package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createLine(x1 int, y1 int, x2 int, y2 int) (out [][]int) {
	if x1 == x2 {

	} else if y1 == y2 {

	}

	return out
}

func d5p1() {
	seabed := make(map[int]map[int]int)

	file, err := os.Open("5.test")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "->")

		coord1 := strings.Split(line[0], ",")
		coord2 := strings.Split(line[1], ",")

		x1 := intme(coord1[0])
		y1 := intme(coord1[1])
		x2 := intme(coord2[0])
		y2 := intme(coord2[1])

		if x1 == x2 {
			if _, ok := seabed[x]; !ok {
				seabed[x] = map[int]int{y: 1}
				continue
			}

			if _, ok := seabed[x][y]; !ok {
				seabed[x][y] = 1
				continue
			}

			seabed[x][y]++
		}

		if y1 == y2 {
			continue
		}
	}

	fmt.Printf("%v\n", seabed)
}
