package handlers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nickfoden/web-log/internal/models"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

type rssItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
	GUID        string   `xml:"guid"`
}

type rssChannel struct {
	XMLName       xml.Name  `xml:"channel"`
	Title         string    `xml:"title"`
	Link          string    `xml:"link"`
	Description   string    `xml:"description"`
	LastBuildDate string    `xml:"lastBuildDate"`
	Items         []rssItem `xml:"item"`
}

type rssFeed struct {
	XMLName xml.Name   `xml:"rss"`
	Version string     `xml:"version,attr"`
	Channel rssChannel `xml:"channel"`
}

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

func (h *BlogHandler) Ai(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "ai.html", map[string]any{
		"Title": "Web Log by Nick Foden",
	})
}

func (h *BlogHandler) About(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "about.html", map[string]any{
		"Title": "Web Log by Nick Foden",
	})
}

func (h *BlogHandler) Post(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	for _, post := range h.posts {
		if post.Slug == slug {
			source, err := os.ReadFile(fmt.Sprintf("internal/content/posts/%s.md", post.Slug))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			md := goldmark.New(
				goldmark.WithRendererOptions(
					html.WithUnsafe(),
				),
			)

			var buf bytes.Buffer
			if err := md.Convert(source, &buf); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			renderTemplate(w, "post.html", map[string]interface{}{
				"Content": template.HTML(buf.String()),
				"Title":   "Web Log by Nick Foden",
				"Post":    post,
			})
			return
		}
	}

	http.NotFound(w, r)
}

func (h *BlogHandler) Feed(w http.ResponseWriter, r *http.Request) {
	baseURL := "https://www.nickfoden.com"

	items := make([]rssItem, 0, len(h.posts))
	for _, post := range h.posts {
		link := baseURL + "/posts/" + post.Slug
		items = append(items, rssItem{
			Title:       post.Title,
			Link:        link,
			Description: post.ContentPreview,
			PubDate:     post.CreatedAt.Format(time.RFC1123Z),
			GUID:        link,
		})
	}

	lastBuildDate := time.Now().Format(time.RFC1123Z)
	if len(h.posts) > 0 {
		lastBuildDate = h.posts[0].CreatedAt.Format(time.RFC1123Z)
	}

	feed := rssFeed{
		Version: "2.0",
		Channel: rssChannel{
			Title:         "Web Log by Nick Foden",
			Link:          baseURL,
			Description:   "A web log by Nick Foden",
			LastBuildDate: lastBuildDate,
			Items:         items,
		},
	}

	w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")

	data, err := xml.MarshalIndent(feed, "", "  ")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Combine header and data into a single write to avoid partial writes
	output := append([]byte(xml.Header), data...)
	if _, err := w.Write(output); err != nil {
		// Cannot send error response after w.Write has been called
		// Log the error if needed
		return
	}
}
