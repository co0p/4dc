package catalog

import (
	"encoding/json"
	"net/http"
)

type CatalogEntry struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
	LastUsed string `json:"lastUsed"`
	User     string `json:"user"`
}

func CatalogHandler(w http.ResponseWriter, r *http.Request) {
	entries := []CatalogEntry{
		{ID: 1, Name: "Drill", ImageURL: "https://example.com/drill.jpg", LastUsed: "2025-11-30", User: "Alice"},
		{ID: 2, Name: "Hammer", ImageURL: "https://example.com/hammer.jpg", LastUsed: "Never", User: "N/A"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}
