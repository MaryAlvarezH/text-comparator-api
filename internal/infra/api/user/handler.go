package user

import (
	"errors"
	"net/http"

	"github.com/MaryAlvarezH/text-comparator/internal/pkg/entity"
	"github.com/MaryAlvarezH/text-comparator/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService ports.UserService
}

func newHandler(service ports.UserService) *userHandler {
	return &userHandler{
		userService: service,
	}
}

func (u *userHandler) CreateUser(c *gin.Context) {
	user := &entity.User{}
	if err := c.Bind(user); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Invalid Input"))
		return
	}

	if err := u.userService.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (u *userHandler) Login(c *gin.Context) {
	credentials := &entity.DefaultCredentials{}

	if err := c.Bind(credentials); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Invalid Input"))
		return
	}

	token, err := u.userService.Login(credentials)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
