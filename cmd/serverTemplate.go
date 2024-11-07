package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
)

var (
	dataDir           = "./data"
	templateDir       = "./templates"
	dataFileMutex     sync.Mutex
	templateFileMutex sync.Mutex
)

func main() {
	http.HandleFunc("/saveData", saveDataHandler)
	http.HandleFunc("/saveTemplate", saveTemplateHandler)
	http.HandleFunc("/render", renderHandler)

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

func saveDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var jsonData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	dataID := uuid.New().String()
	dataFilePath := filepath.Join(dataDir, dataID+".json")

	dataFileMutex.Lock()
	defer dataFileMutex.Unlock()

	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		http.Error(w, "Failed to create data directory", http.StatusInternalServerError)
		return
	}

	file, err := os.Create(dataFilePath)
	if err != nil {
		http.Error(w, "Failed to create data file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(jsonData); err != nil {
		http.Error(w, "Failed to write data to file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(dataID))
}

func saveTemplateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	templateData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	templateID := uuid.New().String()
	templateFilePath := filepath.Join(templateDir, templateID+".tmpl")

	templateFileMutex.Lock()
	defer templateFileMutex.Unlock()

	if err := os.MkdirAll(templateDir, os.ModePerm); err != nil {
		http.Error(w, "Failed to create template directory", http.StatusInternalServerError)
		return
	}

	if err := ioutil.WriteFile(templateFilePath, templateData, os.ModePerm); err != nil {
		http.Error(w, "Failed to write template to file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(templateID))
}

func renderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	type renderRequest struct {
		DataID     string `json:"dataId"`
		TemplateID string `json:"templateId"`
	}

	var req renderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request JSON", http.StatusBadRequest)
		return
	}

	dataFilePath := filepath.Join(dataDir, req.DataID+".json")
	templateFilePath := filepath.Join(templateDir, req.TemplateID+".tmpl")

	dataFile, err := os.Open(dataFilePath)
	if err != nil {
		http.Error(w, "Data file not found", http.StatusNotFound)
		return
	}
	defer dataFile.Close()

	var jsonData map[string]interface{}
	if err := json.NewDecoder(dataFile).Decode(&jsonData); err != nil {
		http.Error(w, "Failed to decode data JSON", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		http.Error(w, "Template file not found or invalid", http.StatusNotFound)
		return
	}

	var result string
	buf := &resultBuffer{buf: &result}
	if err := tmpl.Execute(buf, jsonData); err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

type resultBuffer struct {
	buf *string
}

func (r *resultBuffer) Write(p []byte) (n int, err error) {
	*r.buf += string(p)
	return len(p), nil
}
