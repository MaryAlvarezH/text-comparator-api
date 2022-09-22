package comparison

import (
	"errors"
	"time"

	"github.com/MaryAlvarezH/text-comparator/internal/pkg/entity"
	"github.com/MaryAlvarezH/text-comparator/internal/pkg/ports"
	"github.com/MaryAlvarezH/text-comparator/internal/pkg/utils"
	"github.com/sergi/go-diff/diffmatchpatch"
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

// GetUserTextComparisons gets the text comparison history made by specific user.
func (s *service) GetUserTextComparisons(userID int) ([]*entity.TextComparison, error) {
	var comparisonsResponse []*entity.TextComparison

	if err := s.repo.FindInOrder(&comparisonsResponse, "date desc", "user_id=?", userID); err != nil {
		return nil, errors.New("failed to get text comparisons for this user")
	}

	return comparisonsResponse, nil
}

// CompareTexts returns the difference between two texts and save the comparison in db.
func (s *service) CompareTexts(comparisonRequest *entity.TextComparisonRequest) (*entity.TextComparisonResponse, error) {
	// 1. Create an empty instance of the response
	var response = &entity.TextComparisonResponse{}

	// 2. Get the texts diference using go-diff package.
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(comparisonRequest.Text1, comparisonRequest.Text2, false)

	// 3. Get the text differe using go-diff package in HTML format.
	response.HtmlDiff = dmp.DiffPrettyHtml(diffs)

	// 4. Get min transformation distance to convert text1 in tex2
	response.TransformationDistance = utils.GetTransformationDistance(comparisonRequest.Text1, comparisonRequest.Text2)

	// 5. Save text comparison in db.
	if err := s.saveTextComparison(comparisonRequest, response.TransformationDistance); err != nil {
		return response, err
	}

	return response, nil
}

func (s *service) saveTextComparison(comparison *entity.TextComparisonRequest, transformationDistance int) error {
	comparisonLog := &entity.TextComparison{
		UserID:                 comparison.UserID,
		FirstText:              comparison.Text1,
		SecondText:             comparison.Text2,
		Date:                   time.Now(),
		VowelsFisrtText:        utils.GetVowelsNumber(comparison.Text1),
		VowelsSecondText:       utils.GetVowelsNumber(comparison.Text2),
		TransformationDistance: transformationDistance,
	}

	if err := s.repo.Create(comparisonLog); err != nil {
		return errors.New("failed to save text comparison in history")
	}

	return nil
}
