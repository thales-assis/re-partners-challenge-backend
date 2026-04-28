package viewmodel

type HttpErrorResponse struct {
	Status  int    `json:"status" example:"422"`
	Message string `json:"message" example:"some parameter in the request body is invalid"`
	Code    string `json:"code,omitempty" example:"invalid-request-body"`
	*ValidationErrors
}
