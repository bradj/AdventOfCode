package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func cull(items []string, focusIdx int, getMost bool) []string {
	ones := []string{}
	zeros := []string{}

	for _, v := range items {
		if len(v) == 0 {
			break
		}

		// 48 ord = 0 ascii
		// 49 ord = 1 ascii
		if v[focusIdx] == 48 {
			zeros = append(zeros, v)
		} else {
			ones = append(ones, v)
		}
	}

	if getMost {
		if len(ones) >= len(zeros) {
			return ones
		}

		return zeros
	} else {
		if len(ones) >= len(zeros) {
			return zeros
		}

		return ones
	}
}

func main() {
	// bcontent, err := ioutil.ReadFile("3.input")
	bcontent, err := ioutil.ReadFile("3.input")

	if err != nil {
		log.Fatal(err)
	}

	content := string(bcontent)
	items := strings.Split(content, "\n")

	o2items := cull(items, 0, true)
	idx := 1

	for len(o2items) > 1 {
		o2items = cull(o2items, idx, true)
		idx++
	}

	co2items := cull(items, 0, false)
	idx = 1

	for len(co2items) > 1 {
		co2items = cull(co2items, idx, false)
		idx++
	}

	fmt.Printf("o2 - %v\nco2 - %v\n", o2items, co2items)
	o2, _ := strconv.ParseInt(o2items[0], 2, 32)
	co2, _ := strconv.ParseInt(co2items[0], 2, 32)

	fmt.Printf("%d\n", co2*o2)
}
