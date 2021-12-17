package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func intme(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}

	return int(i)
}

type Board struct {
	Board   [][]int
	matches map[string]int
	Sum     int
	solved  bool
}

func (b *Board) SetBoard(board [][]int) {
	b.Board = board

	for _, row := range b.Board {
		for _, col := range row {
			b.Sum += col
		}
	}

	fmt.Printf("sum: %d\nboard: %v\n", b.Sum, b.Board)
}

func remove(s []Board, i int) []Board {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (b *Board) IsBingo(g int) bool {
	if b.solved {
		return true
	}

	if len(b.matches) == 0 {
		b.matches = make(map[string]int)
	}

	for row, rowv := range b.Board {
		for col, colv := range rowv {
			if colv != g {
				// no match
				continue
			}

			b.Sum -= g

			colidx := fmt.Sprintf("c%d", col)
			rowidx := fmt.Sprintf("r%d", row)

			if _, ok := b.matches[colidx]; ok {
				b.matches[colidx]++
			} else {
				b.matches[colidx] = 1
			}

			if _, ok := b.matches[rowidx]; ok {
				b.matches[rowidx]++
			} else {
				b.matches[rowidx] = 1
			}

			if b.matches[colidx] == 5 || b.matches[rowidx] == 5 {
				b.Sum *= g
				b.solved = true
				return true
			}
		}
	}

	return false
}

func main() {
	file, err := os.Open("4.input")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	draw := strings.Split(scanner.Text(), ",")
	row := 0
	var boards []Board
	var board [][]int

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}

		converted := []int{}

		for _, v := range strings.Fields(scanner.Text()) {
			converted = append(converted, intme(v))
		}

		board = append(board, converted)

		row++

		if row == 5 {
			b := Board{}
			b.SetBoard(board)
			boards = append(boards, b)
			board = nil
			row = 0
		}
	}

	fmt.Println(draw)

	total := len(boards)
	count := 0

	for _, v := range draw {
		for i := range boards {
			if boards[i].IsBingo(intme(v)) {
				count++

				if count == total {
					fmt.Printf("Final bingo!\n%v\nsum: %d\n", boards[i].Board, boards[i].Sum)
					return
				}
			}
		}
	}
}
