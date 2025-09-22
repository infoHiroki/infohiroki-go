package models

import (
	"encoding/json"
	"html/template"
	"strings"
	"time"
	"github.com/russross/blackfriday/v2"
)

// BlogPost represents a blog article
type BlogPost struct {
	ID          uint      `json:"id"`
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Description string    `json:"description"`
	Tags        string    `json:"tags"` // JSON array as string
	Icon        string    `json:"icon"`
	ContentType string    `json:"content_type"` // "html" or "markdown"
	MarkdownPath string   `json:"markdown_path"` // .mdファイルパス
	CreatedDate time.Time `json:"created_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Published   bool      `json:"published"`
}

// TableNameメソッドはファイルベースでは不要

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

// IsIconURL checks if the icon field contains a URL or path
func (b *BlogPost) IsIconURL() bool {
	if b.Icon == "" {
		return false
	}
	return strings.HasPrefix(b.Icon, "http") ||
	       strings.HasPrefix(b.Icon, "./") ||
	       strings.HasPrefix(b.Icon, "/")
}

// RenderContent renders the content based on ContentType
func (b *BlogPost) RenderContent() template.HTML {
	if b.ContentType == "markdown" {
		// MarkdownをHTMLに変換
		renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
			Flags: blackfriday.CommonHTMLFlags,
		})

		extensions := blackfriday.CommonExtensions | blackfriday.AutoHeadingIDs
		html := blackfriday.Run([]byte(b.Content), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extensions))

		return template.HTML(html)
	}

	// デフォルトはHTMLコンテンツ
	return template.HTML(b.Content)
}