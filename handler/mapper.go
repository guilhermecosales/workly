package handler

import "github.com/guilhermecosales/workly/schemas"

func mapToOpening(request CreateOpeningRequest) schemas.Opening {
	logger.Debugf("Mapping request to Opening schema: %+v", request)
	return schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}
}

func mapToOpeningResponse(opening schemas.Opening) schemas.OpeningResponse {
	logger.Debugf("Mapping Opening schema to response: %+v", opening)
	return schemas.OpeningResponse{
		Id:        opening.ID,
		Role:      opening.Role,
		Company:   opening.Company,
		Location:  opening.Location,
		Remote:    opening.Remote,
		Link:      opening.Link,
		Salary:    opening.Salary,
		CreatedAt: opening.CreatedAt,
		UpdatedAt: opening.UpdatedAt,
		DeletedAt: opening.DeletedAt.Time,
	}
}
