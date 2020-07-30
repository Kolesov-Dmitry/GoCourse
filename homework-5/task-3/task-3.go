package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// isFileExists checks if file located on path is exists
// Input:
//  - path to the file
func isFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// openOrCreateFile opens file if the file exists, otherwise cerates it
// Input:
//  - path to the file
// Output:
//  - file descriptor
//  - error, if unable to open or create the file
func openOrCreateFile(path string) (*os.File, error) {
	var file *os.File
	var err error
	if isFileExists(path) {
		file, err = os.Open(path)
	} else {
		file, err = os.Create(path)
	}

	return file, err
}

// readCSVFile reads CSV file into array of records
// Input:
//  - path is the path of the file to read
//  - comma is the delimeter in CSV file
// Output:
//  - array of records
//  - error, if failed to read file or file has wrong format
func readCSVFile(path string, comma rune) ([][]string, error) {
	// open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read the file content
	reader := csv.NewReader(file)
	reader.Comma = comma

	return reader.ReadAll()
}

// saveCSVToFile saves CSV data into a file
// Input:
//  - path is the path of the file to save CSV
//  - comma is the delimeter in CSV file
//  - data is the records of data
// Output:
//  - error, if failed to write into the file
// Creates file if it doesn't exist
func saveCSVToFile(path string, comma rune, data [][]string) error {
	// open or create the file
	file, err := openOrCreateFile(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// save CSV data
	writer := csv.NewWriter(file)
	writer.Comma = comma

	return writer.WriteAll(data)	
}

func main() {
	lines, err := readCSVFile("input.csv", ';')
	if err == nil {
		for _, line := range lines {
			fmt.Println(line)	
		}

		if err = saveCSVToFile("output.csv", ';', lines); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}