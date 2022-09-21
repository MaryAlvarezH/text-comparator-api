package comparison

import (
	"errors"
	"time"
	"unicode"

	"github.com/MaryAlvarezH/text-comparator/internal/pkg/entity"
	"github.com/MaryAlvarezH/text-comparator/internal/pkg/ports"
	"github.com/sergi/go-diff/diffmatchpatch"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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
func (s *service) CompareTexts(comparisonRequest *entity.TextComparisonRequest) (string, error) {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(comparisonRequest.Text1, comparisonRequest.Text2, false)

	diffInHtml := dmp.DiffPrettyHtml(diffs)

	comparisonLog := &entity.TextComparison{
		UserID:           comparisonRequest.UserID,
		FirstText:        comparisonRequest.Text1,
		SecondText:       comparisonRequest.Text2,
		Date:             time.Now(),
		VowelsFisrtText:  getVowelsNumber(comparisonRequest.Text1),
		VowelsSecondText: getVowelsNumber(comparisonRequest.Text2),
	}

	if err := s.repo.Create(comparisonLog); err != nil {
		return "", errors.New("failed to save text comparison in history")
	}

	return diffInHtml, nil
}

func getVowelsNumber(str string) int {
	s := normalizeText(str)
	c := 0
	for _, value := range s {
		switch value {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			c++
		}
	}

	return c
}

// normalizeText apply text normalization to remove special signs in strings
// NFD (descomposición canónica de formato de normalización) - Los caracteres se descomponen según su equivalencia canónica.
// NFD (composición canónica de formato de normalización) - Los caracteres se descomponen y después se recomponen según su equivalencia canónica..
func normalizeText(str string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, str)
	return s
}
