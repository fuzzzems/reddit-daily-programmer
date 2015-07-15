// [2015-07-15] Challenge #223 [Intermediate] Eel of Fortune
// http://www.reddit.com/r/dailyprogrammer/comments/3ddpms/20150715_challenge_223_intermediate_eel_of_fortune/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(problem("synchronized", "snond"))  // True
	fmt.Println(problem("misfunctioned", "snond")) // True
	fmt.Println(problem("mispronounced", "snond")) // False
	fmt.Println(problem("shotgunned", "snond"))    // False
	fmt.Println(problem("snond", "snond"))         // True
}

func problem(s, p string) bool {

	n := 0
	for i := range s {
		if strings.ContainsRune(p, rune(s[i])) {
			if s[i] != p[n] {
				return false
			}
			n++
		}
	}

	if n == len(p) {
		return true
	}

	return false
}
