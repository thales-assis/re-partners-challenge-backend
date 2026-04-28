package viewmodel

type CalculatorPacksRequest struct {
	Amount int `json:"amount" example:"500000"`
}

func (vm CalculatorPacksRequest) Validate() error {

	validationErrors := NewValidationsError()

	if vm.Amount < 0 {
		validationErrors.Append("amount", "must be a valid value greater than -1")
	}

	if validationErrors.HasError() {
		return validationErrors
	}

	return nil
}

type CalculatorPacksResponse struct {
	PackSize int `json:"pack_size" example:"23"`
	Quantity int `json:"quantity" example:"2"`
}
