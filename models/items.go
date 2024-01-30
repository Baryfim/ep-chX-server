package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time
	YaerRefer  int    `json:"year_id"`
	Year       Year   `gorm:"foreignKey:YaerRefer"`
	Date       string `json:"date"`
	Name       string `json:"name"`
	Text       string `json:"text" gorm:"text"`
	SourceLink string `json:"source_link"`
	ImageReal  string `json:"imageReal"`
	ImageAi    string `json:"imageAi"`
	Slug       string `json:"slug"`
}

func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	item.Slug = uuid.String()
	return
}
