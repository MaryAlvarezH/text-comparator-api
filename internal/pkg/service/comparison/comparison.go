package comparison

import (
	"errors"
	"fmt"

	"github.com/MaryAlvarezH/text-comparator/internal/pkg/entity"
	"github.com/MaryAlvarezH/text-comparator/internal/pkg/ports"
)

type service struct {
	repo ports.ComparisonRepository
}

// NewService creates a service instance using UserRepository.
func NewService(repo ports.ComparisonRepository) *service {
	return &service{
		repo: repo,
	}
}

// CompareTexts returns the difference between two texts and save the comparison in db.
func (s *service) CompareTexts(comparisonRequest *entity.TextComparisonRequest) (*entity.TextComparisonResults, error) {
	comparisonResult := &entity.TextComparisonResults{}
	fmt.Println("CompareText")
	return comparisonResult, nil
}

// GetUserTextComparisons gets the text comparison history made by specific user.
func (s *service) GetUserTextComparisons(userID int) ([]*entity.TextComparison, error) {
	var comparisonsResponse []*entity.TextComparison

	if err := s.repo.FindInOrder(&comparisonsResponse, "date desc", "user_id=?", userID); err != nil {
		return nil, errors.New("failed to get text comparisons for this user")
	}

	return comparisonsResponse, nil
}
