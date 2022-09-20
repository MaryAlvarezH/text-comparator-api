package entity

import "time"

// TextComparison represents a compare iteraction between two texts to store an history of comparisons in db.
type TextComparison struct {
	ID               int       `json:"id"`
	UserID           int       `json:"userID"`
	Date             time.Time `json:"date" binding:"required"`
	FirstText        string    `json:"firstText"`
	SecondText       string    `json:"secondText"`
	VowelsFisrtText  int       `json:"vowelsFisrtText"`
	VowelsSecondText int       `json:"vowelsSecondText"`
}

// TextComparisonRequest represents the body payload to make a comparison request.
type TextComparisonRequest struct {
	UserID int    `json:"userID"`
	Text1  string `json:"text1"`
	Text2  string `json:"text2"`
}

type TextComparisonResults struct {
	DiffFirstText  []*ComparisonResult `json:"diffFirstText"`
	DiffSecondText []*ComparisonResult `json:"diffSecondText"`
}

type ComparisonResult struct {
	Character string `json:"character"`
	IsDiff    bool   `json:"status"`
}

func (TextComparison) TableName() string {
	return "app.text_comparisons"
}
