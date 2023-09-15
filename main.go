package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	changeFile()
}


func changeFile(){

	args := os.Args[1:]
	sampleFilePath := args[0]
	resultFilePath := args[1]

	file, err := os.ReadFile(sampleFilePath)
	
	fileString := string(file)
	
	fileString = cap(fileString)
	fileString = up(fileString)

    // cleanedText := clean(fileString)
	
	data := []byte(fileString)

	err = os.WriteFile(resultFilePath, []byte(data), 0644) // Write the modified content to the result file
	if err != nil {
		fmt.Println("Error writing to the result file:", err)
		return
	}

	fmt.Println("File successfully written to", resultFilePath)

}



func cap(input string) string{
	
	re := regexp.MustCompile(`\b\w+\s+\(cap\)`)

	output := re.ReplaceAllStringFunc(input, func(matched string) string {
		// Получите слово перед "(cap)"
		words := strings.Fields(matched) // Разделите строку на слова
		if len(words) > 1 {
			word := words[len(words)-2]  // Получите предпоследнее слово
			capitalized := strings.ToUpper(word[:1]) + word[1:] // Сделайте первую букву заглавной
			return strings.Replace(matched, word, capitalized, 1)
		}
		return matched
	})

	return output
}


func up(input string) string{
	n := 1
	// re := regexp.MustCompile(`(\w+\s+\(cap(?:,\s*(\d+))?\))`)
	exp := fmt.Sprintf(`((\w+\s+){%d}\(cap(?:,\s*(\d+))?\))`, n)

	re:= regexp.MustCompile(exp)

	output := re.ReplaceAllStringFunc(input, func(matched string) string {
		submatches := re.FindStringSubmatch(matched)

		if len(submatches) >= 2 {
			wordList := strings.Split(submatches[1], " ")
			numWords := 1 

			if len(submatches) == 3 {
				// Если указано число, преобразуем его в int
				n, err := strconv.Atoi(submatches[2])
				if err == nil {
					numWords = n
				}
			}

			if numWords > len(wordList) {
				numWords = len(wordList)
			}

			for i := 0; i < numWords; i++ {
				word := wordList[i]
				capitalized := strings.ToUpper(word[:1]) + word[1:]
				wordList[i] = capitalized
			}

			return strings.Join(wordList, " ")
		}

		return matched
	})

	return output
}

// func low(input string) string{
// 	re := regexp.MustCompile(`\b\w+\s+\(up\)`)
	
// 	output := re.ReplaceAllStringFunc(input, func(matched string) string {
		
// 		words := strings.Fields(matched)
		
// 		if len(words) > 1 {
// 			word := words[len(words)-2]                         
// 			capitalized := strings.ToUpper(word)
// 			return strings.Replace(matched, word, capitalized, 1)
// 		}
// 		return matched
// 	})

// 	return output
// }


// func clean(input string) string{
// 	re := regexp.MustCompile(`\s+([,.;:!?])| \(([^)]+)\)`)
//   	output := re.ReplaceAllString(input, "$1")
// 	return output
// }

