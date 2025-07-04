package content

import (
	"github.com/nickfoden/web-log/internal/models"
)

func GetAllPosts() []models.Post {
	posts := make([]models.Post, 0, len(PostsLibrary))
	for _, post := range PostsLibrary {
		posts = append(posts, post)
	}
	return posts
}

func GetPost(slug string) models.Post {
	if post, exists := PostsLibrary[slug]; exists {
		return post
	}
	return models.Post{}
}
