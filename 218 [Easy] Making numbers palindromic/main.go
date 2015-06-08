// [2015-06-08] Challenge #218 [Easy] Making numbers palindromic
// http://www.reddit.com/r/dailyprogrammer/comments/38yy9s/20150608_challenge_218_easy_making_numbers/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(convertToPalindrome(123))
	fmt.Println(convertToPalindrome(286))
	//Overflows. Need to switch to use big.Int
	fmt.Println(convertToPalindrome(196196871))
}

func reverseDigits(n uint64) uint64 {

	var remain, ans uint64
	remain, ans = 0, 0

	for n > 0 {
		remain = n % 10
		ans = ans*10 + remain
		n = n / 10
	}
	return ans
}

func isPalindromic(n uint64) bool {

	s := strconv.FormatUint(n, 10)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func convertToPalindrome(n uint64) string {

	var count, sum uint64
	count, sum = 0, n

	for !isPalindromic(sum) {
		i := reverseDigits(sum)
		sum = i + sum
		count++
	}
	return fmt.Sprintf("%v gets palindromic after %v steps: %v", n, count, sum)
}
