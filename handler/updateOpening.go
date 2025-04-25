package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermecosales/workly/schemas"
	"net/http"
)

func UpdateOpening(context *gin.Context) {
	request := UpdateOpeningRequest{}

	openingId := context.Param("openingId")
	if len(openingId) == 0 {
		logger.Warn("Missing Opening ID in request")
		sendErrorResponse(context, http.StatusBadRequest, "Opening ID is required")
		return
	}

	if err := context.BindJSON(&request); err != nil {
		logger.Errorf("Error binding JSON: %v", err.Error())
		logger.Infof(
			"Request payload: %v",
			context.Request.Body)
		sendErrorResponse(context, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, openingId).Error; err != nil {
		logger.Errorf("Error retrieving opening: %v", err.Error())
		sendErrorResponse(context, http.StatusNotFound, "Opening not found")
		return
	}

	mapRequestToOpening(&request, &opening)

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Error updating opening: %v", err.Error())
		sendErrorResponse(context, http.StatusInternalServerError, "Failed to update opening")
		return
	}

	sendSuccessResponse(context, "Opening updated successfully", opening)
}

func mapRequestToOpening(request *UpdateOpeningRequest, opening *schemas.Opening) {
	logger.Info("Mapping request to opening")
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
}
