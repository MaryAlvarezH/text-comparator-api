package user

import (
	"github.com/MaryAlvarezH/text-comparator/internal/infra/repositories/postgres"
	"github.com/MaryAlvarezH/text-comparator/internal/pkg/service/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	repo := postgres.NewClient()
	service := user.NewService(repo)
	handler := newHandler(service)

	// Create user
	e.POST("/api/v1/users", handler.CreateUser)

	// Login user
	e.POST("/api/v1/authenticate", handler.Login)
}
