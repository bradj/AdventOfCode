package twenttwo

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/bradj/AdventOfCode/util"
)

func D1p1() {
	bcontent, err := os.ReadFile("2022/1.input")

	if err != nil {
		log.Fatal(err)
	}

	content := string(bcontent)
	items := strings.Split(content, "\n")

	elf := 0
	most := 0

	for _, item := range items {
		if item == "" {
			elf = 0
			continue
		}

		elf += util.Intme(item)

		if elf > most {
			most = elf
		}
	}

	fmt.Printf("Most calories: %d\n", most)
}

func D1p2() {
	bcontent, err := os.ReadFile("2022/1.input")

	if err != nil {
		log.Fatal(err)
	}

	content := string(bcontent)
	items := strings.Split(content, "\n")

	elves := []int{}
	elf := 0
	elves = append(elves, 0)

	for _, item := range items {
		if item == "" {
			elf++
			elves = append(elves, 0)
			continue
		}

		elves[elf] += util.Intme(item)
	}

	sort.Ints(elves)

	l := len(elves)

	fmt.Printf("Top 3 calories: %d\n", util.Sum(elves[l-3:]))
}
