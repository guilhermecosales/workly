package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/guilhermecosales/workly/schemas"
	"gorm.io/gorm"
	"net/http"
)

func RetrieveOpeningById(context *gin.Context) {
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

	sendSuccessResponse(context, "Opening retrieved successfully", mapToOpeningResponse(opening))
}
