// [2016-02-16] Challenge #254 [Easy] Atbash Cipher
// https://www.reddit.com/r/dailyprogrammer/comments/45w6ad/20160216_challenge_254_easy_atbash_cipher/

package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(atbash("foobar"))
	fmt.Println(atbash("wizard"))
	fmt.Println(atbash("/r/dailyprogrammer"))
	fmt.Println(atbash("gsrh rh zm vcznkov lu gsv zgyzhs xrksvi"))
	fmt.Println(atbash("Gsrh Rh Zm Vcznkov Lu Gsv Zgyzhs Xrksvi"))
}

func atbash(s string) string {
	var o bytes.Buffer
	for _, c := range s {

		var t rune
		switch {
		case c <= 122 && c >= 97:
			t = 122 - (c % 97)
		case c <= 90 && c >= 65:
			t = 90 - (c % 65)
		default:
			t = c
		}
		_, err := o.WriteRune(t)
		if err != nil {
			panic(err)
		}
	}
	return o.String()
}
