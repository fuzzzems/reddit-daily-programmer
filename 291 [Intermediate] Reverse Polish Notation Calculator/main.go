// [2016-11-09] Challenge #291 [Intermediate] Reverse Polish Notation Calculator
// https://www.reddit.com/r/dailyprogrammer/comments/5c5jx9/20161109_challenge_291_intermediate_reverse/

package main

import (
	"fmt"
	"github.com/alediaferia/stackgo"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%.0f\n", calc("0.5 1 2 ! * 2 1 ^ + 10 + *"))
	fmt.Printf("%.0f\n", calc("1 2 3 4 ! + - / 100 *"))
	fmt.Printf("%.0f\n", calc("100 807 3 331 * + 2 2 1 + 2 + * 5 ^ * 23 10 558 * 10 * + + *"))
}

func calc(input string) interface{} {
	stack := stackgo.NewStack()
	tokens := strings.Split(input, " ")

	for _, v := range tokens {
		if isOperator(v) {
			if v == "!" {
				o0, _ := stack.Pop().(float64)
				for i := o0; i > 1; {
					i--
					o0 *= i
				}
				stack.Push(o0)

			} else {
				o1, _ := stack.Pop().(float64)
				o2, _ := stack.Pop().(float64)

				switch v {
				case "+":
					stack.Push(o2 + o1)
				case "-":
					stack.Push(o2 - o1)
				case "*", "x":
					stack.Push(o2 * o1)
				case "/":
					stack.Push(o2 / o1)
				case "//":
					stack.Push(float64(int(o2) / int(o1)))
				case "%":
					stack.Push(float64(int(o2) % int(o1)))
				case "^":
					stack.Push(math.Pow(o2, o1))
				default:
					fmt.Println("Unsupported operator")
				}
			}
		} else {
			n, err := strconv.ParseFloat(v, 64)
			if err != nil {
				panic(err)
			}
			stack.Push(n)
		}
	}
	return stack.Pop()
}

func isOperator(s string) bool {
	switch s {
	case "+", "-", "*", "x", "/", "//", "%", "^", "!":
		return true
	default:
		return false
	}
}
