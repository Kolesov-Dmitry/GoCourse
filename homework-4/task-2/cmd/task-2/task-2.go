package main

import (
	"fmt"
	"sort"
	"homework.com/v4/task-2/internal/app/phonebook"
)

func main() {
	book := make(phonebook.PhoneBook, 0, 3)

	// insert some elements
	book.Append("Masha", "88 88 888 888")
	book.Append("Olya", "99 99 999 999")
	book.Append("Dima", "77 77 777 777")
			
	fmt.Println("До сортировки:")
	fmt.Print(book)

	sort.Sort(book)

	fmt.Println("")
	fmt.Println("После сортировки:")
	fmt.Print(book)
}