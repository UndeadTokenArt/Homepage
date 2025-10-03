package main

import (
	"encoding/json"
	"os"

	"github.com/undeadtokenart/Homepage/structs"
)

// ParseConfigFile reads a JSON config file and unmarshals it into a Homepage struct.
func ParseConfigFile(path string) (*structs.Homepage, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var homepage structs.Homepage
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&homepage); err != nil {
		return nil, err
	}
	return &homepage, nil
}
