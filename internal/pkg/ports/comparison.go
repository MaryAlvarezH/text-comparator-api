package ports

import "github.com/MaryAlvarezH/text-comparator/internal/pkg/entity"

type ComparisonRepository interface {
	Create(value interface{}) error
	First(out interface{}, conditions ...interface{}) error
	FindInOrder(out interface{}, order string, where ...interface{}) error
}

type ComparisonService interface {
	CompareTexts(*entity.TextComparisonRequest) (*entity.TextComparisonResponse, error)
	GetUserTextComparisons(userID int) ([]*entity.TextComparison, error)
}
