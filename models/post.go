package models

import "github.com/google/uuid"

type Post struct {
	UID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Slug    string    `json:"slug" gorm:"uniqueIndex"`
}

type NewPostMessage struct {
	UID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

type PublishedPostMessage struct {
	Post
}
