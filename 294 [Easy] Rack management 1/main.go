// [2016-12-05] Challenge #294 [Easy] Rack management 1
// https://www.reddit.com/r/dailyprogrammer/comments/5go843/20161205_challenge_294_easy_rack_management_1/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(scrabble("ladilmy", "daily"))
	fmt.Println(scrabble("eerriin", "eerie"))
	fmt.Println(scrabble("orrpgma", "program"))
	fmt.Println(scrabble("orppgma", "program"))

	fmt.Println(scrabble("pizza??", "pizzazz"))
	fmt.Println(scrabble("piizza?", "pizzazz"))
	fmt.Println(scrabble("a??????", "program"))
	fmt.Println(scrabble("b??????", "program"))
}

func scrabble(tiles string, word string) bool {

	dict := make(map[rune]int)

	for _, r := range tiles {
		dict[r]++
	}

	for _, r := range word {
		dict[r]--
		if dict[r] < 0 {
			if dict['?'] > 0 {
				dict['?']--
			} else {
				return false
			}
		}
	}
	return true
}
