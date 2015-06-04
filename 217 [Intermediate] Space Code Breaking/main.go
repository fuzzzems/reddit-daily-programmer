// [2015-06-03] Challenge #217 [Intermediate] Space Code Breaking
// http://www.reddit.com/r/dailyprogrammer/comments/38fjll/20150603_challenge_217_intermediate_space_code/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := strings.TrimPrefix(scanner.Text(), "\"")
		text = strings.TrimSuffix(text, "\"")
		fmt.Println(decodeMessage(text))
	}

}

func decodeMessage(s string) string {

	htrae, a := decodeHtrae(s)
	hoth, b := decodeHoth(s)
	ryza, c := decodeRyza(s)
	omicon, d := decodeOmicron(s)

	message, large := htrae, a

	if large < b {
		message, large = hoth, b
	}
	if large < c {
		message, large = ryza, c
	}
	if large < d {
		message = omicon
	}
	return message
}

func decodeHtrae(message string) (string, int) {

	decoded := []rune(message)
	weight := 0

	for i, n := 0, utf8.RuneCountInString(message)-1; i < n; i, n = i+1, n-1 {
		decoded[i], decoded[n] = decoded[n], decoded[i]
		weight += determineWeight(decoded[i])

	}
	return "Htrea: " + string(decoded), weight
}

func decodeHoth(message string) (string, int) {

	var buffer bytes.Buffer
	weight := 0

	subTen := func(r rune) rune {
		weight += determineWeight(r - 10)
		return r - 10
	}
	buffer.WriteString(strings.Map(subTen, message))

	return "Hoth: " + buffer.String(), weight
}

func decodeRyza(message string) (string, int) {

	var buffer bytes.Buffer
	weight := 0

	addOneUnlessSpace := func(r rune) rune {
		if unicode.IsSpace(r) {
			weight += -2
			return r
		}
		weight += determineWeight(r + 1)
		return r + 1
	}
	buffer.WriteString(strings.Map(addOneUnlessSpace, message))

	return "Ryza: " + buffer.String(), weight
}

func decodeOmicron(message string) (string, int) {

	var buffer bytes.Buffer
	weight := 0

	invertFifthBit := func(r rune) rune {
		weight += determineWeight(r ^ 16)
		return r ^ 16
	}
	buffer.WriteString(strings.Map(invertFifthBit, message))

	return "Omicron: " + buffer.String(), weight

}

func determineWeight(r rune) int {

	weight := 0
	switch {
	case unicode.IsLetter(r):
		weight += 1
		break
	case unicode.IsPunct(r):
		weight += -2
		break
	case unicode.IsSpace(r):
		weight += 2
		break
		weight += -20
	}

	return weight
}
