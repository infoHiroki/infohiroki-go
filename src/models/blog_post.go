package models

import (
	"time"
	"gorm.io/gorm"
)

// BlogPost represents a blog article
type BlogPost struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Slug        string         `json:"slug" gorm:"uniqueIndex;not null"`
	Title       string         `json:"title" gorm:"not null"`
	Content     string         `json:"content" gorm:"type:text"`
	Description string         `json:"description"`
	Tags        string         `json:"tags"` // JSON array as string
	Icon        string         `json:"icon"`
	CreatedDate time.Time      `json:"created_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Published   bool           `json:"published" gorm:"default:true"`
}

// TableName specifies the table name for BlogPost model
func (BlogPost) TableName() string {
	return "blog_posts"
}

// ToMarkdown converts the blog post to markdown format
func (b *BlogPost) ToMarkdown() string {
	return "# " + b.Title + "\n\n" + b.Content
}