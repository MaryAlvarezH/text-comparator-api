package entity

import "time"

// TextComparison represents an compare iteraction between two texts to store an history of comparisons.
type TextComparison struct {
	UserID           int       `json:"userID"`
	Date             time.Time `json:"date" binding:"required"`
	VowelsFisrtText  int       `json:"vowelsFisrtText"`
	VowelsSecondText int       `json:"vowelsSecondText"`
}

func (TextComparison) TableName() string {
	return "app.text_comparisons"
}
