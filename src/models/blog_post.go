package models

import (
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
	Icon        string    `json:"icon"`
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

	result += "---\n\n" + b.Content

	return result
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

// RenderContent renders the markdown content as HTML
func (b *BlogPost) RenderContent() template.HTML {
	// MarkdownをHTMLに変換
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: blackfriday.CommonHTMLFlags,
	})

	extensions := blackfriday.CommonExtensions | blackfriday.AutoHeadingIDs
	html := blackfriday.Run([]byte(b.Content), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extensions))

	return template.HTML(html)
}