package models

import "time"

type Tests struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Title     string     `json:"title_test"`
	Questions []Question `gorm:"foreignKey:TestID"`
}
