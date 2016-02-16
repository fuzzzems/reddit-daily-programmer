// [2016-02-16] Challenge #254 [Easy] Atbash Cipher
// https://www.reddit.com/r/dailyprogrammer/comments/45w6ad/20160216_challenge_254_easy_atbash_cipher/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(atbash("foobar"))
	fmt.Println(atbash("wizard"))
	fmt.Println(atbash("/r/dailyprogrammer"))
	fmt.Println(atbash("gsrh rh zm vcznkov lu gsv zgyzhs xrksvi"))
	fmt.Println(atbash("Gsrh Rh Zm Vcznkov Lu Gsv Zgyzhs Xrksvi"))
}

func atbash(s string) string {
	return strings.Map(func(r rune) rune {
		switch {
		case r <= 'z' && r >= 'a':
			return 'z' - (r % 'a')
		case r <= 'Z' && r >= 'A':
			return 'Z' - (r % 'A')
		default:
			return r
		}
	}, s)
}
