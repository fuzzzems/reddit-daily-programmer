// [2015-07-06] Challenge #222 [Easy] Balancing Words
// http://www.reddit.com/r/dailyprogrammer/comments/3c9a9h/20150706_challenge_222_easy_balancing_words/

package main

import (
	"fmt"
)

func main() {
	s := "TO"
	fmt.Println(balance(s))
}

func balance(s string) string {

	length := len(s)
	pivot := 1

	for ; ; pivot++ {

		left := 0
		right := 0

		for i := 0; i < pivot; i++ {
			left += ((int(s[i]) - 64) * (pivot - i))
		}

		for j := pivot + 1; j < length; j++ {
			right += ((int(s[j]) - 64) * (j - pivot))
		}

		if left > right || pivot == length {
			break
		}

		if left == right {
			return fmt.Sprintf("%v %c %v - %d", s[:pivot], s[pivot], s[pivot+1:], left)
		}
	}

	return s + " does not balance"
}
