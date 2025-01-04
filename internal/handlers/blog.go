package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/nickfoden/web-log/internal/models"
)

type BlogHandler struct {
	posts []models.Post
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Parse the base template and the child template
	t, err := template.ParseFiles(
		"internal/templates/base.html",
		"internal/templates/"+tmpl,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// Execute the template with the data
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewBlogHandler(posts []models.Post) *BlogHandler {
	return &BlogHandler{posts: posts}
}

func (h *BlogHandler) Index(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "index.html", map[string]interface{}{
		"Title": "Web Log by Nick Foden",
		"Posts": h.posts,
	})
}

func (h *BlogHandler) Post(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	for _, post := range h.posts {
		if post.Slug == slug {
			content, err := os.ReadFile(fmt.Sprintf("internal/content/posts/%s.html", post.Slug))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			renderTemplate(w, "post.html", map[string]interface{}{
				"Content": template.HTML(content),
				"Title":   "Web Log by Nick Foden",
				"Post":    post,
			})
			return
		}
	}

	http.NotFound(w, r)
}
