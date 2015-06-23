// [2015-06-22] Challenge #220 [Easy] Mangling sentences
// http://www.reddit.com/r/dailyprogrammer/comments/3aqvjn/20150622_challenge_220_easy_mangling_sentences/

package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(mangle("This challenge doesn't seem so hard."))
	fmt.Println(mangle("There are more things between heaven and earth, Horatio, than are dreamt of in your philosophy."))
	fmt.Println(mangle("Eye of Newt, and Toe of Frog, Wool of Bat, and Tongue of Dog."))
	fmt.Println(mangle("Adder's fork, and Blind-worm's sting, Lizard's leg, and Howlet's wing."))
	fmt.Println(mangle("For a charm of powerful trouble, like a hell-broth boil and bubble."))
}

func mangle(s string) string {
	split := strings.Split(s, " ")
	return SortString(split)
}

func SortString(s []string) string {

	var buffer bytes.Buffer

	for n := range s {

		r := []rune(s[n])

		sort.Sort(sortWord(r))

		buffer.WriteString(string(r))
		buffer.WriteString(" ")

	}

	return buffer.String()
}

type sortWord []rune

func (w sortWord) Len() int {
	return len(w)
}

func (w sortWord) Swap(i, j int) {

	if unicode.IsLetter(w[i]) && unicode.IsLetter(w[j]) {

		if unicode.IsUpper(w[i]) {
			w[i], w[j] = unicode.ToUpper(w[j]), unicode.ToLower(w[i])
		} else if unicode.IsUpper(w[j]) {
			w[i], w[j] = unicode.ToLower(w[j]), unicode.ToUpper(w[i])
		} else {
			w[i], w[j] = w[j], w[i]
		}
	}
}

func (w sortWord) Less(i, j int) bool {

	if !unicode.IsLetter(w[i]) || !unicode.IsLetter(w[j]) {
		return false
	}

	return unicode.ToLower(w[i]) < unicode.ToLower(w[j])
}
