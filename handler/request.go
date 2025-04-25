package handler

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func validateRequiredParam(name, typ string) error {
	return fmt.Errorf("parameter %s (type: %s) is required", name, typ)
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Link == "" && r.Salary <= 0 {
		return fmt.Errorf("invalid request body")
	}

	if r.Role == "" {
		return validateRequiredParam("role", "string")
	}

	if r.Company == "" {
		return validateRequiredParam("company", "string")
	}

	if r.Location == "" {
		return validateRequiredParam("location", "string")
	}

	if r.Remote == nil {
		return validateRequiredParam("remote", "bool")
	}

	if r.Link == "" {
		return validateRequiredParam("link", "string")
	}

	if r.Salary <= 0 {
		return validateRequiredParam("salary", "int64")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
