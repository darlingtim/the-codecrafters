package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(os.Args)
enterArgument:
	fmt.Println("Please enter an operation and the two number with space between them.")
	argument := bufio.NewScanner(os.Stdin)

	if argument.Scan() {
		arg := strings.Fields(argument.Text())

		funcArg1, _ := strconv.Atoi(arg[1])
		funcArg2, _ := strconv.Atoi(arg[2])
		if len(arg) == 1 && argument.Text() == "quit" {
			fmt.Println("Thanks for using our calculator")
		} else if len(arg) < 3 || len(arg) > 3 {
			fmt.Println("incorrect number of argument. E.g. add 3 5")
			goto enterArgument

		} else if strings.ContainsAny(arg[1], "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") || strings.ContainsAny(arg[2], "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			fmt.Println("Function arguments must be a number")
			goto enterArgument
		} else {
			if arg[0] == "add" {
				result := fmt.Sprint(add(funcArg1, funcArg2))
				fmt.Println(result)
				goto enterArgument
			}

			if arg[0] == "sub" {
				result := fmt.Sprint(sub(funcArg1, funcArg2))
				fmt.Println(result)
				goto enterArgument
			}

			if arg[0] == "mul" {
				result := fmt.Sprint(mul(funcArg1, funcArg2))
				fmt.Println(result)
				goto enterArgument
			}

			if arg[0] == "div" {
				result := fmt.Sprint(div(funcArg1, funcArg2))
				fmt.Println(result)
				goto enterArgument
			}

		}

	} else {
		return
	}

	//if len(os.Args)<
}

func add(i int, j int) int {
	return i + j
}

func sub(i int, j int) int {
	return i - j
}

func mul(i int, j int) int {
	return i * j
}

func div(i int, j int) int {
	return i / j
}
