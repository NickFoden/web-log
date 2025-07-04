package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nickfoden/web-log/internal/content"
	"github.com/nickfoden/web-log/internal/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filesDir := http.Dir(filepath.Join(workDir, "static"))
	FileServer(r, "/static", filesDir)

	posts := content.GetAllPosts()

	// Handlers
	blogHandler := handlers.NewBlogHandler(posts)

	// Pages
	r.Get("/", blogHandler.Index)
	r.Get("/about", blogHandler.About)
	r.Get("/posts/{slug}", blogHandler.Post)

	// API
	r.Get("/get_current_year", (func(w http.ResponseWriter, r *http.Request) {
		year := time.Now().Year()
		fmt.Fprintf(w, "%d", year)
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Cache-Control", "max-age=3600")

		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
