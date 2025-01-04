package models

import "time"

type Post struct {
	Title          string
	Content        string
	ContentPreview string
	Slug           string
	CreatedAt      time.Time
}
