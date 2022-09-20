package comparison

import (
	"github.com/MaryAlvarezH/text-comparator/internal/infra/api/middlewares"
	"github.com/MaryAlvarezH/text-comparator/internal/infra/repositories/postgres"
	"github.com/MaryAlvarezH/text-comparator/internal/pkg/service/comparison"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	repo := postgres.NewClient()
	service := comparison.NewService(repo)
	handler := newHandler(service)

	// Make a text comparison
	e.POST("/api/v1/compare", middlewares.Authenticate(), handler.CompareText)

	// Get text comparison history
	e.GET("/api/v1/comparisons", middlewares.Authenticate(), handler.GetUserTextComparisons)
}
