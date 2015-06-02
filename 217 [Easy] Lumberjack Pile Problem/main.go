// [2015-06-01] Challenge #217 [Easy] Lumberjack Pile Problem
// http://www.reddit.com/r/dailyprogrammer/comments/3840rp/20150601_challenge_217_easy_lumberjack_pile/

package main

import (
	"fmt"
)

func main() {

	n, m := 3, 7

	var grid = []int{1, 1, 1, 2, 1, 3, 1, 4, 1}

	for i := 0; i < m; i++ {

		l := findSmall(grid)
		grid[l] = grid[l] + 1

	}

	for i := range grid {
		fmt.Print(grid[i], " ")
		if (i+1)%n == 0 {
			fmt.Println()
		}
	}

}

func findSmall(grid []int) int {

	small := 0

	for index := range grid {
		if grid[small] > grid[index] {
			small = index
		}
	}

	return small
}
