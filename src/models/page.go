package models

import (
	"time"
	"gorm.io/gorm"
)

// Page represents a static page (about, services, contact, etc.)
type Page struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Slug            string         `json:"slug" gorm:"uniqueIndex;not null"`
	Title           string         `json:"title" gorm:"not null"`
	Content         string         `json:"content" gorm:"type:text"`
	Template        string         `json:"template" gorm:"default:'page'"`
	MetaDescription string         `json:"meta_description"`
	MetaKeywords    string         `json:"meta_keywords"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName specifies the table name for Page model
func (Page) TableName() string {
	return "pages"
}