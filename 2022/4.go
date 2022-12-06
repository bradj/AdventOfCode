package twenttwo

import (
	"fmt"
	"strings"

	"github.com/bradj/AdventOfCode/util"
)

func contains(start1 int, end1 int, start2 int, end2 int) bool {
	if start1 >= start2 && end1 <= end2 {
		return true
	}

	if start2 >= start1 && end2 <= end1 {
		return true
	}

	return false
}

func D4p1(items []string) {
	total := 0

	for _, line := range items {
		if line == "" {
			continue
		}

		pairs := strings.Split(line, ",")
		p1 := strings.Split(pairs[0], "-")
		p2 := strings.Split(pairs[1], "-")

		if contains(util.Intme(p1[0]), util.Intme(p1[1]), util.Intme(p2[0]), util.Intme(p2[1])) {
			total++
		}
	}

	fmt.Printf("%d\n", total)
}

func p2contains(start1 int, end1 int, start2 int, end2 int) bool {
	if start1 <= end2 && start1 >= start2 {
		return true
	}

	if end1 <= end2 && end1 >= start2 {
		return true
	}

	if start2 <= end1 && start2 >= start1 {
		return true
	}

	if end2 <= end1 && end2 >= start1 {
		return true
	}

	return false
}

func D4p2(items []string) {
	total := 0

	for _, line := range items {
		if line == "" {
			continue
		}

		pairs := strings.Split(line, ",")
		p1 := strings.Split(pairs[0], "-")
		p2 := strings.Split(pairs[1], "-")

		if p2contains(util.Intme(p1[0]), util.Intme(p1[1]), util.Intme(p2[0]), util.Intme(p2[1])) {
			total++
		}
	}

	fmt.Printf("%d\n", total)
}
