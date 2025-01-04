package content

import (
	"time"

	"github.com/nickfoden/web-log/internal/models"
)

var posts = []models.Post{
	{Title: "New Web Log Alert",
		Content:   "",
		CreatedAt: time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC),
		Slug:      "1"},
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
