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
	"unicode"
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

	var processed []string
	linesRead := 0
	linesRemoved := 0

	for _, line := range lineOfWords {
		linesRead++

		line = replaceTODO(line)
		line = convertAllCaps(line)
		line = convertLowerCase(line)
		if removeBlankDash(line) {
			linesRemoved++
			continue
		}
		line = truncateLongLine(line)

		processed = append(processed, line)

	}

	// writing HEADER to file
	file, _ := os.OpenFile(outputFile, os.O_WRONLY, 0644)
	fmt.Fprintln(file, "PROCESSED SENTINEL FIELD REPORT")

	// write processed lines to output file
	for i, processedLine := range processed {
		fmt.Fprintf(file, "%d. %s\n", i+1, processedLine)
	}

	// Print Terminal Summary
	fmt.Println("Summary Block")
	fmt.Println("Lines read: ", linesRead)
	fmt.Println("Lines written: ", len(processed))
	fmt.Println("Lines removed: ", linesRemoved)
	fmt.Print("Rules applied:\n", "1. [Replace TODO: with ✦ ACTION:]\n", "2. [Convert ALL CAPS lines to Title Case]\n", "3. [Convert all lowercase lines to uppercase]\n",
		"4. [Remove lines that are only dashes or blanks]\n",
		"5. [Flag any line longer than 80 characters with [TRUNCATED] at the end]\n")

	//os.WriteFile(outputFile, []byte(line+"\n"), 0644)

}

//transformation functions

// rule 5 Flag lines longer than 80 characters.
func truncateLongLine(line string) string {
	if len(line) > 80 {
		truncated := line[:81]
		return truncated + " [Truncated]"
	}
	return line
}

// Rule 4 Remove lines that are only dashes or blank
func removeBlankDash(line string) bool {
	l := strings.TrimSpace(line)
	if l == "" || allDashes(l) {
		return true
	}
	return false
}

func allDashes(s string) bool {
	for _, v := range s {
		if v != '-' {
			return false
		}
	}
	return len(s) > 0
}

// Rule 2 Convert ALL CAPS lines to Title Case
func convertAllCaps(line string) string {
	if isAllCaps(line) {
		return toTitleCase(line)
	}
	return line
}

func isAllCaps(s string) bool {
	hasLetter := false
	for _, v := range s {
		if unicode.IsLetter(v) {
			hasLetter = true
			if !unicode.IsUpper(v) {
				return false
			}
		}
	}
	return hasLetter
}

func toTitleCase(s string) string {
	return strings.Title(strings.ToLower(s))
}

// Rule 1 replace TODO: with ACTION:
func replaceTODO(line string) string {
	return strings.Replace(line, strings.ToLower("TODO:"), "ACTION:", 1)
}

// Rule 3
func convertLowerCase(line string) string {
	if line == strings.ToLower(line) && strings.TrimSpace(line) != "" {
		return strings.ToUpper(line)
	}
	return line
}
