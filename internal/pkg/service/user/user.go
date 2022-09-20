package user

import (
	"net/mail"
	"time"

	"github.com/MaryAlvarezH/text-comparator/internal/pkg/entity"
	"github.com/MaryAlvarezH/text-comparator/internal/pkg/ports"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo ports.UserRepository
}

// NewService creates a service instance using UserRepository.
func NewService(repo ports.UserRepository) *service {
	return &service{
		repo: repo,
	}
}

// Create adds a new user in db.
func (s *service) Create(user *entity.User) error {
	// 1. Valid email address
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return err
	}

	// 2. Hash password
	user.Password = hashAndSalt(user.Password)

	// 3. Save into db
	return s.repo.Create(user)
}

// Login uses credentials to authenticate the user returning a Bearer token.
func (s *service) Login(credentials *entity.DefaultCredentials) (string, error) {
	user := &entity.User{}

	// 1. Loking for user that matches with given email
	if err := s.repo.First(user, "email = ?", credentials.Email); err != nil {
		return "", err
	}

	// 2. Try match passwords
	if err := tryMatchPasswords(user.Password, credentials.Password); err != nil {
		return "", err
	}

	// 3. Create session token
	return createToken(user)

}

func createToken(user *entity.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = user.ID
	claims["user"] = user
	claims["exp"] = time.Hour * 24

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := jwtToken.SignedString([]byte("textcomparisonapipass754267"))
	if err != nil {
		return "", err
	}

	return strToken, nil
}

func tryMatchPasswords(userPassword, crerentialPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(crerentialPassword)); err != nil {
		return err
	}
	return nil
}

func hashAndSalt(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		log.Error(err)
	}

	return string(hash)
}
