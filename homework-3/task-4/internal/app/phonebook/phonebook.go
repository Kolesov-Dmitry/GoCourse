package phonebook

import (
	"errors"	
	"log"
)

// PhoneBook structure
type PhoneBook struct {
	// path to file to dump phone book
	dumpFilePath string

	// abonents contains the phone numbers
	abonents map[string]string
}

// NewPhoneBook creates empty PhoneBook
// Input:
//  - path is the path to the JSON file to save abonents data
// Output:
//  - pointer to the created instance of PhoneBook
func NewPhoneBook(path string) *PhoneBook {	
	if len(path) == 0 {		
		path = "phonebook.json"				
	} 
	
	return &PhoneBook {
		dumpFilePath: path,
		abonents: make(map[string]string),
	}
}

// LoadPhoneBook creates PhoneBook instance and loads the abonents data from the file into it
// Input:
//  - path is the path to the JSON file
// Output:
//  - pointer to the created instance of PhoneBook
// Creates new PhoneBook if path is empty
func LoadPhoneBook(path string) *PhoneBook {
	// if path is not empty load the abonents from the file
	if len(path) != 0 {		
		return &PhoneBook {
			dumpFilePath: path,
			abonents: loadFromJsonFile(path),			
		}		
	} 

	// create empty PhoneBook if path is empty
	return NewPhoneBook(path)
}

// Append inserts new number into the phone book
// Input:
//  - abonent is the name of the abonent
//  - number is the phone number of the abonent
// Rewrite phone number if abonent is already exists in the phone book
func (pb *PhoneBook) Append(abonent, number string) {
	pb.abonents[abonent] = number

	err := saveToJsonFile(pb.dumpFilePath, pb.abonents)
	if err != nil {
		// needs some sane error handling		
		log.Printf("Unabel to save the phone book to the file &v", err)
	}
}

// Find searches for the phone number of the abonent in the phone book
// Input:
//  - abonent is the name of the abonent
// Output:
//  - the abonent's phone number, in case if success
//  - error if abonent was not found in the phone book
func (pb *PhoneBook) Find(abonent string) (string, error) {
	number, exists := pb.abonents[abonent]
	if exists == false {
		return "", errors.New("Abonent was not found")
	}

	return number, nil
}

