// [2016-03-23] Challenge #259 [Intermediate] Mahjong Hands
// https://www.reddit.com/r/dailyprogrammer/comments/4bmdwz/20160323_challenge_259_intermediate_mahjong_hands/

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	circle, character, bamboo := make([]int, 9), make([]int, 9), make([]int, 9)

	file, err := os.Open("sample-input.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2

	// Get rid of that first line
	// Ignore the output and error.
	reader.Read()

	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		val, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}

		val--

		switch line[0] {
		case "Circle":
			circle[val] += 1
		case "Bamboo":
			bamboo[val] += 1
		case "Character":
			character[val] += 1
		}
	}

	if isWinningHand(circle, bamboo, character) {
		fmt.Println("Winning hand")
	} else {
		fmt.Println("Not a winning hand")
	}
}

func isWinningHand(inputs ...[]int) bool {

	for _, v := range inputs {

		v = findMatches(v)

		for _, m := range v {
			if m != 0 {
				return false
			}
		}
	}

	return true
}

func findMatches(input []int) []int {
	input = findSequence(input)
	input = findMulti(input)
	return input
}

func findMulti(input []int) []int {
	for k, v := range input {
		if v == 2 || v == 3 {
			input[k] = 0
		}
	}

	return input
}

func findSequence(input []int) []int {

	n := 0
	for k, v := range input {
		if v != 0 {
			n++
		} else {
			n = 0
		}

		if n == 3 {
			input[k]--
			input[k-1]--
			input[k-2]--

			n = 0
		}

	}

	return input

}
