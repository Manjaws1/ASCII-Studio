package storage

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type Project struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Banner    string `json:"banner"`
	Color     string `json:"color"`
	Alignment string `json:"alignment"`
}

var mutex sync.Mutex
const dataDir = "data"

// SaveProject stores the project details in a JSON file and returns the generated ID.
func SaveProject(p Project) (string, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Ensure data directory exists
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return "", err
	}

	// Generate ID if not present
	if p.ID == "" {
		bytes := make([]byte, 8)
		if _, err := rand.Read(bytes); err != nil {
			return "", err
		}
		p.ID = hex.EncodeToString(bytes)
	}

	filePath := filepath.Join(dataDir, p.ID+".json")
	fileData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "", err
	}

	if err := os.WriteFile(filePath, fileData, 0644); err != nil {
		return "", err
	}

	return p.ID, nil
}

// LoadProject loads a project from its JSON file by ID.
func LoadProject(id string) (Project, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var p Project
	filePath := filepath.Join(dataDir, id+".json")

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return p, fmt.Errorf("project not found")
	}

	if err := json.Unmarshal(fileData, &p); err != nil {
		return p, err
	}

	return p, nil
}
