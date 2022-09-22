package comparison

import (
	"errors"
	"testing"
	"time"

	"github.com/MaryAlvarezH/text-comparator/internal/pkg/entity"
	"github.com/MaryAlvarezH/text-comparator/internal/pkg/ports"
	"github.com/MaryAlvarezH/text-comparator/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type comparisonServiceSuite struct {
	suite.Suite
	repo    *mocks.ComparisonRepository
	service ports.ComparisonService
}

func (s *comparisonServiceSuite) SetupTest() {
	s.repo = new(mocks.ComparisonRepository)
	s.service = NewService(s.repo)
}

func (s *comparisonServiceSuite) TearDownTest() {
	s.SetupTest()
}

// TestGetUserTextComparisonsPositive returns valid text comparisons.
func (s *comparisonServiceSuite) TestGetUserTextComparisonsPositive() {
	var expectedTextComparisons []*entity.TextComparison

	mockedTextComparisons := []*entity.TextComparison{
		{
			ID:                     1,
			UserID:                 1,
			Date:                   time.Now(),
			FirstText:              "dummy text 1",
			SecondText:             "dummy text 2",
			VowelsFisrtText:        2,
			VowelsSecondText:       2,
			TransformationDistance: 1,
		},
		{
			ID:                     2,
			UserID:                 1,
			Date:                   time.Now(),
			FirstText:              "orange",
			SecondText:             "apple",
			VowelsFisrtText:        3,
			VowelsSecondText:       2,
			TransformationDistance: 5,
		},
	}

	s.repo.On("FindInOrder",
		&expectedTextComparisons,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int")).Run(func(args mock.Arguments) {
		r := args.Get(0).(*[]*entity.TextComparison)
		*r = mockedTextComparisons
		expectedTextComparisons = *r
	}).Return(nil)

	got, err := s.service.GetUserTextComparisons(1)

	s.NoError(err)
	s.ElementsMatch(expectedTextComparisons, got)
}

// TestGetUserTextComparisonsPositive returns an error from db getting text comparisons.
func (s *comparisonServiceSuite) TestGetUserTextComparisonsNegative() {
	var expectedTextComparisons []*entity.TextComparison

	s.repo.On("FindInOrder",
		&expectedTextComparisons,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("int"),
	).Return(errors.New("Something went wrong getting data from db."))

	got, err := s.service.GetUserTextComparisons(1)

	s.Error(err)
	s.ElementsMatch(expectedTextComparisons, got)
}

func TestComparisonService(t *testing.T) {
	suite.Run(t, new(comparisonServiceSuite))
}
