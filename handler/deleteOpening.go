package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/guilhermecosales/workly/schemas"
	"gorm.io/gorm"
	"net/http"
)

func DeleteOpening(context *gin.Context) {
	openingId := context.Param("openingId")

	// Param validation
	if openingId == "" {
		sendErrorResponse(
			context,
			http.StatusBadRequest,
			validateRequiredParam("openingId", "pathParameter").Error(),
		)
		return
	}

	opening := schemas.Opening{}

	// Search for the opening
	if err := db.First(&opening, openingId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sendErrorResponse(
				context,
				http.StatusNotFound,
				"Opening not found",
			)
		} else {
			sendErrorResponse(
				context,
				http.StatusInternalServerError,
				"Failed to retrieve opening",
			)
		}
		return
	}

	// Delete the opening
	if err := db.Delete(&opening).Error; err != nil {
		sendErrorResponse(
			context,
			http.StatusInternalServerError,
			"Failed to delete opening",
		)
		return
	}

	openingResponse := mapToOpeningResponse(opening)

	sendSuccessResponse(context, "Opening deleted successfully", openingResponse)
}
