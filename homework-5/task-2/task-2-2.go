package main

// Если уж сильно захотелось иметь такой же код, но свой
// я бы вынес его в отдельную функцию и докрутил error в качестве
// возвращаемого значения, чтоб было понятно что же там произошло.

// Более того, не дурно будет ограничить максимальный размер файла
// для чтения таким образом, иначе может не влезть в оперативку

import (
	"errors"
	"fmt"
	"os"
	"io"
)

// Maximum file size to read
const MaxRead = 500_000_000

func readFile(path string) []byte, error {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	
	// check if file size is less than MaxRead
	if fi.Size() > MaxRead {
		return nil, errors.New("File is too big to read")
	}

	buffer := make([]byte, fi.Size())
	if _, err = file.Read(buffer); err != nil {
		return nil, err
	}
	
	return buffer, nil
}

func main() {
	if fileData, err := readFile("fileread.go"); err == nil {
		fmt.Println(string(fileData))
	}
}