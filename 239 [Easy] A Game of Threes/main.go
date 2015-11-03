// [2015-11-02] Challenge #239 [Easy] A Game of Threes
// https://www.reddit.com/r/dailyprogrammer/comments/3r7wxz/20151102_challenge_239_easy_a_game_of_threes/

package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func main(){
	r := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a number")
	r.Scan()
	input, err := strconv.ParseUint(r.Text(), 10, 64)
	if err != nil{
		panic(err)
	}

	for input > 1 {
		c := input % 3
		switch	c {
			case 0:
				fmt.Println(input, 0)
				break
			case 1:
				fmt.Println(input, -1)
				break
			case 2:
				fmt.Println(input, 1)
				input += 1
				break
		}
		input /= 3
	}

	if input ==  1 {
		fmt.Println(1)
	}
}
