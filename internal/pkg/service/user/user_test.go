package user

import (
	"errors"
	"testing"

	"github.com/MaryAlvarezH/text-comparator/internal/pkg/entity"
	"github.com/MaryAlvarezH/text-comparator/internal/pkg/ports"
	"github.com/MaryAlvarezH/text-comparator/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type userServiceSuite struct {
	// Testify suite
	suite.Suite

	// Dependencies
	repo *mocks.UserRepository

	// Service to test
	service ports.UserService
}

// SetupTest runs before any tests.
func (s *userServiceSuite) SetupTest() {
	s.repo = new(mocks.UserRepository)
	s.service = NewService(s.repo)
}

// TearDownTest runs after every tests reseting changes generated to the service suite dependencies.
func (s *userServiceSuite) TearDownTest() {
	s.SetupTest()
}

// TextCreateWithEmailNegative tests user sign up with an invalid email.
func (s *userServiceSuite) TextCreateWithEmailNegative() {
	user := &entity.User{
		Name:     "dummy name",
		Lastname: "dummy lastname",
		Email:    "dummy_invalid_email",
		Password: "dummu password",
	}

	// Mocking result returned by our repository
	s.repo.On(
		"Create",
		mock.Anything,
	).Return(errors.New("Failed to save the user in postgres"))

	// Calling the method we want to test
	err := s.service.Create(user)

	// Assertions
	s.Error(err)
}

// TextCreateWithEmailPositive tests user sign up with a valid email.
func (s *userServiceSuite) TextCreateWithEmailPositive() {
	user := &entity.User{
		Name:     "dummy name",
		Lastname: "dummy lastname",
		Email:    "dummy_invalid_email",
		Password: "dummy_password@epa.digital",
	}

	// Mocking a result returned by our repository
	s.repo.On(
		"Create",
		mock.Anything,
	).Return(nil)

	// Calling the method we want to test
	err := s.service.Create(user)

	// Assertions
	s.NoError(err)
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(userServiceSuite))
}
