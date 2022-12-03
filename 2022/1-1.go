package twenttwo

import (
	"fmt"
	"log"
	"os"
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
