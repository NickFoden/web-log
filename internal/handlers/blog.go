package handlers

import (
	"html/template"
	"net/http"

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
		"Title": "Nick Web Log",
		"Posts": h.posts,
	})
}

func (h *BlogHandler) Post(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	for _, post := range h.posts {
		if post.Slug == slug {
			renderTemplate(w, "post.html", map[string]interface{}{
				"Title": post.Title,
				"Post":  post,
			})
			return
		}
	}

	http.NotFound(w, r)
}
