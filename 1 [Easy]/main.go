// [easy] challenge #1
// http://www.reddit.com/r/dailyprogrammer/comments/pih8x/easy_challenge_1/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("What is your name?")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("How old are you?")
	scanner.Scan()
	age := scanner.Text()

	fmt.Println("What is your username?")
	scanner.Scan()
	un := scanner.Text()

	output := fmt.Sprintf("Your name is %s, you are %s years old, and your username is %s", name, age, un)

	fmt.Println(output)

	file, err := os.Create("./output.txt")
	if err != nil {
		panic(err)
	}

	file.WriteString(output)
}
