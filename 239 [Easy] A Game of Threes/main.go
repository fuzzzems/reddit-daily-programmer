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
	n, err := strconv.ParseUint(r.Text(), 10, 64)
	if err != nil{
		panic(err)
	}

	for n > 1 {
		switch n % 3{
			case 1:
				fmt.Println(n, -1)
				n--
				break;
			case 2:
				fmt.Println(n, 1)
				n++
				break;
		default:
			fmt.Println(n, 0)
			break
		}
		n /= 3;
	}
	fmt.Println(n)
}
