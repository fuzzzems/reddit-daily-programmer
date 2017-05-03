// [2017-05-01] Challenge #313 [Easy] Subset sum
// https://www.reddit.com/r/dailyprogrammer/comments/68oda5/20170501_challenge_313_easy_subset_sum/

package main

import (
	"fmt"
)

func main() {
	// False
	fmt.Println(sumZero([]int{1, 2, 3}))
	fmt.Println(sumZero([]int{-5, -3, -1, 2, 4, 6}))
	fmt.Println(sumZero([]int{}))

	// True
	fmt.Println(sumZero([]int{-1, 1}))
	fmt.Println(sumZero([]int{-97364, -71561, -69336, 19675, 71561, 97863}))
	fmt.Println(sumZero([]int{-53974, -39140, -36561, -23935, -15680, 0}))
}

func sumZero(data []int) bool {
	set := make(map[int]bool, len(data))

	for _, value := range data {
		n := abs(value)

		if set[n] == true || n == 0 {
			return true
		}
		set[n] = true
	}

	return false
}

func abs(n int) int {
	if n > 0 {
		return n
	} else {
		return n * -1
	}
}
