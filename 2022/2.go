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

var outcomes map[string]map[string]string

func D2p2() {
	items := util.GetItems("2022/2.input")

	outcomes = make(map[string]map[string]string)

	outcomes[OpRock] = make(map[string]string)
	outcomes[OpRock]["X"] = Scissors // lose
	outcomes[OpRock]["Y"] = Rock     // draw
	outcomes[OpRock]["Z"] = Paper    // win

	outcomes[OpPaper] = make(map[string]string)
	outcomes[OpPaper]["X"] = Rock
	outcomes[OpPaper]["Y"] = Paper
	outcomes[OpPaper]["Z"] = Scissors

	outcomes[OpScissors] = make(map[string]string)
	outcomes[OpScissors]["X"] = Paper
	outcomes[OpScissors]["Y"] = Scissors
	outcomes[OpScissors]["Z"] = Rock

	total := 0

	for _, item := range items {
		if item == "" {
			continue
		}

		round := strings.Split(item, " ")
		myChoice := outcomes[round[0]][round[1]]
		result := choices[round[0]][myChoice] + choiceScore[myChoice]
		total += result
	}

	fmt.Printf("total: %d\n", total)
}
