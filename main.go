package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// for testing: go run . sample.txt result.txt
func main() {
	changeFile()
}

// read sample.txt args[0], change and save to result.txt
func changeFile() {
	args := os.Args[1:]
	sampleFilePath := args[0]
	resultFilePath := args[1]

	file, err := os.ReadFile(sampleFilePath)

	fileString := string(file)
	fileString = processText(fileString)
	fileString = clean(fileString)

	data := []byte(fileString)

	err = os.WriteFile(resultFilePath, []byte(data), 0o644) // Write the modified content to the result file
	if err != nil {
		fmt.Println("Error writing to the result file:", err)
		return
	}

	fmt.Println("File successfully written to", resultFilePath)
}

// chacking and change text
func processText(input string) string {
	inputArr := strings.Fields(string(input)) // splits the string
	var count int

	for i := 0; i < len(inputArr); i++ { // takes every element separately from the slice input
		if inputArr[i] == "(hex)" {
			r, _ := strconv.ParseInt(inputArr[i-1], 16, 64) // converts the string from hex to dec
			inputArr[i-1] = strconv.FormatInt(r, 10)        // converts int to string
		}
		if inputArr[i] == "(bin)" {
			q, _ := strconv.ParseInt(inputArr[i-1], 2, 64) // converts the string from bin to dec
			inputArr[i-1] = strconv.FormatInt(q, 10)       // converts int to string
		}

		if inputArr[i] == "(cap)" {
			inputArr[i-1] = strings.Title(inputArr[i-1])
		}

		if inputArr[i] == "(cap," {
			count, _ = strconv.Atoi(strings.Replace(inputArr[i+1], ")", "", 1))
			for e := 1; e <= count; e++ {
				inputArr[i-e] = strings.Title(inputArr[i-e])
			}
		}

		if inputArr[i] == "(low)" {
			inputArr[i-1] = strings.ToLower(inputArr[i-1])
		}

		if inputArr[i] == "(low," {
			count, _ = strconv.Atoi(strings.Replace(inputArr[i+1], ")", "", 1))
			for e := 1; e <= count; e++ {
				inputArr[i-e] = strings.ToLower(inputArr[i-e])
			}
			inputArr[i] = strings.Replace(inputArr[i], inputArr[i], "", 1)
			inputArr[i+1] = strings.Replace(inputArr[i+1], inputArr[i+1], "", 1)
		}
		if inputArr[i] == "(up)" {
			inputArr[i-1] = strings.ToUpper(inputArr[i-1])
		}
		if inputArr[i] == "(up," {
			count, _ = strconv.Atoi(strings.Replace(inputArr[i+1], ")", "", 1))
			for e := 1; e <= count; e++ {
				inputArr[i-e] = strings.ToUpper(inputArr[i-e])
			}
			inputArr[i] = strings.Replace(inputArr[i], inputArr[i], "", 1)
			inputArr[i+1] = strings.Replace(inputArr[i+1], inputArr[i+1], "", 1)
		}
	}

	justString := strings.Join(inputArr, " ")

	// check and change article
	re := regexp.MustCompile(`a\s+([aeiouhAEIOUH])`)

	output := re.ReplaceAllStringFunc(justString, func(matched string) string {
		submatches := re.FindStringSubmatch(matched)
		if len(submatches) > 1 {
			return "an " + strings.TrimSpace(submatches[1])
		}
		return matched
	})
	return output
}

// clean extra spacing, double-spacing and hashtags (cap, low, etc.)
func clean(input string) string {
	re := regexp.MustCompile(`\s+([,.;:!?])|\s+\(([^)]+)\)|('...''!?')|\w\s{2}`)
	output := re.ReplaceAllString(input, "$1")
	return output
}
