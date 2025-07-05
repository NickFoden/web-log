package content

import (
	"sort"

	"github.com/nickfoden/web-log/internal/models"
)

func GetAllPosts() []models.Post {
	posts := make([]models.Post, 0, len(PostsLibrary))
	for _, post := range PostsLibrary {
		posts = append(posts, post)
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})

	return posts
}

func GetPost(slug string) models.Post {
	if post, exists := PostsLibrary[slug]; exists {
		return post
	}
	return models.Post{}
}
