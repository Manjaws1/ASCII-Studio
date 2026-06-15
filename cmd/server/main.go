package main

import (
	"ascii-studio/internal/renderer"
	"ascii-studio/internal/storage"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

type GenerateRequest struct {
	Text      string `json:"text"`
	Banner    string `json:"banner"`
	Color     string `json:"color"`
	Alignment string `json:"alignment"`
}

type GenerateResponse struct {
	Result string `json:"result"`
	Error  string `json:"error,omitempty"`
}

func main() {
	// API routes
	http.HandleFunc("/api/generate", handleGenerate)
	http.HandleFunc("/api/save-project", handleSaveProject)
	http.HandleFunc("/api/load-project/", handleLoadProject)
	http.HandleFunc("/api/banners", handleBanners)

	// Static files and SPA
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handleIndex)

	port := ":8087"
	log.Printf("Server is starting on port %s...", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/index.html" {
		http.NotFound(w, r)
		return
	}
	content, err := os.ReadFile("web/templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(content)
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req GenerateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(w, http.StatusBadRequest, GenerateResponse{Error: "Invalid JSON format"})
		return
	}

	if req.Text == "" || req.Banner == "" {
		jsonResponse(w, http.StatusBadRequest, GenerateResponse{Error: "Text and Banner are required"})
		return
	}

	result, err := renderer.Process(req.Text, req.Banner, req.Color, req.Alignment)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, GenerateResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, http.StatusOK, GenerateResponse{Result: result})
}

func handleSaveProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var p storage.Project
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	id, err := storage.SaveProject(p)
	if err != nil {
		http.Error(w, "Failed to save project", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, http.StatusOK, map[string]string{"id": id})
}

func handleLoadProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/api/load-project/")
	if id == "" {
		http.Error(w, "Project ID is required", http.StatusBadRequest)
		return
	}

	p, err := storage.LoadProject(id)
	if err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	jsonResponse(w, http.StatusOK, p)
}

func handleBanners(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Simplified static list for now, ideally scan banners/ dir
	banners := []string{"standard", "shadow", "thinkertoy"}
	jsonResponse(w, http.StatusOK, banners)
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
