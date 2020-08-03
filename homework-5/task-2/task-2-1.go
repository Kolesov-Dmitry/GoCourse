package main

// Первое, что приходит на ум, это заменить всё 
// на функцию ReadFile из пакета io/util, 
// поскольку приведённый код практически дублирует её.

import (
	"fmt"
	"io/util"
)

func main() {
	if fileData, err := ioutil.ReadFile("fileread.go"); err == nil {
		fmt.Println(string(fileData))
	}
}