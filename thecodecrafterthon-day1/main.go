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
	fmt.Print("Welcome to CodeCrafters Calculator..\n This calculator can only help you to add, subtract, multiply and divide any two numbers of your choice.\n")
	exampleOperations := []string{
		"\n",
		"Here are the examples of all operations you can perform.\n",
		"\n",
		"add 2 5 - add 2 and 5 to give you 7\n",
		"sub 5 3 - subtract two 5 and 3 to give you 2\n",
		"mul 7 4 - multiplies 7 and 4 to give you 28\n",
		"div 8 2 - divides 8 with 2 to give you 4\n",
		"help - will show you all you need to effectively use this calculator\n",
		"quit - will stop and exit the calculator program\n",
	}

	fmt.Print("\n", exampleOperations, "\n", "\n")

enterArgument:
	fmt.Println("Please enter an operation")
	argument := bufio.NewScanner(os.Stdin)

	if argument.Scan() {
		arg := strings.Fields(argument.Text())

		if len(arg) == 1 {
			if argument.Text() == "quit" {
				fmt.Print("\n", "Thanks for using codeCrafters calculator\n")
				return
			} else if argument.Text() == "help" {
				//helpFile, err3 := os.ReadFile("README.md")
				/*if err3 != nil {
					fmt.Println("An error occured in fetching help file, try again or carry out another operation")
					goto enterArgument
				}*/
				fmt.Print("\n", exampleOperations, "\n")
				goto enterArgument

			} else if argument.Text() == "add" {
				fmt.Print("\n", "you did not provide the two numbers that you want to add.", "\n")
				goto enterArgument

			} else if argument.Text() == "sub" {
				fmt.Print("\n", "you did not provide the two numbers that you want to subtract.", "\n")
				goto enterArgument

			} else if argument.Text() == "mul" {
				fmt.Print("\n", "you did not provide the two numbers that you want to multiply.", "\n")
				goto enterArgument

			} else if argument.Text() == "div" {
				fmt.Print("\n", "you did not provide the two numbers that you want to divide.", "\n")
				goto enterArgument

			} else {
				fmt.Print("\n", "you entered a wrong operation.", "\n")
				goto enterArgument

			}
		}

		if len(arg) == 2 || len(arg) > 3 {
			fmt.Print("\n", "incorrect number of argument. E.g. add 3 5", "\n")
			goto enterArgument

		} else if len(arg) == 3 {

			funcArg1, err := strconv.Atoi(arg[1])
			funcArg2, err1 := strconv.Atoi(arg[2])

			if err != nil || err1 != nil {
				fmt.Print("\n", "An error occured, Try again with the right parameters", "\n")

			}

			if strings.ContainsAny(arg[1], "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") || strings.ContainsAny(arg[2], "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
				fmt.Println("Function arguments must be a number")
				goto enterArgument
			}

			if arg[0] == "add" {
				result := fmt.Sprint(add(funcArg1, funcArg2))
				fmt.Print("\n", result, "\n", "\n")
				goto enterArgument
			}

			if arg[0] == "sub" {
				result := fmt.Sprint(sub(funcArg1, funcArg2))
				fmt.Print("\n", result, "\n", "\n")
				goto enterArgument
			}

			if arg[0] == "mul" {
				result := fmt.Sprint(mul(funcArg1, funcArg2))
				fmt.Print("\n", result, "\n", "\n")
				goto enterArgument
			}

			if arg[0] == "div" {
				result := fmt.Sprint(div(funcArg1, funcArg2))
				fmt.Print("\n", result, "\n", "\n")
				goto enterArgument
			}
			fmt.Print("\n", "The math operation you provided is either invalid or not supported at the moment", "\n")
			goto enterArgument
		}

	} else {
		return
	}

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
