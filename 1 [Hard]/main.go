// [difficult] challenge #1
// http://www.reddit.com/r/dailyprogrammer/comments/pii6j/difficult_challenge_1/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	min, max := 1, 100
	count := 1
	for {

		guess := (min + max) / 2
		fmt.Printf("Is your number %d \n", guess)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		response := scanner.Text()

		if response == "yes" {
			fmt.Printf("Yay, we guessed your number, %d, in %d guesses \n", guess, count)
			break
		} else if response == "no" {
			count++
			fmt.Println("too high or too low?")
			scanner.Scan()
			response = scanner.Text()
			if response == "high" || response == "too high" {
				max = guess
			} else if response == "low" || response == "too low" {
				min = guess
			} else {
				fmt.Println("Unregconized input. Valid input: high, low")
			}
		} else {
			fmt.Println("Unregconized input. Valid input: yes, no")
		}
	}
}
