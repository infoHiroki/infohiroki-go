package models

import (
	"encoding/json"
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
	result := "# " + b.Title + "\n\n"

	if b.Description != "" {
		result += b.Description + "\n\n"
	}

	result += "**作成日:** " + b.CreatedDate.Format("2006年01月02日") + "\n\n"

	if b.Tags != "" {
		result += "**タグ:** " + b.Tags + "\n\n"
	}

	result += "---\n\n" + b.Content

	return result
}

// GetTagsSlice parses the JSON tags string and returns a slice of strings
func (b *BlogPost) GetTagsSlice() []string {
	if b.Tags == "" {
		return []string{}
	}

	var tags []string
	err := json.Unmarshal([]byte(b.Tags), &tags)
	if err != nil {
		// If JSON parsing fails, treat as a single tag
		return []string{b.Tags}
	}

	return tags
}