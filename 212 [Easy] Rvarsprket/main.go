// [2015-04-27] Challenge #212 [Easy] Rövarspråket
// http://www.reddit.com/r/dailyprogrammer/comments/341c03/20150427_challenge_212_easy_rövarspråket/

package main

import (
	"bytes"
	"unicode"
)

var consonants = []rune{'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z'}

func Robberify(s string) string {

	var buffer bytes.Buffer

	for _, c := range s {

		if IsConsonant(unicode.ToLower(c)) {
			buffer.WriteRune(c)
			buffer.WriteRune('o')
			buffer.WriteRune(unicode.ToLower(c))
		} else {
			buffer.WriteRune(c)
		}
	}

	return buffer.String()
}

func IsConsonant(c rune) bool {
	for _, v := range consonants {
		if c == v {
			return true
		}
	}
	return false
}
