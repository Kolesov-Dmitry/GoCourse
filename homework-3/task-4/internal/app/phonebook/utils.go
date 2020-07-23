package phonebook

import (	
	"encoding/json"
	"io/ioutil"
)

// loadFromJsonFile loads phone book from JSON file
// Input:
//  - path is the path to the JSON file
// Output:
//  - map which contains abonents and their phone numbers
func loadFromJsonFile(path string) map[string]string {
	result := make(map[string]string)

	data, err := ioutil.ReadFile(path)
	if err != nil {		
		json.Unmarshal(data, &result)
	}

	return result
}

// loadFromJsonFile loads phone book from JSON file
// Input:
//  - path is the path to the JSON file
//  - abonents is a map which contains abonents and their phone numbers
// Output:
//  - nil if there is no error occures during the save process
//  - othrewise, returns the error
func saveToJsonFile(path string, abonents map[string]string) error {
	// represent the abonents as JSON
	jsonData, err := json.Marshal(abonents)
	if err != nil {
		return err
	}

	// save JSON to the file	
	if err := ioutil.WriteFile(path, jsonData, 0644); err != nil {		
		return err			
	}
	
	return nil	
}