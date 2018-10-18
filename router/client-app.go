package router

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
)

func renderClientApp(appPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, appPath)
	}
}

func serveStatic(staticPath string, root http.FileSystem) http.HandlerFunc {
	fs := http.StripPrefix(staticPath, http.FileServer(root))

	return func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}
}

func ClientAppRouter() http.Handler {
	workDir, _ := os.Getwd()
	staticDir := filepath.Join(workDir, "./client/dist")
	appPath := filepath.Join(workDir, "./template/app.html")

	r := chi.NewRouter()
	r.Get("/static/*", serveStatic("/static/", http.Dir(staticDir)))
	r.Get("/*", renderClientApp(appPath))

	return r
}
