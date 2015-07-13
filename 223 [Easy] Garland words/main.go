// [2015-07-13] Challenge #223 [Easy] Garland words
// http://www.reddit.com/r/dailyprogrammer/comments/3d4fwj/20150713_challenge_223_easy_garland_words/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(garland("programmer"))
	fmt.Println(garland("ceramic"))
	fmt.Println(garland("onion"))
	fmt.Println(garland("alfalfa"))

}

func garland(s string) int {
	n := 0
	l := len(s)
	for i := range s {
		if s[:i] == s[l-i:] {
			n = i
		}
	}
	return n
}
