package models

type Question struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	TestID        uint
	Answer        string   `json:"answer"`
	Questions     []string `json:"questions"`
	QuestionTitle string   `json:"title_question"`
	ItemRefer     int      `json:"item_id"`
	Item          Item     `gorm:"foreignKey:ItemRefer"`
}
