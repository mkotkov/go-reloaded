package main

import (
    "fmt"
    "os"
)

func main() {
	// readFile()
	changeFile()
}

// func readFile (){
// 	args := os.Args[1:]
// 	filePath := args[0]

// 	file, err := os.ReadFile(filePath) // For read access.
// 	fileString := string(file)
// 		if err != nil {
// 		fmt.Println("File isn't")
// 	}else{
// 		fmt.Println(fileString)
// 	}
// }

func changeFile(){

	args := os.Args[1:]
	sampleFilePath := args[0]
	resultFilePath := args[1]

	file, err := os.ReadFile(sampleFilePath)
	
	fileString := string(file)

	data := []byte(fileString) // Change

	err = os.WriteFile(resultFilePath, []byte(data), 0644) // Write the modified content to the result file
	if err != nil {
		fmt.Println("Error writing to the result file:", err)
		return
	}

	fmt.Println("File successfully written to", resultFilePath)

}