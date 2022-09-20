package api

import (
	"github.com/MaryAlvarezH/text-comparator/internal/infra/api/comparison"
	"github.com/MaryAlvarezH/text-comparator/internal/infra/api/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	// User routes
	user.RegisterRoutes(e)

	// Text comparison routes
	comparison.RegisterRoutes(e)
}
