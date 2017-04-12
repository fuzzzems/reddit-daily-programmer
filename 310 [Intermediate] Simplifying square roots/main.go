// [2017-04-12] Challenge #310 [Intermediate] Simplifying square roots
// https://www.reddit.com/r/dailyprogrammer/comments/64y4cf/20170412_challenge_310_intermediate_simplifying/

package main

import "fmt"

func main() {

	// (1√2)/5
	fmt.Println(simplify(2, 5, 5, 10))
	// (15√879)/26
	fmt.Println(simplify(45, 1465, 26, 15))
}

/*
 *  a(√(b*d)) / cd
 */
func simplify(a int, b int, c int, d int) string {

	oa, oc := a, c

	i, r := simplifyRadical(b * d)

	c = c * d
	a = a * i

	gcf := getGCF(a, c)

	if gcf != 1 {
		a, c = a/gcf, c/gcf
	}

	return fmt.Sprintf("Input: (%d√%d)/(%d√%d) \nOutput: (%d√%d)/%d", oa, b, oc, d, a, r, c)
}

/*
 *	index√radicand
 */
func simplifyRadical(radicand int) (int, int) {
	index, n := 1, 2
	for n*n <= radicand {
		if radicand%(n*n) == 0 {
			radicand = radicand / (n * n)
			index = index * n
		} else {
			n++
		}
	}

	return index, radicand
}

func getGCF(a int, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}

	return a
}
