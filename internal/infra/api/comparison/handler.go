package comparison

import (
	"errors"
	"net/http"

	"github.com/MaryAlvarezH/text-comparator/internal/pkg/entity"
	"github.com/MaryAlvarezH/text-comparator/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type comparisonHandler struct {
	comparisonService ports.ComparisonService
}

func newHandler(service ports.ComparisonService) *comparisonHandler {
	return &comparisonHandler{
		comparisonService: service,
	}
}

func (co *comparisonHandler) CompareText(c *gin.Context) {
	comparisonRequest := &entity.TextComparisonRequest{}

	if err := c.Bind(comparisonRequest); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Invalid text comparison format"))
		return
	}

	comparisonRequest.UserID = int(c.MustGet("userID").(float64))

	comparisonResult, err := co.comparisonService.CompareTexts(comparisonRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, comparisonResult)
}

func (co *comparisonHandler) GetUserTextComparisons(c *gin.Context) {
	userID := int(c.MustGet("userID").(float64))

	textComparisons, err := co.comparisonService.GetUserTextComparisons(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, textComparisons)
}
