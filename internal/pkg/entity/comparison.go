package entity

import "time"

// TextComparison represents a compare iteraction between two texts to store an history of comparisons in db.
type TextComparison struct {
	ID                     int       `json:"id"`
	UserID                 int       `json:"userID"`
	Date                   time.Time `json:"date" binding:"required"`
	FirstText              string    `json:"firstText"`
	SecondText             string    `json:"secondText"`
	VowelsFisrtText        int       `json:"vowelsFisrtText"`
	VowelsSecondText       int       `json:"vowelsSecondText"`
	TransformationDistance int       `json:"transformationDistance"`
}

// TextComparisonRequest represents the body payload to make a comparison request.
type TextComparisonRequest struct {
	UserID int    `json:"userID"` // used to save the comparison result in db.
	Text1  string `json:"text1"`
	Text2  string `json:"text2"`
}

// TextComparisonResponse represets the response format for a texts comparison.
type TextComparisonResponse struct {
	HtmlDiff               string `json:"htmlDiff"`
	TransformationDistance int    `json:"transformationDistance"`
}

func (TextComparison) TableName() string {
	return "app.text_comparisons"
}
