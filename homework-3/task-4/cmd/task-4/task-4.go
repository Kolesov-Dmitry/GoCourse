package main

// This task was tricky one 
// I almost dived in OOP with interfaces and dependency inversion 
// and exstracted the Saver and Loader entities. But stoped myself 
// in time and replaced all this stuff with two simple functions :)

import (
	"fmt"	
	"homework.com/v3/task-4/internal/app/phonebook"
)

func main() {
	// create new phonebook
	book := phonebook.NewPhoneBook("phonebook.json")

	// insert some elements
	book.Append("Dima", "77 77 777 777")
	book.Append("Olya", "88 88 888 888")
	book.Append("Masha", "99 99 999 999")
	
	anotherBook := phonebook.LoadPhoneBook("phonebook.json")

	phone, _ := anotherBook.Find("Dima")
	fmt.Println(phone)
}