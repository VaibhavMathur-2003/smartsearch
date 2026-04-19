package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveJson(data any, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create json: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("failed to write json data: %w", err)
	}

	return nil
}

func CreateFolders(folderPath string) {
	os.MkdirAll(folderPath, 0755)
}

func SaveParquet() {

}
