package viewmodel

import "encoding/json"

type validationError struct {
	Parameter string `json:"parameter" example:"quantity"`
	Message   string `json:"message" example:"must be a valid quantity with value greater than 0"`
}

func (vm validationError) Error() string {
	b, err := json.Marshal(vm)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

type ValidationErrors struct {
	Errors []validationError `json:"errors,omitempty"`
}

func (vm ValidationErrors) Error() string {
	b, err := json.Marshal(vm)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (vm ValidationErrors) HasError() bool {
	return len(vm.Errors) > 0
}

func (vm *ValidationErrors) Append(parameter string, message string) {
	vm.Errors = append(vm.Errors, validationError{
		Parameter: parameter,
		Message:   message,
	})
}

func NewValidationsError() *ValidationErrors {

	vm := &ValidationErrors{
		Errors: make([]validationError, 0),
	}

	return vm
}
