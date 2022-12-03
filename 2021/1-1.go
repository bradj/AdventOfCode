package twentyone

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bradj/AdventOfCode/util"
)

func getWindowSum(items []int, idx int) int {
	length := len(items)

	if length < idx+3 {
		return 0
	}

	return util.Sum(items[idx : idx+3])
}

func D1() {
	bcontent, err := os.ReadFile("1.input")

	if err != nil {
		log.Fatal(err)
	}

	content := string(bcontent)
	sitems := strings.Split(content, "\n")

	items := []int{}

	for _, v := range sitems {
		cast, _ := strconv.Atoi(v)
		items = append(items, cast)
	}

	inc := 0

	for idx := range items {
		a := getWindowSum(items, idx)
		b := getWindowSum(items, idx+1)

		if b > a {
			inc++
		}
	}

	fmt.Printf("%d increases found\n", inc)
}
