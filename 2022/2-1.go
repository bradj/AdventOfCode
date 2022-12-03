package twenttwo

import (
	"fmt"
	"strings"

	"github.com/bradj/AdventOfCode/util"
)

const (
	OpRock     = "A"
	Rock       = "X"
	OpPaper    = "B"
	Paper      = "Y"
	OpScissors = "C"
	Scissors   = "Z"
	Win        = 6
	Loss       = 0
	Draw       = 3
)

var choices map[string]map[string]int
var choiceScore map[string]int

func init() {
	choices = make(map[string]map[string]int)

	choices[OpRock] = make(map[string]int)
	choices[OpRock][Rock] = Draw
	choices[OpRock][Paper] = Win
	choices[OpRock][Scissors] = Loss

	choices[OpPaper] = make(map[string]int)
	choices[OpPaper][Rock] = Loss
	choices[OpPaper][Paper] = Draw
	choices[OpPaper][Scissors] = Win

	choices[OpScissors] = make(map[string]int)
	choices[OpScissors][Rock] = Win
	choices[OpScissors][Paper] = Loss
	choices[OpScissors][Scissors] = Draw

	choiceScore = make(map[string]int)

	choiceScore[Rock] = 1
	choiceScore[Paper] = 2
	choiceScore[Scissors] = 3
}

func D2p1() {
	items := util.GetItems("2022/2.input")

	total := 0

	for _, item := range items {
		if item == "" {
			continue
		}

		round := strings.Split(item, " ")
		result := choices[round[0]][round[1]] + choiceScore[round[1]]
		total += result
	}

	fmt.Printf("total: %d\n", total)
}
