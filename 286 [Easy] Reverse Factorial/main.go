// [2016-10-03] Challenge #286 [Easy] Reverse Factorial
// https://www.reddit.com/r/dailyprogrammer/comments/55nior/20161003_challenge_286_easy_reverse_factorial/

package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(formatOutput(1))
	fmt.Println(formatOutput(0))
	fmt.Println(formatOutput(-1))
	fmt.Println(formatOutput(150))
	fmt.Println(formatOutput(120))
	fmt.Println(formatOutput(3628800))
	fmt.Println(formatOutput(479001600))
	fmt.Println(formatOutput(6))
	fmt.Println(formatOutput(18))
}

func formatOutput(n int) string {
	i, err := reverseFactorial(n)

	if err != nil {
		return fmt.Sprintf("%d %s", n, err)
	}
	return fmt.Sprintf("%d = %d!", n, i)
}

func reverseFactorial(n int) (int, error) {
	f := 1

	for ; n%f == 0 && n != 0; f++ {

		if n/f == 1 {
			return f, nil
		}

		n = n / f
	}

	return f, errors.New("NONE")

}
