package property

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/rmh-softengineer/locqube/api/model"
)

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) GetProperties() ([]model.Property, error) {
	// Find and read the properties.json file
	filePath := filepath.Join("json", "properties.json")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Parse JSON into properties slice
	var properties []model.Property
	err = json.Unmarshal(bytes, &properties)
	if err != nil {
		return nil, err
	}

	return properties, nil
}
