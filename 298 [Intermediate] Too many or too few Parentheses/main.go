// [2017-01-04] Challenge #298 [Intermediate] Too many or too few Parentheses
// https://www.reddit.com/r/dailyprogrammer/comments/5m034l/20170104_challenge_298_intermediate_too_many_or/

package main

import (
	"strings"
	"fmt"
	"bytes"
)

func main(){
	inputs := []string{
	")(asdf)))",
	"((((asdf)))",
	"((((asdf))",
	"(ab)((cd)(asdf)))",
	"(ab)((cd)(asdf)())",
	"(ab)(((cd)(asdf)",
	"(ab)(((cd)(asdf",
	"(ab)(((cd)(asdf)))))",
	}

	for _, v := range inputs{
		if b, i := isBalanced(v); b {
			fmt.Println(v)
		} else{
			fmt.Println(boldError(v, i))
		}
	}

}

func isBalanced(s string) (bool, int) {
	stack := make([]int, 0)

	chars := strings.Split(s, "")

	for k, v := range chars{
		if v == "(" {
			stack = append(stack, k)
		} else if v == ")"{

			if len(stack) == 0 {
				return false, k
			}

			stack = append(stack[:len(stack) - 1])
		}
	}

	if len(stack) == 0 {
		return true, 0
	}

	return false, stack[len(stack) - 1]
}

func boldError(s string, i int) string{

	var b bytes.Buffer

	b.WriteString(s[:i])
	b.WriteString("**")
	b.WriteByte(s[i])
	b.WriteString("**")
	b.WriteString(s[i + 1:])

	return b.String()
}
