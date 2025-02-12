package content

import (
	"time"

	"github.com/nickfoden/web-log/internal/models"
)

var posts = []models.Post{
	{Title: "New Web Log Alert",
		Content:        "",
		ContentPreview: "A \"simple\" web log built with Go and html + htmx.",
		CreatedAt:      time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC),
		Slug:           "1"},
	{Title: "No Bugs Here",
		Content:        "",
		ContentPreview: "C'mon just use a proper tsconfig",
		CreatedAt:      time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC),
		Slug:           "no-bugs"},
	{Title: "Ruby for React Native",
		Content:        "",
		ContentPreview: "RN setup issues? Ensure Ruby is current",
		CreatedAt:      time.Date(2025, 1, 16, 0, 0, 0, 0, time.UTC),
		Slug:           "react-native-requires-current-ruby"},
}

func GetAllPosts() []models.Post {
	return posts
}

func GetPost(slug string) models.Post {
	posts := GetAllPosts()
	for _, post := range posts {
		if post.Slug == slug {
			return post
		}
	}
	return models.Post{}
}
