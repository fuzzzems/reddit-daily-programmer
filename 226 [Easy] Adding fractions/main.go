// [2015-08-03] Challenge #226 [Easy] Adding fractions
// http://www.reddit.com/r/dailyprogrammer/comments/3fmke1/20150803_challenge_226_easy_adding_fractions/

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	cl := bufio.NewReader(os.Stdin)

	num, err := cl.ReadString('\n')

	if err != nil {
		panic(err)
	}

	num = strings.TrimSpace(num)

	n, err := strconv.Atoi(num)

	if err != nil {
		panic(err)
	}

	fractions := make([]string, n)

	for i := 0; i < n; i++ {
		fractions[i], err = cl.ReadString('\n')
		if err != nil {
			panic(err)
		}
	}

	total := fractions[0]
	for i := 0; i < n; i++ {
		if i != n - 1 {
			total = addFraction(total, fractions[i+1])
		}
	}


	fmt.Println(total)
}

func addFraction(a, b string) string {

	i := regexp.MustCompile("/").Split(a, 2)
	j := regexp.MustCompile("/").Split(b, 2)

	i[1] = strings.TrimSpace(i[1])
	j[1] = strings.TrimSpace(j[1])

	aa, _ := strconv.Atoi(i[0])
	ad, _ := strconv.Atoi(i[1])
	bb, _ := strconv.Atoi(j[0])
	bd, _ := strconv.Atoi(j[1])

	an, bn, bd := crossMultiply(aa, ad, bb, bd)

	total := an + bn

	gcf := gcf(total, bd)

	return fmt.Sprintf("%d/%d", total/gcf, bd/gcf)
}

func crossMultiply(an, ad, bn, bd int) (int, int, int) {

	bn *= ad
	an *= bd
	ad *= bd

	return an, bn, ad
}

func gcf(a, b int) int {

	c := -1

	for {
		a, b = b, a%b

		if b == 0 {
			return c
		}

		c = b
	}

	return -1
}
