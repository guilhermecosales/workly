package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpening(context *gin.Context) {
	request := CreateOpeningRequest{}

	// Bind JSON request to struct
	if err := context.BindJSON(&request); err != nil {
		logger.Errorf("Error binding JSON: %v", err.Error())
		sendErrorResponse(context, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the request
	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	// Map the request to the Opening schema
	opening := mapToOpening(request)

	// Insert into the database
	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error inserting into DB: %v", err.Error())
		sendErrorResponse(context, http.StatusInternalServerError, "Failed to create opening")
		return
	}

	// Map the opening schema to the response schema
	openingResponse := mapToOpeningResponse(opening)

	// Send success response
	sendSuccessResponse(context, "Create Opening successfully", openingResponse)
}
