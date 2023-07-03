package work

import (
	"fmt"
	"os"

	"encoding/json"
)

// LoadList converts a file into a list
func LoadList(fileName string) (*List, error) {
	// initialize list
	var list List

	// if file does not exist, just return an empty list
	// otherwise, return an error
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return &list, nil
	} else if err != nil {
		return nil, fmt.Errorf("error reading file info: %w", err)
	}

	// read the file
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// unmarsal file bytes into a list
	err = json.Unmarshal(fileBytes, &list)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling file bytes into list: %w", err)
	}

	return &list, nil
}

// SaveList stores the list in a file
func SaveList(fileName string, list *List) error {
	// convert object to JSON bytes
	fileBytes, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling list into file bytes: %w", err)
	}

	// write JSON bytes to file
	err = os.WriteFile(fileName, fileBytes, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
