// [2015-05-04] Challenge #213 [Easy] Pronouncing Hex
// http://www.reddit.com/r/dailyprogrammer/comments/34rxkc/20150504_challenge_213_easy_pronouncing_hex/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./values.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	tens := map[string]string{
		"A": "atta",
		"B": "bibbity",
		"C": "city",
		"D": "dickety",
		"E": "ebbity",
		"F": "fleventy",
		"0": "",
	}

	ones := map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
		"4": "four",
		"5": "five",
		"6": "six",
		"7": "seven",
		"8": "eight",
		"9": "nine",
		"A": "a",
		"B": "bee",
		"C": "cee",
		"D": "dee",
		"E": "e",
		"F": "eff",
		"0": "zero",
	}

	var output bytes.Buffer
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		output.WriteString(text)
		output.WriteString(" ")

		for i := 2; i < len(text); i = i + 2 {

			ten, one := tens[string(text[i])], ones[string(text[i+1])]

			output.WriteString(ten)
			output.WriteRune('-')

			if one != "zero" {
				output.WriteString(one)
				output.WriteString(" ")
			}

			if len(text) > 4 && i == 2 {
				output.WriteString("bitey ")
			}
		}
		output.WriteRune('\n')
	}
	fmt.Println(output.String())
}
