package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermecosales/workly/schemas"
	"net/http"
)

func RetrieveAllOpenings(context *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendErrorResponse(context, http.StatusInternalServerError, "Error retrieving all openings")
		return
	}

	openingResponses := make([]schemas.OpeningResponse, len(openings))
	for i, opening := range openings {
		openingResponses[i] = mapToOpeningResponse(opening)
	}

	sendSuccessResponse(context, "Retrieve all openings", openingResponses)
}
