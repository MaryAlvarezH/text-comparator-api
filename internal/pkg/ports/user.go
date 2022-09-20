package ports

import "github.com/MaryAlvarezH/text-comparator/internal/pkg/entity"

type UserRepository interface {
	Create(value interface{}) error
	First(out interface{}, conditions ...interface{}) error
}

type UserService interface {
	Create(user *entity.User) error
	Login(credentials *entity.DefaultCredentials) (string, error)
}
