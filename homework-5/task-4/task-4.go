package main

import (
	"errors"
	"fmt"
	"os"
//	"io"
)

// isFileExists checks if file located on path is exists
// Input:
//  - path to the file
func isFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// openOrCreateDestFile opens destination file if the file exists, otherwise cerates it
// Input:
//  - path to the file
// Output:
//  - file descriptor
//  - error, if unable to open or create the file
func openOrCreateDestFile(path string) (*os.File, error) {
	var destFile *os.File
	var err error
	if isFileExists(path) {
		destFile, err = os.Open(path)
	} else {
		destFile, err = os.Create(path)
	}

	return destFile, err
}

// getFileSize return size of the given file
func getFileSize(file *os.File) (int64, error) {
	if file == nil {
		return 0, errors.New("Fatal error!")
	}

	stat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return stat.Size(), nil
}

func copyFile(srcFile, destFile *os.File) error {
	if srcFile == nil || destFile == nil {
		return errors.New("Fatal error!")
	}

	srcFileSize, err := getFileSize(srcFile)
	if err != nil {
		return err
	}

	var copyedBytes int64 = 0
	buffer := make([]byte, 100_000)	
	for copyedBytes < srcFileSize {		
		// read data from source file
		c, err := srcFile.Read(buffer)
		if err != nil {
			return err
		}

		// write data to destination file
		_, err = destFile.Write(buffer[:c])
		if err != nil {
			return err
		}
		
		copyedBytes += int64(c)

		srcFile.Seek(copyedBytes, 0)
		destFile.Seek(copyedBytes, 0)
	}

	return nil
}

func Copy(src, dest string) error {
	// open source file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// open or create destination file
	destFile, err := openOrCreateDestFile(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	err = copyFile(srcFile, destFile)
	if err != nil {
		os.Remove(dest)
	}

	return err
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: copy (SRC_FILE_PATH) (DEST_FILE_PATH)")
		return
	}

	src := os.Args[1]
	dest := os.Args[2]
	if err := Copy(src, dest); err != nil {
		fmt.Println(err)
	}
}