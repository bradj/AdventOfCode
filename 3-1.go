package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("3.input")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var bitcounter [12][2]int

	for scanner.Scan() {
		for idx, v := range scanner.Text() {
			bitcounter[idx][int(v-'0')]++
		}
	}

	// gamma is most common bit
	// epsilon is least common bit

	gamma := ""
	epsilon := ""

	for _, v := range bitcounter {
		if v[0] > v[1] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	ep, _ := strconv.ParseInt(epsilon, 2, 64)
	ga, _ := strconv.ParseInt(gamma, 2, 64)

	fmt.Printf("%d\n", ep*ga)
}
