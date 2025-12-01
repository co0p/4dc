// removed duplicate package declaration
package static

import (
	"net/http"
	"os"
	"path/filepath"
)

const StaticDir = "../../.svelte-kit/output/client"

func StaticFileHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) >= 5 && r.URL.Path[:5] == "/api/" {
		http.NotFound(w, r)
		return
	}
	filePath := r.URL.Path
	if filePath == "/" {
		filePath = "/index.html"
	}
	fullPath := filepath.Join(StaticDir, filePath)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, fullPath)
}
