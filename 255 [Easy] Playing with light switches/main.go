// [2016-02-22] Challenge #255 [Easy] Playing with light switches
// https://www.reddit.com/r/dailyprogrammer/comments/46zm8m/20160222_challenge_255_easy_playing_with_light/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("lots_of_switches.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	n, err := strconv.ParseInt(scanner.Text(), 10, 32)
	if err != nil {
		panic(err)
	}

	switches := make([]int, n)

	for scanner.Scan() {
		start, end := getNum(scanner.Text())
		for i := start; i <= end; i++ {
			switches[i] = switches[i] ^ 1
		}
	}

	var sum int
	for _, i := range switches {
		sum = i + sum
	}
	fmt.Println(sum)
}

func getNum(s string) (int64, int64) {
	q := strings.Fields(s)

	start, err := strconv.ParseInt(q[0], 10, 32)
	if err != nil {
		panic(err)
	}

	end, err := strconv.ParseInt(q[1], 10, 32)

	if err != nil {
		panic(err)
	}

	if start > end {
		start, end = end, start
	}

	return start, end
}
