// [2016-11-07] Challenge #291 [Easy] Goldilocks' Bear Necessities
// https://www.reddit.com/r/dailyprogrammer/comments/5bn0b7/20161107_challenge_291_easy_goldilocks_bear/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	l := strings.Split(scanner.Text(), " ")

	w, err := strconv.Atoi(l[0])
	t, err := strconv.Atoi(l[1])
	if err != nil {
		panic(err)
	}

	for i := 1; scanner.Scan(); i++ {

		q := strings.Split(scanner.Text(), " ")

		wn, err := strconv.Atoi(q[0])
		tn, err := strconv.Atoi(q[1])
		if err != nil {
			panic(err)
		}

		if w <= wn && t >= tn {
			fmt.Printf("%d ", i)
		}
	}

	fmt.Println()
}
