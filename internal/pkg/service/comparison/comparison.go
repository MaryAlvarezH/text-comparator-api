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
func (s *service) CompareTexts(comparisonRequest *entity.TextComparisonRequest) (*entity.TextComparisonResponse, error) {
	// 1. Create an empty instance of the response
	var response = &entity.TextComparisonResponse{}

	// 2. Get the texts diference using go-diff package.
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(comparisonRequest.Text1, comparisonRequest.Text2, false)

	// 3. Get the text differe using go-diff package in HTML format.
	response.HtmlDiff = dmp.DiffPrettyHtml(diffs)

	// 4. Get min transformation distance to convert text1 in tex2
	response.TransformationDistance = getTransformationDistance(comparisonRequest.Text1, comparisonRequest.Text2)

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
		VowelsFisrtText:        getVowelsNumber(comparison.Text1),
		VowelsSecondText:       getVowelsNumber(comparison.Text2),
		TransformationDistance: transformationDistance,
	}

	if err := s.repo.Create(comparisonLog); err != nil {
		return errors.New("failed to save text comparison in history")
	}

	return nil
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

func getTransformationDistance(text1, text2 string) int {
	// 1. Get unicode characters for text1 and text2 strings using rune data type
	t1 := []rune(text1)
	t2 := []rune(text2)

	t1Len := len(t1)
	t2Len := len(t2)

	// 2. Create slice with length of text1 + 1
	column := make([]int, t1Len+1)

	// 3. Assign incremental values to column slice
	for y := 1; y <= t1Len; y++ {
		column[y] = y
	}

	// 3. Iterate the items of text 1 with the items of text 2
	for x := 1; x <= t2Len; x++ {
		// assign the last index position of text 2 in column[0] slice
		column[0] = x
		lastkey := x - 1 // last index position before the new assigment

		for y := 1; y <= t1Len; y++ {
			oldkey := column[y]
			var incr int

			// if there's a difference between a character of text 1 and a character of text 2

			if t1[y-1] != t2[x-1] {
				incr = 1
			}

			// gets the minimun value of insert, delete and remove operation
			// it is assumed that there will be at least one update operation after an insert, delete or replace
			column[y] = minimum(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey
		}

	}
	return column[t1Len]
}

func minimum(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
