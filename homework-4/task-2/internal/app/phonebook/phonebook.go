package phonebook

import (	
	"fmt"
	"strings"
)

// PhoneBook data type
type PhoneBook []Abonent

// Append inserts new number into the phone book
// Input:
//  - abonentName is the name of the abonent
//  - abonentNumber is the phone number of the abonent
func (pb *PhoneBook) Append(abonentName, abonentNumber string) {
	*pb = append(*pb, Abonent { abonentName, abonentNumber })
}

// Stringer interface implementation of PhoneBook data type
func (pb PhoneBook) String() string {	
	var builder strings.Builder
	for _, abonent := range pb {		
		str := fmt.Sprintf("%s: %s\n", abonent.Name, abonent.Phone)		
		builder.WriteString(str)
	}

	return builder.String()
}

// Sorting interface implemetation of PhoneBook data type
func (pb PhoneBook) Len() int {
	return len(pb)
}

func (pb PhoneBook) Less(i, j int) bool {
	return pb[i].Name < pb[j].Name
}

func (pb PhoneBook) Swap(i, j int) {
	pb[i], pb[j] = pb[j], pb[i]
}