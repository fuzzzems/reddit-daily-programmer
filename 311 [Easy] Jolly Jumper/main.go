// [2017-04-17] Challenge #311 [Easy] Jolly Jumper
// https://www.reddit.com/r/dailyprogrammer/comments/65vgkh/20170417_challenge_311_easy_jolly_jumper/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Functions nested like this to demonstrate usage.
	fmt.Println(isJollyToString(isJolly(strings.Split("8 1 6 -1 8 9 5 2 7", " "))))
	fmt.Println(isJollyToString(isJolly(strings.Split("4 1 4 2 3", " "))))
	fmt.Println(isJollyToString(isJolly(strings.Split("5 1 4 2 -1 6", " "))))
	fmt.Println(isJollyToString(isJolly(strings.Split("4 19 22 24 21", " "))))
	fmt.Println(isJollyToString(isJolly(strings.Split("4 19 22 24 25", " "))))
	fmt.Println(isJollyToString(isJolly(strings.Split("4 2 -1 0 2", " "))))

}

func isJollyToString(b bool, s []string) string {
	if b {
		return fmt.Sprintf("%v JOLLY", s)
	}
	return fmt.Sprintf("%v NOT JOLLY", s)
}

func isJolly(data []string) (bool, []string) {
	n, err := strconv.Atoi(data[0])
	if err != nil {
		return false, data
	}

	ans := make([]bool, n-1)

	for i := 1; i <= len(ans); i++ {
		k, err := strconv.Atoi(data[i])
		if err != nil {
			return false, data
		}

		q, err := strconv.Atoi(data[i+1])
		if err != nil {
			return false, data
		}

		a := abs(k - q)

		if a > len(ans) || ans[a-1] == true {
			return false, data
		}
		ans[a-1] = true
	}

	return true, data
}

func abs(i int) int {
	if i > 0 {
		return i
	}

	return i * -1
}
