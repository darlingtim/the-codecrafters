package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binary(b string) int64 {
	result, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		fmt.Printf("An error occured while converting from binary to decimal:\n %v", err)

	}
	return result
}

func hexadecimal(b string) int64 {
	result, err := strconv.ParseInt(b, 16, 64)
	if err != nil {
		fmt.Printf("an error occured while converting from hexadecimal to decimal:\n %v", err)

	}
	return result
}

func decimal(dec string) string {
	return dec
}

// main function
func main() {

	fmt.Print("Welcome to CodeCrafters NumberBase Converter..\n This Converter can only help you convert binary and hexadecimal to decimal and can also give you the binary and hexadecimal value of any decimal number of your choice .\n")
	exampleOperations := []string{
		"\n",
		"Here are the examples of all operations you can perform.\n",
		"\n",
		"convert 1E hex - Decimal: 30\n",
		"convert 10 bin - Decimal: 2\n",
		"convert 255 dec - Binary:  11111111  Hex: FF\n",
		"help - will show you all you need to effectively use this calculator\n",
		"quit - will stop and exit the calculator program\n",
	}

	fmt.Print("\n", exampleOperations, "\n", "\n")

enterArgument:
	fmt.Println("Please enter an operation")

	userInput := bufio.NewScanner(os.Stdin)
	arg := strings.Fields(userInput.Text())

	if userInput.Scan() {
		if len(arg) == 1 {
			if userInput.Text() == "quit" {
				fmt.Print("\n", "Thanks for using codeCrafters Base Converter\n")
				return
			} else if userInput.Text() == "help" {
				//helpFile, err3 := os.ReadFile("README.md")
				/*if err3 != nil {
					fmt.Println("An error occured in fetching help file, try again or carry out another operation")
					goto enterArgument
				}*/
				fmt.Print("\n", exampleOperations, "\n")
				goto enterArgument

			} else {
				fmt.Printf("%v is not a valid operation. First Arguement must be convert, help or quit", userInput.Text())
				goto enterArgument
			}
		} else if len(arg) == 2 || len(arg) > 3 {
			fmt.Println("invalid number of inputs recievied. enter the right format or press help to see the needed input format")
			goto enterArgument
		} else {
			if arg[0] == "convert" || arg[0] == "Convert" || arg[0] == "CONVERT" {
				switch arg[2] {
				case "bin":
					binary(arg[1])
				case "hex":
					hexadecimal(arg[1])
				case "dec":
					decimal(arg[1])
				default:
					fmt.Printf("converting %v is not currently support", arg[2])
					goto enterArgument

				}
			}
		}

		fmt.Println(len(userInput.Text()))

	}

}
