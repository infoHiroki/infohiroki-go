package models

import (
	"time"
)

// Page represents a static page (about, services, contact, etc.)
type Page struct {
	ID              uint      `json:"id"`
	Slug            string    `json:"slug"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	Template        string    `json:"template"`
	MetaDescription string    `json:"meta_description"`
	MetaKeywords    string    `json:"meta_keywords"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// TableNameメソッドはファイルベースでは不要