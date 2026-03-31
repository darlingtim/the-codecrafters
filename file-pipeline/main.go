/*CodeCrafters — Operation Gopher Protocol
Module: File Pipeline
Author: [Timothy Ododo]
Squad:  [Debuggers]
*/

// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: [Debuggers]
// ───────────────────────────────────────────
// Input line types:
//  1. Lines in ALL CAPS
//  2. Lines in all lowercase
//  3. Lines starting with TODO:
//  4. Lines that are only dashes or blanks
//
//
// Transformation rules (in order):
//   1. [Replace TODO: with ✦ ACTION:]
//   2. [Convert ALL CAPS lines to Title Case]
//   3. [Convert all lowercase lines to uppercase]
//   4. [Remove lines that are only dashes or blanks]
//   5. [Flag any line longer than 80 characters with [TRUNCATED] at the end]
//
// Output format:
//   [Header: PROCESSED SENTINEL FIELD REPORT]
//   [Line numbering format: "1."]
//   [Summary block:  Lines read : [number] ✦ Lines written : [number] ✦ Lines removed : [number] ✦ Rules applied : [] ]
//
// Terminal summary fields:
//   [List what your squad agreed on]
// ═══════════════════════════════════════════

// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [Your Name]
// Squad:  [Debuggers]

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args

	if len(arg) < 3 || arg == nil {
		fmt.Println("Usage: go run main.go input.txt output.txt")
		return
	}

	inputFile := arg[1]
	outputFile := arg[2]

	if inputFile == outputFile {
		fmt.Println("Input and output cannot be the same file.")
	}
	// stat for input and ouput
	inputInfo, inputFileTypeError := os.Stat(inputFile)
	_, outputFileTypeError := os.Stat(outputFile)
	//empty input file check

	//check if argument for input is a file or directory
	if inputFileTypeError != nil {
		fmt.Println("Cannot write to input: path is a directory, not a file.")
		return
	}

	//check if argument for output is a file or directory
	if outputFileTypeError != nil {

		fmt.Println("Cannot write to output: path is a directory, not a file.")
		return
	}

	inputData, err := os.ReadFile(inputFile)
	lineOfWords := strings.Split(string(inputData), "\n")

	// Error reading input file or file does not exist
	if err != nil {
		fmt.Println("input.txt file does not exist or Error in reading input.txt file")
	}

	if inputInfo.Size() == 0 {
		fmt.Println("Input file is empty. Nothing to process.")
	}

	for _, line := range lineOfWords {
		if strings.Split(line, "-") == nil || strings.Split(line, " ") == nil {

		} else {

			file, _ := os.OpenFile(outputFile, os.O_WRONLY, 0644)
			fmt.Fprintln(file, line)

			//os.WriteFile(outputFile, []byte(line+"\n"), 0644)
		}
	}

}
